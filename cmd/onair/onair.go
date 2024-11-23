package main

import (
	"bytes"
	"context"
	"flag"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ebitengine/oto/v3"
	"github.com/hajimehoshi/go-mp3"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nh3000-org/radio/config"
)

var playingday string  // MON .. SUN
var playinghour string // 00 .. 23
var schedday string    // MON .. SUN
var schedhour string   // 00 .. 23
var logto bool
var moh = 0

func minOfHour() int {
	if logto {
		log.Println("[minOfHour]")
	}
	t := time.Now()
	return t.Minute()
}
func currentHour() int {
	if logto {
		log.Println("[currentHour]")
	}
	t := time.Now()
	return t.Hour()
}
func adjustToTopOfHour() {
	if logto {
		log.Println("[adjustToTopOfHour]", playingday, playinghour)
	}
}
func getNextDay() {
	if playingday == "MON" {
		schedday = "TUE"
	}
	if playingday == "TUE" {
		schedday = "WED"
	}
	if playingday == "WED" {
		schedday = "THU"
	}
	if playingday == "THU" {
		schedday = "FRI"
	}
	if playingday == "FRI" {
		schedday = "SAT"
	}
	if playingday == "SAT" {
		schedday = "SUN"
	}
	if playingday == "SUN" {
		clearSpinsPerDayCount()
		playingday = "MON"
		schedday = "MON"
	}
	playingday = schedday

}
func checkTop40ForIntroOutroPlay() {
	if logto {
		log.Println("[checkTop40ForExtraPlay]", playingday, playinghour)
	}
	// determine what should play
	// if min of hour is even play intro
	// if min of hour is odd play extro
}
func clearSpinsPerDayCount() {
	if logto {
		log.Println("[clearSpinsPerDayCount]")
	}
}
func deleteExpiredInventory(id int) {
	if logto {
		log.Println("[deleteExpiredInventory]", id)
	}
}
func getNextHourPart() string {
	adjustToTopOfHour()
	hp, hperr := strconv.Atoi(playinghour)
	if hperr != nil {
		playinghour = "00"
		schedhour = "00"
		return schedhour
	}
	if hperr == nil {
		hp++
		if hp > 23 {
			playinghour = "00"
			schedhour = "00"
			return schedhour
		}
	}
	newhp := strconv.Itoa(hp)
	if len(newhp) == 1 {
		newhp = "0" + newhp
	}
	return newhp
}
func main() {

	schedday := flag.String("schedday", "Day: MON TUE WED THU FRI SAT SUN", "-schedday MON")
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

	var thedb = config.DBaddress + config.DBname + "?user=" + config.DBuser + "&password=" + config.DBpassword
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	pool, errp := pgxpool.New(ctx, thedb)
	if errp != nil {
		log.Println("Unable to connect to database: ", errp, thedb)
		os.Exit(1)
	}
	Daysget, _ := pgconn.Prepare(ctx, "daysget", "select * from days order by dayofweek")

	Hoursget, _ := conn.Prepare(ctxsql, "hoursget", "select * from hours order by id")
	d.hoursget = Hoursget
	Categoriesget, _ := conn.Prepare(ctxsql, "categoriesget", "select * from categories order by id")
	d.categoriesget = Categoriesget
	Scheduleget, _ := conn.Prepare(ctxsql, "scheduleget", "select * from schedule where days = $1 and hours = $2 order by position")
	d.scheduleget = Scheduleget
	Inventoryreset, _ := conn.Prepare(ctxsql, "inventoryreset", "update inventory  set lastplayed = '2024-01-01 00:00:00',spinstoday = 0,spinsweek = 0")
	d.inventoryreset = Inventoryreset
	Inventorygetschedule, _ := conn.Prepare(ctxsql, "inventorygetschedule", "select * from inventory where category = $1 order by lastplayed desc limit 10")
	d.inventorygetschedule = Inventorygetschedule
	Inventoryget, _ := conn.Prepare(ctxsql, "inventoryget", "select * from inventory where id = $1")
	d.inventoryget = Inventoryget
	Inventoryadd, _ := conn.Prepare(ctxsql, "inventoryadd", "insert insert into inventory (id,category,artist,song,album,length,lastplayed,dateadded,spinstoday,spinsweek,spinstotal) values($1,$2,$3,$4,$5,#6,$7,$8,$9,$10,$11)")
	d.inventoryadd = Inventoryadd
	pool.Exec(scheduleget, *schedday, schedhour)

	// Read the mp3 file into memory
	fileBytes, err := os.ReadFile("./my-file.mp3")
	if err != nil {
		panic("reading my-file.mp3 failed: " + err.Error())
	}

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
		time.Sleep(time.Millisecond)
	}

	// Now that the sound finished playing, we can restart from the beginning (or go to any location in the sound) using seek
	// newPos, err := player.(io.Seeker).Seek(0, io.SeekStart)
	// if err != nil{
	//     panic("player.Seek failed: " + err.Error())
	// }
	// println("Player is now at position:", newPos)
	// player.Play()

	// If you don't want the player/sound anymore simply close
	err = player.Close()
	if err != nil {
		panic("player.Close failed: " + err.Error())
	}
}
