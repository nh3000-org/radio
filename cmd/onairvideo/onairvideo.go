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
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nh3000-org/radio/config"
	//"github.com/nh3000-org/radio/config"
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

var hourtimingstart time.Time
var nextgetconn *pgxpool.Conn
var errnextget error
var nextrows pgx.Rows
var nextrowserr error
var nexterr error

func playNext() {
	nextgetconn, _ = config.SQL.Pool.Acquire(context.Background())
	_, errnextget = nextgetconn.Conn().Prepare(context.Background(), "next", "select * from inventory where category = 'VIDEONEXT'")
	if errnextget != nil {
		log.Println("Prepare nextgetconn", errnextget)
		config.Send("messages.NEXT", "Prepare Next Get "+errnextget.Error(), "onairvideo")
	}
	nextrows, nextrowserr = nextgetconn.Query(context.Background(), "next")
	if nextrowserr != nil {
		config.Send("messages.NEXT", "Prepare Inventory Read "+nextrowserr.Error(), "onairvideo")
		log.Fatal("Error reading inventory ", nextrowserr, " cat: ", categories)
	}

	for nextrows.Next() {
		nexterr = nextrows.Scan(&rowid, &category, &artist, &song, &album, &songlength, &rndorder, &startson, &expireson, &lastplayed, &dateadded, &today, &week, &total, &sourcelink)
		if nexterr != nil {
			log.Println("processing inventory video get " + nexterr.Error())
			config.Send("messages.NEXT", "Inventory Video Get "+nexterr.Error(), "onairvideo")
		}
		// play the item
		config.SendONAIRmp4( artist+" - "+album+" - "+song)
		itemlength = Play(otoctx, rowid, category)

	}
	nextgetconn.Release()
}

var tohtime time.Time
var tohmin float64
var tohleft float64
var tohspinsf float64
var tohspins int
var tohgetconn *pgxpool.Conn
var tohnextget error
var tohrows pgx.Rows
var tohrowserr error

// var toherr error
var toonairvideohinverr error
var tohinvupdconn *pgxpool.Conn
var toherrinventoryupd error

var cspwgetconn *pgxpool.Conn
var cspwerr error

func clearSpinsPerWeekCount() {
	if logto {
		log.Println("[clearSpinsPerWeekCount]")
	}

	cspwgetconn, _ = config.SQL.Pool.Acquire(context.Background())

	_, cspwerr = cspwgetconn.Query(context.Background(), "update inventory set spinsweek = 0")
	if cspwerr != nil {
		log.Println("Clear Spins Per Week Clear "+cspwerr.Error(), "CSPW", playingday, playinghour, categories)
		config.Send("messages."+StationId, "Clear Spins Per Week Clear "+cspwerr.Error(), "onair")
	}
	cspwgetconn.Release()

}

func clearSpinsPerDayCount() {
	if logto {
		log.Println("[clearSpinsPerDayCount]")
	}

	config.ToPDF("SpinsPerDay", StationId)

	cspdgetconn, _ := config.SQL.Pool.Acquire(context.Background())

	_, cspderr := cspdgetconn.Query(context.Background(), "update inventory set spinstoday = 0")
	if cspderr != nil {
		log.Println("Clear Spins Per Day Clear "+cspderr.Error(), "CSPD", playingday, playinghour, categories)
		config.Send("messages."+StationId, "Clear Spins Per Day Clear "+cspderr.Error(), "onair")
	}
	cspdgetconn.Release()

}

var hp int
var hperr error

var elapsed = 0
var fileid string
var otoCtx *oto.Context
var otoreadyChan chan struct{}
var otoerr error

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

	otoCtx, otoreadyChan, otoerr = oto.NewContext(op)
	if otoerr != nil {
		panic("oto.NewContext failed: " + otoerr.Error())
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-otoreadyChan

	return *otoCtx

}

var fileBytes []byte
var fileBytesReader *bytes.Reader
var t time.Time
var decodedMp3 *mp3.Decoder
var decodedMp3err error
var player *oto.Player

func Play(ctx oto.Context, song string, cat string) int {

	elapsed = 0

	if cat == "top40" {
		t = time.Now()
		if t.Minute()%2 == 0 {
			song += "INTRO"
		} else {
			song += "OUTRO"
		}

	}
	// Read the mp3 file into memory
	fileBytes = config.GetBucket("mp3", song, StationId)
	/* 	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	} */

	// Convert the pure bytes into a reader object that can be used with the mp3 decoder
	fileBytesReader = bytes.NewReader(fileBytes)

	// Decode file
	decodedMp3, decodedMp3err = mp3.NewDecoder(fileBytesReader)
	if decodedMp3err != nil {
		config.Send("messages."+StationId, "MP3 Decoder Error "+song, "onair")
		log.Println("mp3.NewDecoder failed: ", decodedMp3err.Error(), "for song:", song)
		return 0
	}
	// Create a new 'player' that will handle our sound. Paused by default.
	player = ctx.NewPlayer(decodedMp3)

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
var StationId = ""
var spinstoplay int
var spinstoplayerr error
var invupdconn *pgxpool.Conn
var trafficaddconn *pgxpool.Conn
var invdelconn *pgxpool.Conn
var errinventorygetschedule error
var errinventoryupd error
var invupderr error
var trafficaddconnerr error
var errtrafficadd error
var trafficadderr error
var errinventorydelete error
var invdelerr error

// var invrows pgx.Row
// var invrowserr error
var spinsweek int
var spinstoday int
var spinstotal int
var lp time.Time
var played = "YYYY-MM-DD HH:mm:SS"
var month string
var day string
var min string
var sec string
var ex time.Time
var exerr error
var intro string
var outro string
var errremove error
var errremovei error
var errremoveo error
var otoctx oto.Context

func main() {
	hourtimingstart = time.Now()
	//schedDay := flag.String("schedday", "MON", "-schedday MON || TUE || WED || THU || FRI || SAT || SUN")
	stationId := flag.String("stationid", "WRRW", "-station WRRW")
	StationId = *stationId
	Logging := flag.String("logging", "true", "-logging true || false")
	flag.Parse()

	pm := time.Now().Minute()
	log.Println("TEST day hour minute station logging", playingday, playinghour, pm, *stationId, *Logging)

	otoctx = playsetup()

	if *Logging == "true" {
		logto = true
	} else {
		logto = false
	}
	log.Println("Startup Parms:",  *stationId, *Logging)

	config.NewPGSQL()
	config.NewNatsJS()
	//config.NewNatsJSOnAirmp4()

	var connectionspool *pgxpool.Conn
	var connectionspoolerr error
	var errscheduleget error
	var schedulerows pgx.Rows
	var schedulerowserr error
	var invgetconn *pgxpool.Conn
	var invrows pgx.Rows
	var invrowserr error
	var inverr error
	//clearSpinsPerDayCount()

	// toh to get in sync
	//adjustToTopOfHour()
	for {

		runtime.GC()
		runtime.ReadMemStats(&memoryStats)
		log.Println("Memory start: day:", playingday, ":hour:", playinghour, ":mem:", strconv.FormatUint(memoryStats.Alloc/1024/1024, 10)+" Mib")
		connectionspool, connectionspoolerr = config.SQL.Pool.Acquire(context.Background())
		if connectionspoolerr != nil {
			config.Send("messages."+*stationId, "Connection Pool Acquire FATAL "+connectionspoolerr.Error(), "onairvideo")
			log.Fatal("Error while acquiring connection from the database pool!!")
		}
		_, errscheduleget = connectionspool.Conn().Prepare(context.Background(), "scheduleget", "select * from schedule where days = 'VID' and hours = '00' order by position")
		if errscheduleget != nil { // get an inventory item to play
			config.Send("messages."+*stationId, "Prepare Schedule Get FATAL "+errscheduleget.Error(), "onairvideo")
			log.Fatal("Prepare scheduleget", errscheduleget)
		}
		schedulerows, schedulerowserr = connectionspool.Query(context.Background(), "scheduleget", playingday, playinghour)
		//log.Println("reading schedule next ", playingday, playinghour, categories)
		for schedulerows.Next() {
			runtime.GC()
			runtime.ReadMemStats(&memoryStats)
			//log.Println("Memory cat:", categories, ":day:", playingday, ":hour:", playinghour, ":mem:", strconv.FormatUint(memoryStats.Alloc/1024/1024, 10)+" Mib")
1212
			scheduleerr := schedulerows.Scan(&rowid, &days, &hours, &position, &categories, &toplay)
			//log.Println("reading schedule: ", days, hours, position, categories, toplay, " schedule", playingday, playinghour, categories)
			spinstoplay, spinstoplayerr = strconv.Atoi(toplay)
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
				runtime.GC()
				runtime.ReadMemStats(&memoryStats)
				log.Println("CAT:", categories, ":spins:", spinstoplay, ":day:", playingday, ":hour:", playinghour, ":mem:", strconv.FormatUint(memoryStats.Alloc/1024/1024, 10)+" Mib")
				if spinstoplay <= 0 {
					break
				}
				// check for NEXT
				playNext()
				invgetconn, _ = config.SQL.Pool.Acquire(context.Background())
				_, errinventorygetschedule = invgetconn.Conn().Prepare(context.Background(), "inventorygetschedule", "select * from inventory where category = $1 order by lastplayed, rndorder limit 10")
				if errinventorygetschedule != nil {
					log.Println("Prepare inventorygetschedule", errinventorygetschedule)
					config.Send("messages."+*stationId, "Prepare Inventory Get "+errinventorygetschedule.Error(), "onair")
				}
				invrows, invrowserr = invgetconn.Query(context.Background(), "inventorygetschedule", categories)
				if invrowserr != nil {
					config.Send("messages."+*stationId, "Prepare Inventory Read "+invrowserr.Error(), "onair")
					log.Fatal("Error reading inventory ", invrowserr, " cat: ", categories)
				}

				for invrows.Next() {
					if spinstoplay <= 0 {
						break
					}
					inverr = invrows.Scan(&rowid, &category, &artist, &song, &album, &songlength, &rndorder, &startson, &expireson, &lastplayed, &dateadded, &today, &week, &total, &sourcelink)
					//log.Println("processing inventory song get"+song, " schedule", playingday, playinghour, categories)
					if inverr != nil {
						log.Println("processing inventory video get " + inverr.Error())
						config.Send("messages."+*stationId, "Inventory Video Get "+inverr.Error(), "onair")
					}
					// play the item
					config.SendONAIRmp4("onairmp4",*stationId, artist+" - "+album+" - "+song)
					itemlength = Play(otoctx, rowid, category)
					// update statistics
					spinsweek, _ = strconv.Atoi(week)
					spinsweek++
					spinstoday, _ = strconv.Atoi(today)
					spinstoday++
					spinstotal, _ = strconv.Atoi(total)
					spinstotal++
					lp = time.Now()

					played = strings.Replace(played, "YYYY", strconv.Itoa(lp.Year()), 1)
					month = strconv.Itoa(int(lp.Month()))
					if len(month) == 1 {
						month = "0" + month
					}
					played = strings.Replace(played, "MM", month, 1)
					day = strconv.Itoa(int(lp.Day()))
					if len(day) == 1 {
						day = "0" + day
					}
					played = strings.Replace(played, "DD", day, 1)

					hours = strconv.Itoa(int(lp.Hour()))
					if len(hours) == 1 {
						hours = "0" + hours
					}
					played = strings.Replace(played, "HH", hours, 1)

					min = strconv.Itoa(int(lp.Minute()))
					if len(min) == 1 {
						min = "0" + min
					}
					played = strings.Replace(played, "mm", min, 1)

					sec = strconv.Itoa(int(lp.Second()))
					if len(sec) == 1 {
						sec = "0" + sec
					}
					played = strings.Replace(played, "SS", sec, 1)

					//log.Println("last played", played, " schedule", playingday, playinghour, categories)
					invupdconn, _ = config.SQL.Pool.Acquire(context.Background())
					_, errinventoryupd = invupdconn.Conn().Prepare(context.Background(), "inventoryupdate", "update inventory set spinstoday = $1, spinsweek = $2, spinstotal = $3, lastplayed = $4, songlength= $5 where rowid = $6")
					if errinventoryupd != nil {
						log.Println("Prepare inventory upd", errinventoryupd, " schedule", playingday, playinghour, categories)
						config.Send("messages."+*stationId, "Prepare Inventory Update "+errinventorygetschedule.Error(), "onair")
					}
					_, invupderr = invupdconn.Exec(context.Background(), "inventoryupdate", spinstoday, spinsweek, spinstotal, played, itemlength, rowid)
					if invupderr != nil {
						log.Println("updating inventory "+invupderr.Error(), " schedule", playingday, playinghour, categories)
						config.Send("messages."+*stationId, "Inventory Update "+invupderr.Error(), "onair")
					}
					invupdconn.Release()
					if strings.HasPrefix(category, "ADDS") {
						//log.Println("adding inventory to traffic", song)
						trafficaddconn, trafficaddconnerr = config.SQL.Pool.Acquire(context.Background())
						if trafficaddconnerr != nil {
							log.Println("Prepare trafficadd", trafficaddconnerr)
							config.Send("messages."+*stationId, "Prepare trafficadd conn "+trafficaddconnerr.Error(), "onair")

						}
						_, errtrafficadd = trafficaddconn.Conn().Prepare(context.Background(), "trafficadd", "insert into  traffic (artist, album,song,playedon) values($1,$2,$3,$4)")
						if errtrafficadd != nil {
							log.Println("Prepare trafficadd", errtrafficadd)
							config.Send("messages."+*stationId, "Prepare trafficadd "+errtrafficadd.Error(), "onair")
						}
						//log.Println("adding inventory to traffic adding", song)
						_, trafficadderr = trafficaddconn.Exec(context.Background(), "trafficadd", artist, song, album, played)
						if trafficadderr != nil {
							log.Println("updating inventory " + trafficadderr.Error())
							config.Send("messages."+*stationId, "Updating Inventory "+trafficadderr.Error(), "onair")
						}
						trafficaddconn.Release()
					}
					expireson = strings.Replace(expireson, " ", "T", 1)
					expireson = expireson + "Z"
					if !strings.HasPrefix(expireson, "9999") {
						log.Println("Expires on", expireson)
					}
					ex, exerr = time.Parse(time.RFC3339, expireson)
					if exerr != nil {
						log.Println("inventory time parse "+exerr.Error(), " schedule", playingday, playinghour, categories)
						config.Send("messages."+*stationId, "Inventory Time Parse "+exerr.Error(), "onair")
					}
					//log.Println("EXPIRES: ", ex.String())
					if time.Now().After(ex) {
						log.Println("deleting  expired inventory: ", fileid)
						invdelconn, _ = config.SQL.Pool.Acquire(context.Background())
						_, errinventorydelete = invdelconn.Conn().Prepare(context.Background(), "inventorydelete", "delete from inventory where rowid = $1")
						if errinventorydelete != nil {
							log.Println("Prepare inventory delete", errinventorydelete)
							config.Send("messages."+*stationId, "Prepage Indentory Delete "+errinventorydelete.Error(), "onair")
						}

						_, invdelerr = invdelconn.Exec(context.Background(), "inventorydelete", rowid)
						if invdelerr != nil {
							log.Println("deleting inventory " + invdelerr.Error())
							config.Send("messages."+*stationId, "Indentory Delete "+invdelerr.Error(), "onair")
						}
						invdelconn.Release()

						//fileid = strconv.FormatUint(rowid, 10)
						intro = rowid + "INTRO"
						outro = rowid + "OUTRO"
						errremove = config.DeleteBucket("mp3s", fileid)
						if errremove != nil {
							log.Println("deleting  failed: ", errremove.Error(), fileid)
							config.Send("messages."+*stationId, "MP3 Bucket Delete "+fileid+" "+errremove.Error(), "onair")
						}
						errremovei = config.DeleteBucket("mp3s", intro)
						if errremovei != nil {
							log.Println("deleting  failed: ", errremovei.Error(), intro)
							config.Send("messages."+*stationId, "MP3 Bucket Delete "+intro+" "+errremove.Error(), "onair")
						}
						errremoveo = config.DeleteBucket("mp3s", outro)
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
					//log.Println("spinstoplay inventory rows", spinstoplay, " schedule", playingday, playinghour, categories)
					spinstoplay--
				} // inventory rows
				spinstoplay = 0
				invgetconn.Release()
				break

			} // spins to play
			connectionspool.Release()
		}
		if schedulerowserr != nil {
			log.Println("Schedule eof", schedulerowserr, " schedule", playingday, playinghour, categories)
			config.Send("messages."+*stationId, "Prepare Schedule Get rows error "+schedulerowserr.Error(), "onair")
		}
		schedulerows.Close()
		//adjustToTopOfHour()
		//getNextHourPart()

	}

}
