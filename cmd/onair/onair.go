package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nh3000-org/radio/config"
)

var memoryStats runtime.MemStats
var playingday string  // MON .. SUN
var playinghour string // 00 .. 23
var schedday string    // MON .. SUN
var schedhour string   // 00 .. 23
var logto bool

//var moh = 0

// schedule
var days string
var hours string
var position string
var categories string

// inventory
var rowid string
var category string
var artist string
var song string
var album string
var songlength string
var rndorder string
var expireson string
var lastplayed string
var dateadded string
var today string
var week string
var total string
var toplay string
var sourcelink string

func adjustToTopOfHour() {
	if logto {
		log.Println("[adjustToTopOfHour]", playingday, playinghour)
	}
}
func getNextDay() {
	clearSpinsPerDayCount()
	if playingday == "MON" {
		schedday = "TUE"
		clearSpinsPerDayCount()
	}
	if playingday == "TUE" {
		schedday = "WED"
		clearSpinsPerDayCount()
	}
	if playingday == "WED" {
		schedday = "THU"
		clearSpinsPerDayCount()
	}
	if playingday == "THU" {
		schedday = "FRI"
		clearSpinsPerDayCount()
	}
	if playingday == "FRI" {

		schedday = "SAT"
		clearSpinsPerDayCount()
	}
	if playingday == "SAT" {
		schedday = "SUN"
		clearSpinsPerDayCount()
	}
	if playingday == "SUN" {
		clearSpinsPerWeekCount()
		schedday = "MON"
	}
	playingday = schedday

}
func clearSpinsPerWeekCount() {
	if logto {
		log.Println("[clearSpinsPerWeekCount]")
	}
	// print daily report to text file
	// print weekly report to text file
}
func clearSpinsPerDayCount() {
	if logto {
		log.Println("[clearSpinsPerDayCount]")
	}
	// print daily report to text file
}

func getNextHourPart() {
	adjustToTopOfHour()
	hp, hperr := strconv.Atoi(playinghour)
	if hperr != nil {
		playinghour = "00"
		schedhour = "00"
		return
	}

	hp++
	if hp > 23 {
		playinghour = "00"
		schedhour = "00"
		getNextDay()
		return
	}

	newhp := strconv.Itoa(hp)
	if len(newhp) == 1 {
		newhp = "0" + newhp
	}
	playinghour = newhp
	schedhour = newhp
	//return newhp
}

var elapsed = 0
var fileid string

func playsetup() oto.Context {

	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.

	op := &oto.NewContextOptions{}

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	op.SampleRate = 44100

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	op.ChannelCount = 2

	// Format of the source. go-mp3's format is signed 16bit integers.
	op.Format = oto.FormatSignedInt16LE

	// Remember that you should **not** create more than one context

	otoCtx, readyChan, err := oto.NewContext(op)
	if err != nil {
		panic("oto.NewContext failed: " + err.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	/* 	// Create a new 'player' that will handle our sound. Paused by default.
	   	player := otoCtx.NewPlayer(decodedMp3)

	   	// Play starts playing the sound and returns without waiting for it (Play() is async).
	   	player.Play()

	   	// We can wait for the sound to finish playing using something like this
	   	for player.IsPlaying() {
	   		elapsed++
	   		time.Sleep(time.Second)
	   	} */

	return *otoCtx

}
func Play(ctx oto.Context, song string, cat string) int {
	elapsed = 0
	log.Println("Playit", song, cat)
	//fileid = strconv.FormatUint(song, 10)
	// if min of hour is even play intro
	// if min of hour is odd play extro
	if cat == "top40" {
		t := time.Now()
		if t.Minute()%2 == 0 {
			song += "INTRO"
		} else {
			song += "OUTRO"
		}

	}
	// Read the mp3 file into memory
	fileBytes := config.GetBucket("mp3", song)
	/* 	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	} */

	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
	fileBytesReader := bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		panic("mp3.NewDecoder failed: " + err.Error())
	}
	// Create a new 'player' that will handle our sound. Paused by default.
	player := ctx.NewPlayer(decodedMp3)

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	player.Play()

	// We can wait for the sound to finish playing using something like this
	for player.IsPlaying() {
		elapsed++
		time.Sleep(time.Second)
	}

	return elapsed
}

var itemlength = 0

func main() {

	schedDay := flag.String("schedday", "MON", "-schedday MON || TUE || WED || THU || FRI || SAT || SUN")
	stationId := flag.String("stationid", "WRRW", "-station WRRW")
	schedHour := flag.String("schedhour", "00", "-schedhour 00..23")
	Logging := flag.String("logging", "true", "-logging true || false")
	flag.Parse()

	playingday = *schedDay
	playinghour = *schedHour
	otoctx := playsetup()

	if *Logging == "true" {
		logto = true
	} else {
		logto = false
	}
	log.Println("Startup Parms:", *schedDay, *schedHour, *stationId, *Logging)

	var TheDB = "postgresql://" + config.DBuser + ":" + config.DBpassword + "@" + config.DBaddress
	dbConfig, err := pgxpool.ParseConfig(TheDB)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), dbConfig)
	if err != nil {
		log.Println("Unable to connect to database: ", err)
		os.Exit(1)
	}
	connectionsched, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}

	_, errscheduleget := connectionsched.Conn().Prepare(context.Background(), "scheduleget", "select * from schedule where days = $1 and hours = $2 order by position")
	if errscheduleget != nil {
		log.Panicln("Prepare scheduleget", errscheduleget)
	}
	/* 	_, errinventoryresetdaily := connection.Conn().Prepare(context.Background(), "inventoryresetdaily", "update inventory  set lastplayed = '2024-01-01 00:00:00',spinstoday = 0")
	   	if errinventoryresetdaily != nil {
	   		log.Panicln("Prepare inventoryresetdaily", errinventoryresetdaily)
	   	}
	   	_, errinventoryresetweekly := connection.Conn().Prepare(context.Background(), "inventoryresetweekly", "update inventory  set lastplayed = '2024-01-01 00:00:00',spinsweek = 0")
	   	if errinventoryresetweekly != nil {
	   		log.Panicln("Prepare inventoryresetweekly", errinventoryresetweekly)
	   	} */
	/* 	_, errinventorygetschedule := connection.Conn().Prepare(context.Background(), "inventorygetschedule", "select * from inventory where category = $1 order by lastplayed, rndorder limit 10")
	   	if errinventorygetschedule != nil {
	   		log.Panicln("Prepare inventorygetschedule", errinventorygetschedule)
	   	} */
	/* 	_, errinventoryget := connection.Conn().Prepare(context.Background(), "inventoryget", "select * from inventory where rowid = $1")
	   	if errinventoryget != nil {
	   		log.Panicln("Prepare errinventoryget", errinventoryget)
	   	}
	   	_, errinventoryupdate := connection.Conn().Prepare(context.Background(), "inventoryupdate", "update inventory set spinstoday = $1, spinsweek = $2, spinstotal = $3, lastplayed = $4, songlength= $5 where rowid = $6")
	   	if errinventoryupdate != nil {
	   		log.Panicln("Prepare inventoryupdate", errinventoryupdate)
	   	} */

	// determine start schedule
	var terminate = 0
	for {
		terminate++
		if terminate > 2 {
			log.Panicln("Reached Termination Point")
		}
		runtime.GC()
		runtime.ReadMemStats(&memoryStats)
		log.Println("Memory start:", playingday, playinghour, strconv.FormatUint(memoryStats.Alloc/1024/1024, 10)+" Mib")

		schedulerows, schedulerowserr := connectionsched.Query(context.Background(), "scheduleget", playingday, playinghour)
		log.Println("reading schedule next ", playingday, playinghour, categories)
		for schedulerows.Next() {
			/* 			invgetconn, _ := connPool.Acquire(context.Background())
			   			_, errinventorygetschedule := invgetconn.Conn().Prepare(context.Background(), "inventorygetschedule", "select * from inventory where category = $1 order by lastplayed, rndorder limit 10")
			   			if errinventorygetschedule != nil {
			   				log.Panicln("Prepare inventorygetschedule", errinventorygetschedule)
			   			} */
			scheduleerr := schedulerows.Scan(&rowid, &days, &hours, &position, &categories, &toplay)
			log.Println("reading schedule: ", days, hours, position, categories, toplay)
			spinstoplay, _ := strconv.Atoi(toplay)
			if scheduleerr != nil {
				log.Panicln("Error scheduleerr " + scheduleerr.Error())
			}
			//if scheduleerr == nil {
			for spinstoplay > 0 {
				// get an inventory item to play
				invgetconn, _ := connPool.Acquire(context.Background())
				_, errinventorygetschedule := invgetconn.Conn().Prepare(context.Background(), "inventorygetschedule", "select * from inventory where category = $1 order by lastplayed, rndorder limit 10")
				if errinventorygetschedule != nil {
					log.Panicln("Prepare inventorygetschedule", errinventorygetschedule)
				}
				invrows, invrowserr := invgetconn.Query(context.Background(), "inventorygetschedule", categories)
				if invrowserr != nil {
					log.Fatal("Error reading inventory ", invrowserr, " cat: ", categories)
				}
				log.Println("inventory schedule get ", categories, invrowserr)
				for invrows.Next() {

					inverr := invrows.Scan(&rowid, &category, &artist, &song, &album, &songlength, &rndorder, &expireson, &lastplayed, &dateadded, &today, &week, &total, &sourcelink)
					log.Println("processing inventory song get" + song)
					if inverr != nil {
						log.Println("processing inventory song get " + inverr.Error())
					}
					// play the item
					itemlength = Play(otoctx, rowid, category)
					// update statistics
					spinsweek, _ := strconv.Atoi(week)
					spinsweek++
					spinstoday, _ := strconv.Atoi(today)
					spinstoday++
					spinstotal, _ := strconv.Atoi(total)
					spinstotal++
					lastplayed = time.Now().String()
					log.Println("last played", lastplayed)
					invupdconn, _ := connPool.Acquire(context.Background())
					_, errinventoryupd := invupdconn.Conn().Prepare(context.Background(), "inventoryupdate", "update inventory set spinstoday = $1, spinsweek = $2, spinstotal = $3, lastplayed = $4, songlength= $5 where rowid = $6")
					if errinventoryupd != nil {
						log.Panicln("Prepare inventory upd", errinventorygetschedule)
					}
					_, invupderr := invupdconn.Exec(context.Background(), "inventoryupdate", spinstoday, spinsweek, spinstoday, lastplayed, itemlength, rowid)
					if invupderr != nil {
						log.Println("updating inventory " + invupderr.Error())
					}

					config.SendONAIR(*stationId, " - "+song)
					log.Println("Expires on", expireson)
					ex, exerr := time.Parse(time.RFC3339, expireson)
					if exerr != nil {
						log.Panicln("inventory time parse " + exerr.Error())
					}
					log.Println("EXPIRES: ",ex.String()," NOW: ",time.Now().String())
					if time.Now().After(ex) {

						invdelconn, _ := connPool.Acquire(context.Background())
						_, errinventorydelete := invdelconn.Conn().Prepare(context.Background(), "inventorydelete", "delete from inventory where rowid = $1")
						if errinventorydelete != nil {
							log.Panicln("Prepare inventory delete", errinventorydelete)
						}

						_, invdelerr := invdelconn.Exec(context.Background(), "inventorydelete", rowid)
						if invdelerr != nil {
							log.Println("deleting inventory " + invdelerr.Error())
						}
						invdelconn.Release()

						//fileid = strconv.FormatUint(rowid, 10)
						var intro = rowid + "INTRO"
						var outro = rowid + "OUTRO"
						errremove := config.DeleteBucket("mp3s", fileid)
						if errremove != nil {
							log.Println("deleting  failed: ", errremove.Error(), fileid)
						}
						errremovei := config.DeleteBucket("mp3s", intro)
						if errremovei != nil {
							log.Println("deleting  failed: ", errremovei.Error(), intro)
						}
						errremoveo := config.DeleteBucket("mp3s", outro)
						if errremoveo != nil {
							log.Println("deleting  failed: ", errremoveo.Error(), outro)
						}

					}
					// check inventory expired

					if invrowserr != nil {
						log.Println("reading inventory " + invrowserr.Error())
					}
					log.Println("spins to play ", spinstoplay)
					//conninv.Release()
				}
				//spinstoplay--
				invgetconn.Release()

				log.Println("spinstoplay", spinstoplay)
				spinstoplay--
			} // spins to play

			// process the category

			//}
			log.Println("Schedule item", categories)
		}
		if schedulerowserr != nil {
			log.Println("Schedule eof", schedulerowserr)
			adjustToTopOfHour()
			getNextHourPart()
		}
		getNextHourPart()
		// Now that the sound finished playing, we can restart from the beginning (or go to any location in the sound) using seek
		// newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
		// if err != nil{
		//     panic("player.Seek failed: " + err.Error())
		// }
		// println("Player is now at position:", newPos)
		// player.Play()

		// If you don't want the player/sound anymore simply close
		/* 	err = player.Close()
		   	if err != nil {
		   		panic("player.Close failed: " + err.Error())
		   	} */
	}
	// get next hour
}
