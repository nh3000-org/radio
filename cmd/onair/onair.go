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
var position uint64
var categories string
var spinstoplay uint64

// inventory
var rowid uint64
var category string
var artist string
var song string
var album string
var length uint64
var expireson string
var lastplayed string
var dateadded string
var spinstoday uint64
var spinsweek uint64
var spinstotal uint64

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
		playingday = "MON"
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

func playit(song uint64, cat string) int {
	elapsed = 0
	fileid = strconv.FormatUint(song, 10)
	// if min of hour is even play intro
	// if min of hour is odd play extro
	if cat == "top40" {
		t := time.Now()
		if t.Minute()%2 == 0 {
			fileid += "INTRO"
		} else {
			fileid += "OUTRO"
		}

	}
	// Read the mp3 file into memory
	fileBytes := config.GetBucket("mp3", fileid)
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

	// Create a new 'player' that will handle our sound. Paused by default.
	player := otoCtx.NewPlayer(decodedMp3)

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
var layout = "2006-01-02 15:04:05"

func main() {
	schedday := flag.String("schedday", "Day: MON TUE WED THU FRI SAT SUN", "-schedday MON")
	stationid := flag.String("stationid", "Call Letter of Station", "-station WRRW")
	schedhour := flag.String("schedhour", "HOUR 00 .. 23", "-schedhour 00")
	logging := flag.String("logging", "TRUE OR FALSE", "-logging TRUE")
	flag.Parse()
	playingday = *schedday
	playinghour = *schedhour

	if *logging == "true" {
		logto = true
	} else {
		logto = false
	}
	var TheDB = "postgresql://" + config.DBuser + ":" + config.DBpassword + "@" + config.DBaddress
	dbConfig, err := pgxpool.ParseConfig(TheDB)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	connPool, err := pgxpool.NewWithConfig(context.Background(), *&dbConfig)
	if err != nil {
		log.Println("Unable to connect to database: ", err)
		os.Exit(1)
	}
	connection, err := connPool.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from the database pool!!")
	}
	defer connection.Release()

	connection.Conn().Prepare(context.Background(), "scheduleget", "select * from schedule where days = $1 and hours = $2 order by position")
	connection.Conn().Prepare(context.Background(), "inventoryresetdaily", "update inventory  set lastplayed = '2024-01-01 00:00:00',spinstoday = 0")
	connection.Conn().Prepare(context.Background(), "inventoryresetweekly", "update inventory  set lastplayed = '2024-01-01 00:00:00',spinsweek = 0")

	connection.Conn().Prepare(context.Background(), "inventorygetschedule", "select * from inventory where category = $1 order by lastplayed desc limit 10")

	connection.Conn().Prepare(context.Background(), "inventoryget", "select * from inventory where id = $1")
	connection.Conn().Prepare(context.Background(), "inventoryupdate", "update inventory set spinstoday = $1, spinsweek = $2, spinstotal = $3, lastplayed = $4, length= $5 where rowid = $6")
	connection.Conn().Prepare(context.Background(), "inventorydelete", "delete inventory where rowid = $1")
	for {
		runtime.GC()
		runtime.ReadMemStats(&memoryStats)
		log.Panicln("Memory start:", strconv.FormatUint(memoryStats.Alloc/1024/1024, 10)+" Mib")

		schedulerows, schedulerowserr := connection.Query(context.Background(), "scheduleget", playinghour, playingday)
		for schedulerows.Next() {

			scheduleerr := schedulerows.Scan(&days, &hours, &position, &categories, &spinstoplay)
			if scheduleerr != nil {
				log.Println("reading schedule " + err.Error())
			}
			if scheduleerr == nil {
				for spinstoplay > 0 {
					// get an inventory item to play
					invrows, invrowserr := connection.Query(context.Background(), "inventorygetschedule", &categories)
					for invrows.Next() {

						inverr := invrows.Scan(&rowid, &category, &artist, &song, &album, &length, &expireson, &lastplayed, &dateadded, &spinstoday, &spinsweek, &spinstotal)
						if inverr != nil {
							log.Println("processing inventory " + inverr.Error())
						}
						// play the item
						itemlength = playit(rowid, categories)
						// update statistics
						spinstoday++
						spinsweek++
						spinstotal++
						lastplayed = time.Now().String()
						_, invupderr := connection.Exec(context.Background(), "inventoryupdate", spinstoday, spinsweek, spinstoday, lastplayed, itemlength, rowid)
						if invupderr != nil {
							log.Println("updating inventory " + invupderr.Error())
						}
						spinstoplay--

						config.SendONAIR(*stationid, " - "+song)
						//oa.Onair.PutString(*stationid, " - "+song)

					}
					// check inventory expired

					if invrowserr != nil {
						log.Println("reading inventory " + invrowserr.Error())
					}

					ex, exerr := time.Parse(layout, expireson)
					if exerr != nil {
						log.Println("reading inventory " + exerr.Error())
					}
					if ex.After(time.Now()) {
						_, invdelerr := connection.Exec(context.Background(), "inventorydelete", rowid)
						if invdelerr != nil {
							log.Println("deleting inventory " + invdelerr.Error())
						}

						fileid = strconv.FormatUint(rowid, 10)
						var intro = fileid + "INTRO"
						var outro = fileid + "OUTRO"
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
				}

				// process the category

			}

		}
		if schedulerowserr != nil {
			adjustToTopOfHour()
			getNextHourPart()
		}
		getNextDay()
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
