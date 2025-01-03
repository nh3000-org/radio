package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
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
var startson string
var expireson string
var lastplayed string
var dateadded string
var today string
var week string
var total string
var toplay string
var sourcelink string
var playcount int

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
		playinghour = "00"
		schedhour = "00"
	}
	if playingday == "TUE" {
		schedday = "WED"
		clearSpinsPerDayCount()
		playinghour = "00"
		schedhour = "00"
	}
	if playingday == "WED" {
		schedday = "THU"
		clearSpinsPerDayCount()
		playinghour = "00"
		schedhour = "00"
	}
	if playingday == "THU" {
		schedday = "FRI"
		clearSpinsPerDayCount()
		playinghour = "00"
		schedhour = "00"
	}
	if playingday == "FRI" {

		schedday = "SAT"
		clearSpinsPerDayCount()
		playinghour = "00"
		schedhour = "00"
	}
	if playingday == "SAT" {
		schedday = "SUN"
		clearSpinsPerDayCount()
		playinghour = "00"
		schedhour = "00"
	}
	if playingday == "SUN" {
		clearSpinsPerWeekCount()
		schedday = "MON"
		playinghour = "00"
		schedhour = "00"
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

	return *otoCtx

}
func Play(ctx oto.Context, song string, cat string) int {
	playcount++
	elapsed = 0
	log.Println("Playit count:", playcount, ":", song, ":", cat, ":")
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

	sql, sqlerr := config.NewPGSQL()
	if sqlerr != nil {
		log.Fatal("Could not connect to DB")
	}

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
		connectionspool, connectionspoolerr := sql.Pool.Acquire(context.Background())
		if connectionspoolerr != nil {
			config.Send("messages."+*stationId, "Connection Pool Acquire FATAL "+connectionspoolerr.Error(), "onair")
			log.Fatal("Error while acquiring connection from the database pool!!")
		}
		_, errscheduleget := connectionspool.Conn().Prepare(context.Background(), "scheduleget", "select * from schedule where days = $1 and hours = $2 order by position")
		if errscheduleget != nil {
			config.Send("messages."+*stationId, "Prepare Schedule Get FATAL "+errscheduleget.Error(), "onair")
			log.Panicln("Prepare scheduleget", errscheduleget)
		}
		schedulerows, schedulerowserr := connectionspool.Query(context.Background(), "scheduleget", playingday, playinghour)
		log.Println("reading schedule next ", playingday, playinghour, categories)
		for schedulerows.Next() {

			scheduleerr := schedulerows.Scan(&rowid, &days, &hours, &position, &categories, &toplay)
			log.Println("reading schedule: ", days, hours, position, categories, toplay, " schedule", playingday, playinghour, categories)
			spinstoplay, spinstoplayerr := strconv.Atoi(toplay)
			if spinstoplayerr != nil {
				config.Send("messages."+*stationId, "Schedule spinstoplayerr "+spinstoplayerr.Error(), "onair")
				log.Panicln("Error spinstoplayerr "+scheduleerr.Error(), " schedule", playingday, playinghour, categories)
			}
			if scheduleerr != nil {
				config.Send("messages."+*stationId, "Schedule Get "+scheduleerr.Error(), "onair")
				log.Panicln("Error scheduleerr "+scheduleerr.Error(), " schedule", playingday, playinghour, categories)
			}
			//if scheduleerr == nil {
			for spinstoplay > 0 {
				log.Println("START FOR SPINSTOPLAY ", spinstoplay)
				if spinstoplay <= 0 {
					break
				}
				// get an inventory item to play
				invgetconn, _ := sql.Pool.Acquire(context.Background())
				_, errinventorygetschedule := invgetconn.Conn().Prepare(context.Background(), "inventorygetschedule", "select * from inventory where category = $1 order by lastplayed, rndorder limit 10")
				if errinventorygetschedule != nil {
					log.Println("Prepare inventorygetschedule", errinventorygetschedule)
					config.Send("messages."+*stationId, "Prepare Inventory Get "+errinventorygetschedule.Error(), "onair")
				}
				invrows, invrowserr := invgetconn.Query(context.Background(), "inventorygetschedule", categories)
				if invrowserr != nil {
					config.Send("messages."+*stationId, "Prepare Inventory Read "+invrowserr.Error(), "onair")
					log.Fatal("Error reading inventory ", invrowserr, " cat: ", categories)
				}
				log.Println("inventory schedule get ", categories, "err", invrowserr)
				/* 				var stp int64 = int64(spinstoplay)
				   				if invrows.CommandTag().RowsAffected() < stp {
				   					log.Println("adjusting spins to play rowsaffected", invrows.CommandTag().RowsAffected(), "spins", spinstoplay)
				   					spinstoplay = int(invrows.CommandTag().RowsAffected())

				   				} */
				for invrows.Next() {
					if spinstoplay <= 0 {
						break
					}
					inverr := invrows.Scan(&rowid, &category, &artist, &song, &album, &songlength, &rndorder, &startson, &expireson, &lastplayed, &dateadded, &today, &week, &total, &sourcelink)
					log.Println("processing inventory song get"+song, " schedule", playingday, playinghour, categories)
					if inverr != nil {
						log.Println("processing inventory song get " + inverr.Error())
						config.Send("messages."+*stationId, "Inventory Song Get "+inverr.Error(), "onair")
					}
					// play the item
					config.SendONAIR(*stationId, artist+" - "+album+" - "+song)
					itemlength = Play(otoctx, rowid, category)
					// update statistics
					spinsweek, _ := strconv.Atoi(week)
					spinsweek++
					spinstoday, _ := strconv.Atoi(today)
					spinstoday++
					spinstotal, _ := strconv.Atoi(total)
					spinstotal++
					lp := time.Now()
					played := "YYYY-MM-DD HH:mm:SS"
					played = strings.Replace(played, "YYYY", strconv.Itoa(lp.Year()), 1)
					m := strconv.Itoa(int(lp.Month()))
					if len(m) == 1 {
						m = "0" + m
					}
					played = strings.Replace(played, "MM", m, 1)
					d := strconv.Itoa(int(lp.Day()))
					if len(d) == 1 {
						d = "0" + d
					}
					played = strings.Replace(played, "DD", d, 1)

					h := strconv.Itoa(int(lp.Hour()))
					if len(h) == 1 {
						d = "0" + d
					}
					played = strings.Replace(played, "HH", h, 1)

					min := strconv.Itoa(int(lp.Minute()))
					if len(min) == 1 {
						d = "0" + d
					}
					played = strings.Replace(played, "mm", min, 1)

					sec := strconv.Itoa(int(lp.Second()))
					if len(sec) == 1 {
						sec = "0" + d
					}
					played = strings.Replace(played, "SS", sec, 1)

					log.Println("last played", played, " schedule", playingday, playinghour, categories)
					invupdconn, _ := sql.Pool.Acquire(context.Background())
					_, errinventoryupd := invupdconn.Conn().Prepare(context.Background(), "inventoryupdate", "update inventory set spinstoday = $1, spinsweek = $2, spinstotal = $3, lastplayed = $4, songlength= $5 where rowid = $6")
					if errinventoryupd != nil {
						log.Println("Prepare inventory upd", errinventoryupd, " schedule", playingday, playinghour, categories)
						config.Send("messages."+*stationId, "Prepare Inventory Update "+errinventorygetschedule.Error(), "onair")
					}
					_, invupderr := invupdconn.Exec(context.Background(), "inventoryupdate", spinstoday, spinsweek, spinstoday, played, itemlength, rowid)
					if invupderr != nil {
						log.Println("updating inventory "+invupderr.Error(), " schedule", playingday, playinghour, categories)
						config.Send("messages."+*stationId, "Inventory Update "+invupderr.Error(), "onair")
					}

					if strings.HasPrefix(category, "ADDS") {
						log.Println("adding inventory to traffic", song)
						trafficaddconn, trafficaddconnerr := sql.Pool.Acquire(context.Background())
						if trafficaddconnerr != nil {
							log.Println("Prepare trafficadd", trafficaddconnerr)
							config.Send("messages."+*stationId, "Prepare trafficadd conn "+trafficaddconnerr.Error(), "onair")

						}
						_, errtrafficadd := trafficaddconn.Conn().Prepare(context.Background(), "trafficadd", "insert into  traffic (artist, album,song,playedon) values($1,$2,$3,$4)")
						if errtrafficadd != nil {
							log.Println("Prepare trafficadd", errtrafficadd)
							config.Send("messages."+*stationId, "Prepare trafficadd "+errtrafficadd.Error(), "onair")
						}
						log.Println("adding inventory to traffic adding", song)
						_, trafficadderr := trafficaddconn.Exec(context.Background(), "trafficadd", artist, song, album, played)
						if trafficadderr != nil {
							log.Println("updating inventory " + trafficadderr.Error())
							config.Send("messages."+*stationId, "Updating Inventory "+trafficadderr.Error(), "onair")
						}
						trafficaddconn.Release()
					}
					expireson = strings.Replace(expireson, " ", "T", 1)
					expireson = expireson + "Z"
					log.Println("Expires on", expireson)
					ex, exerr := time.Parse(time.RFC3339, expireson)
					if exerr != nil {
						log.Println("inventory time parse "+exerr.Error(), " schedule", playingday, playinghour, categories)
						config.Send("messages."+*stationId, "Inventory Time Parse "+exerr.Error(), "onair")
					}
					log.Println("EXPIRES: ", ex.String())
					if time.Now().After(ex) {
						log.Println("deleting  expired inventory: ", fileid)
						invdelconn, _ := sql.Pool.Acquire(context.Background())
						_, errinventorydelete := invdelconn.Conn().Prepare(context.Background(), "inventorydelete", "delete from inventory where rowid = $1")
						if errinventorydelete != nil {
							log.Println("Prepare inventory delete", errinventorydelete)
							config.Send("messages."+*stationId, "Prepage Indentory Delete "+errinventorydelete.Error(), "onair")
						}

						_, invdelerr := invdelconn.Exec(context.Background(), "inventorydelete", rowid)
						if invdelerr != nil {
							log.Println("deleting inventory " + invdelerr.Error())
							config.Send("messages."+*stationId, "Indentory Delete "+invdelerr.Error(), "onair")
						}
						invdelconn.Release()

						//fileid = strconv.FormatUint(rowid, 10)
						var intro = rowid + "INTRO"
						var outro = rowid + "OUTRO"
						errremove := config.DeleteBucket("mp3s", fileid)
						if errremove != nil {
							log.Println("deleting  failed: ", errremove.Error(), fileid)
							config.Send("messages."+*stationId, "MP3 Bucket Delete "+fileid+" "+errremove.Error(), "onair")
						}
						errremovei := config.DeleteBucket("mp3s", intro)
						if errremovei != nil {
							log.Println("deleting  failed: ", errremovei.Error(), intro)
							config.Send("messages."+*stationId, "MP3 Bucket Delete "+intro+" "+errremove.Error(), "onair")
						}
						errremoveo := config.DeleteBucket("mp3s", outro)
						if errremoveo != nil {
							log.Println("deleting  failed: ", errremoveo.Error(), outro)
							config.Send("messages."+*stationId, "MP3 Bucket Delete "+outro+" "+errremove.Error(), "onair")
						}

					}

					if invrowserr != nil {
						log.Println("reading inventory "+invrowserr.Error(), " schedule", playingday, playinghour, categories)
						config.Send("messages."+*stationId, "Inventory Read "+invrowserr.Error(), "onair")
					}
					//conninv.Release()
					log.Println("spinstoplay inventory rows", spinstoplay, " schedule", playingday, playinghour, categories)
					spinstoplay--
				} // inventory rows
				spinstoplay = 0
				invgetconn.Release()
				break

			} // spins to play

			// process the category

			//}
			//getNextHourPart()
			log.Println("Schedule item", categories, " schedule", playingday, playinghour, categories)
		}
		if schedulerowserr != nil {
			log.Println("Schedule eof", schedulerowserr, " schedule", playingday, playinghour, categories)
			config.Send("messages."+*stationId, "Prepare Schedule Get rows error "+schedulerowserr.Error(), "onair")
		}
		schedulerows.Close()
		adjustToTopOfHour()
		getNextHourPart()

	}

}
