package config

import (
	"context"
	"log"
	"os"
	"strconv"
	"strings"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2"
)

/*
* posgresql support only
* hostssl  all  all  0.0.0.0/0  cert in pg_hba.conf
* psql "sslmode=require host=localhost dbname=test"
 */
//var conn *pgx.Conn

// var DBaddress = "localhost:5432/radio"
var DBaddress = "db.newhorizons3000.org:5432/radio?sslmode=verify-ca"
var DBuser = "postgres"
var DBpassword = "postgres"

//var dburlpropocol = "postgresql://localhost:5432/radio?user=root&password=password"

type SQLconn struct {
	//conn   pgx.Conn
	Pool   pgxpool.Pool
	Ctx    context.Context
	Ctxcan context.CancelFunc
}

var myerror error

var DBClientcert = "-----BEGIN CERTIFICATE-----\nMIICyTCCAbECFCJlOZ058bh90IyWT+Z+VS+3K42pMA0GCSqGSIb3DQEBCwUAMCEx\nHzAdBgNVBAMMFmRiLm5ld2hvcml6b25zMzAwMC5vcmcwHhcNMjQxMjA2MTI1MjA1\nWhcNMzQxMjA0MTI1MjA1WjAhMR8wHQYDVQQDDBZkYi5uZXdob3Jpem9uczMwMDAu\nb3JnMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsuUB/1oGzIHbZ3Jj\nRdzqZ3mU2phrwuWrwW02DSbrSqhdrWkhwEDZqcNPclAgU45cP9I0ZS9gMYY+pcv6\naqC5XVZ32kHdLSKW0UGWliyDCU3X7fQEighvutCQqktph0EpY7578WrA7uC8Zap6\nT/zQI9hMf1+YkepfH8oB9m9ekXA/wc0Bf9dHNFzlCl6+UieEL/jUTs4TcOfa/YGK\nOyODoyNmmTTLlf9nr8HwZ5MEHflh2v92AXMG0oB0dD+iNmeOvj9Zq/Hw5QzrkFGZ\nSCp7jxiTJkH/9h+63hECk35ttOFqa0v+ccmTPdrFKOeMc1thqXjQ/k4W/AgFpvuD\nJq6kNQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAaU7ZUC+bscXNnWLjG3YMoz9RW\n/OpgK6i634PXOUyMGmT6985PBTOPaVsA4FNXDVAjhrbtCZhavzuD+NYluTW9IR5c\nZKWFILTLpzkoEGSphlWIm6I7m60mbYugl1I7bZZX3vGM7IZkvu9OtGXjuESWUSwm\nYmYUSJG0bQ2gMa7DkZUXtLbVT3ZGpQrYW3gVA9LqAvCCRnC9YlN/wmWZHjyNPp8s\nDnSyV9HEtRdk3jTQR7ocUOOX5vVXTEG5K84cgk8DlwwQlIZ/WQVFHsV7eqDh4Lfo\nQLuVg4VCt45Rbo4mVkyY/FX/13guFV2tM04wGU+2WBgIFQvZYWtqh2O4N0zB\n-----END CERTIFICATE-----\n"
var DBClientkey = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCy5QH/WgbMgdtn\ncmNF3OpneZTamGvC5avBbTYNJutKqF2taSHAQNmpw09yUCBTjlw/0jRlL2Axhj6l\ny/pqoLldVnfaQd0tIpbRQZaWLIMJTdft9ASKCG+60JCqS2mHQSljvnvxasDu4Lxl\nqnpP/NAj2Ex/X5iR6l8fygH2b16RcD/BzQF/10c0XOUKXr5SJ4Qv+NROzhNw59r9\ngYo7I4OjI2aZNMuV/2evwfBnkwQd+WHa/3YBcwbSgHR0P6I2Z46+P1mr8fDlDOuQ\nUZlIKnuPGJMmQf/2H7reEQKTfm204WprS/5xyZM92sUo54xzW2GpeND+Thb8CAWm\n+4MmrqQ1AgMBAAECggEAAmq5z0QdWuZ3OyDOarYhxuzw1U+MBymjW304/COrjUo/\nMmr88n69zL9vg5fFtAifQkUkrD2gCFNBZevepyb4PM9WAQbKRkQ/yDXMDcfYq7Js\n9Cag6BIrRkQalukRaisidBnxopDtaId0ltAJ5SpBJUwpuRUmzR6Jk0xJ9f/KWzR5\n/gofzpJmlLl4Ymfx4/JAhJFyfUOfq1/1BIIXBEUxhCFDX8fJAog1eJD4i7N2DfrY\nIj2PK15WOizweKnigRGXFghh3dnwLNytm18QW5a/CdGOjCuf7uPPuNfnoW7rZok8\nr7geQh+FeXt3Wbqg3Kxc9Da+IRRix7R8VLAvqNWEcQKBgQDapzQCUkopOBSEIz1/\nQeuvfJLlpDqXtYZ7/KwRPcaBzZfpq17X7SpJG/HDaVm9+RRJxF1f9ft5SBxiGet4\nb1h0uCG3NbTUJXTAk1PmlyxuAQ9Xzd+YlLI+U0ih468Q4tpZrd4kUgtc97uzN7t0\nR0RbeeU5ZGcEfsOefVjK4HNe0QKBgQDRc1OgQz+JvxGhGevjdwa2tRhI0EZaaQOg\nvfc4ppBRLdE3F3j2AFPIBigWPyfHsO9fn5nAYxJIHzI6B7o/6M/JJw/C69/cQg7p\np5/VV73m10dcbtzR9/5O1sl+WGHOdgtvM5hP+IiGsvKRgx3N7F5B/UJkZ4R4scOD\nX/BYzr7wJQKBgApTcSJW7oepzVY8L9BNtaqw8GMF8XpuqS47zYh26WQB6JWxcSYz\nXhbbyfwXgpR1Kd8d9ebtP/YHUMfVP4iNgZjphTYYxDRsnGnny0ONihyb0jSsVU3o\nX86PslNq5D6g5/zqOB5w/XZjgKrDDAg+wVyskgW21yKgNe7LLqFOHkSxAoGBAKWa\nQq9/HDikCqNO5HRXwsYpD0da7ZVEXKr2KAbxoz+cM0QU2f3fKl8HhyB31NMNsWXw\nwdccPfMqP0Mkov0u7UMFEHA0oS38SOAzOausESj4Y6LQwOV+5+Kb7npoFQTxzn6g\n07e/MOsXh7THb4RGdAxG2vyZ4GKxYn14GIePB+bFAoGAKzqQzkb5VP5ML5QwGAmM\n2vn8SU1brebCnnRDSaiBvagLZNeA0X0vVqUWSJ29uKO52x2QpOnBPATJ0th+mHOZ\n/APTv+8hrBkARuu5Kplvfl+MLXbZH+iYWTPbD74kVaJk+Km8tePglR919LvZ8Ysl\nL/QCVUYaP4f/REVsnoShzpg=\n-----END PRIVATE KEY-----\n"

func NewPGSQL() error {

	var d = new(SQLconn)
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 2048*time.Hour)

	d.Ctxcan = ctxsqlcan
	d.Ctx = ctxsql
	pooltime, pterr := time.ParseDuration("4096h")
	if pterr != nil {
		log.Fatal("DB001 Failed to parse time: ", pterr)
		return pterr
	}
	var TheDB = "postgresql://" + DBuser + ":" + DBpassword + "@" + DBaddress
	//var thedb = DBaddress + DBname + "?user=" + DBuser + "&password=" + DBpassword
	mydb, mydberr := pgxpool.ParseConfig(TheDB)
	mydb.MaxConnIdleTime = pooltime
	mydb.MaxConns = 50
	mydb.MaxConnLifetime = pooltime
	if mydberr != nil {
		log.Fatal("DB002 Unable to connect to parse config database: ", mydberr)
		return mydberr
	}

	mypool, mypoolerr := pgxpool.NewWithConfig(ctxsql, mydb)
	if mypoolerr != nil {
		log.Fatal("DB003 Unable to create connection pool: ", myerror)
		return mypoolerr
	}

	d.Pool = *mypool
	SQL = d
	return nil
}

type DaysStruct struct {
	Row  int    // rowid
	Day  string // message id
	Desc string // alias
	Dow  int    // hostname
}

var DaysStore = make(map[int]DaysStruct)
var SelectedDay int

func DaysGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	DaysStore = make(map[int]DaysStruct)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	rows, rowserr := conn.Query(ctxsql, "select * from days order by dayofweek")
	var rowid int
	var day string
	var desc string
	var dow int
	for rows.Next() {

		err := rows.Scan(&rowid, &day, &desc, &dow)
		if err != nil {
			log.Println("DB003 GetDays row", err)
		}
		ds := DaysStruct{}
		ds.Row = rowid
		ds.Day = day
		ds.Desc = desc
		ds.Dow = dow
		DaysStore[len(DaysStore)] = ds
		//log.Println("GetDays Got", day, desc)
	}
	if rowserr != nil {
		log.Println("DB004 GetDays row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()

}
func DaysDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "delete from days where rowid =$1", row)

	if rowserr != nil {
		log.Println("DB005 Delete Days row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func DaysUpdate(row int, day string, desc string, dow int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Exec(ctxsql, "update days set id =$1, description = $2, dayofweek = $3 where rowid = $4", day, desc, dow, row)

	if rowserr != nil {
		log.Println("DB006 Update Days row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func DaysAdd(day string, desc string, dow int) {

	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "insert into  days (id, description, dayofweek) values($1,$2,$3)", day, desc, dow)

	if rowserr != nil {
		log.Println("DB007 Add Days row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}

type HoursStruct struct {
	Row  int    // rowid
	Id   string // hours id
	Desc string // alias
}

var HoursStore = make(map[int]HoursStruct)
var SelectedHour int

func HoursGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	HoursStore = make(map[int]HoursStruct)
	rows, rowserr := conn.Query(ctxsql, "select * from hours order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("DB008 GetHours row", err)
		}
		ds := HoursStruct{}
		ds.Row = rowid
		ds.Id = id
		ds.Desc = desc

		HoursStore[len(HoursStore)] = ds

	}
	if rowserr != nil {
		log.Println("DB009 Gethours row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()

}
func HoursDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "delete from hours where rowid =$1", row)

	if rowserr != nil {
		log.Println("DB010 Delete Hours row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func HoursUpdate(row int, id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Exec(ctxsql, "update hours set id =$1, description = $2 where rowid = $3", id, desc, row)

	if rowserr != nil {
		log.Println("DB011 Update Hours row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func HoursAdd(id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "insert into  hours (id, description) values($1,$2)", id, desc)

	if rowserr != nil {
		log.Println("DB012 Add Hours row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}

type CategoriesStruct struct {
	Row  int    // rowid
	Id   string // hours id
	Desc string // alias
}

var CategoriesStore = make(map[int]CategoriesStruct)
var SelectedCategory int

func CategoriesGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	CategoriesStore = make(map[int]CategoriesStruct)
	rows, rowserr := conn.Query(ctxsql, "select * from categories order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("DB013 Get Categories row", err)
		}
		ds := CategoriesStruct{}
		ds.Row = rowid
		ds.Id = id
		ds.Desc = desc

		CategoriesStore[len(CategoriesStore)] = ds

	}
	if rowserr != nil {
		log.Println("DB014 Get Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()

}

var instructions = "Radio Stub Instructions\nBrowse to this file to initiate import\nSongs are identified by ARTIST-SONG-ALBUM.mp3 and ARTIST-SONG-ALBUM-INTRO.mp3 and ARTIST-SONG-ALBUM-OUTRO.mp3 where INTRO and OUTRO are for TOP40 anouncements in the following categories\nADDS, ADDSDRIVETIME and ADDSTOH are used to add advertising to system.\nFILLTOTOH is a phantom category used internally\nIMAGINGID is used to hold artist station plugs\nLIVE is phantom category to indicate live segments and suspend player for an hour\nMUSIC is the music category\nNEXT is phantom category\nROOTS is accompanying music category\nSTATIONID is ids for sprinkling\nTOP40 is currect hits\nNWS is News Weather Sports and will play once then delete"

func CategoriesWriteStub(withinventory bool) {
	userHome, usherr := os.UserHomeDir()
	if usherr != nil {
		log.Println("DB015 Write Categories User Home", usherr)
	}
	log.Println("DB016 User Home", userHome)
	/* 	db, dberr := NewPGSQL()
	   	if dberr != nil {
	   		log.Println("WriteCategories", dberr)
	   	} */
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	log.Println("DB017 Writing Categories to Stub ")
	CategoriesStore = make(map[int]CategoriesStruct)
	var stubname = "/radio/stub"
	if withinventory {
		stubname = "/backup" + stubname
	}
	err4 := os.RemoveAll(userHome + stubname)
	if err4 != nil {
		log.Println("DB018 Remove Stub", err4)
	}

	err3 := os.MkdirAll(userHome+stubname, os.ModePerm)
	if err3 != nil {
		log.Println("DB019 Get Categories row for Stub", err3)
	}
	os.WriteFile(userHome+stubname+"/README.txt", []byte(instructions), os.ModePerm)
	rows, rowserr := conn.Query(ctxsql, "select * from categories order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("DB020 Get Categories row for Stub", err)
		}
		log.Println("DB021 Writing Stub", userHome+stubname+"/"+id)
		err2 := os.Mkdir(userHome+stubname+"/"+id, os.ModePerm)
		if err2 != nil {
			log.Println("DB022 Create Stub", err2)
		}
		if err2 == nil {
			//get all inv items or category read and write
			ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
			conn, _ := SQL.Pool.Acquire(ctxsql)

			ScheduleStore = make(map[int]ScheduleStruct)
			rows, rowserr := conn.Query(ctxsql, "select rowid,category,artist,song,album from inventory  where category = $1", id)

			for rows.Next() {
				var rowid int       // rowid
				var category string // category
				var artist string   // artist
				var song string     // song
				var album string    // Album

				err := rows.Scan(&rowid, &category, &artist, &song, &album)
				if err != nil {
					log.Println("DB029 Get Schedule row", err)
				}
				var invitem = artist + "-" + song + "-" + album
				if err == nil {
					data := GetBucket("mp3", strconv.Itoa(rowid), "COPY")
					os.WriteFile(userHome+stubname+"/"+id+"/"+invitem+".mp3", data, os.ModePerm)
					log.Println("DB022 Write Stub", userHome+stubname+"/"+id+"/"+invitem+".mp3")
				}

			}
			if rowserr != nil {
				log.Println("DB100 Get Schedule row error", rowserr)
			}
			conn.Release()
			ctxsqlcan()

		}
	}
	if rowserr != nil {
		log.Println("DB023 Create Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()

}

var CategoryArray []string

func CategoriesToArray() []string {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	CategoryArray = []string{}
	rows, rowserr := conn.Query(ctxsql, "select * from categories order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("DB024 Get Categories to Array row", err)
		}
		CategoryArray = append(CategoryArray, id)

	}
	if rowserr != nil {
		log.Println("DB025 Get Categories to Array row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
	return CategoryArray

}

func CategoriesDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "delete from categories where rowid =$1", row)

	if rowserr != nil {
		log.Println("DB026 Delete Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func CategoriesUpdate(row int, id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Exec(ctxsql, "update categories set id =$1, description = $2 where rowid = $3", id, desc, row)

	if rowserr != nil {
		log.Println("DB027 Update Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func CategoriesAdd(id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "insert into  categories (id, description) values($1,$2)", id, desc)

	if rowserr != nil {
		log.Println("DB028 Add Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}

type ScheduleStruct struct {
	Row         int    // rowid
	Days        string // days id
	Hours       string // hour part to play
	Position    string // position on schedule
	Category    string // category to play
	Spinstoplay int    // number of items to play
}

var ScheduleStore = make(map[int]ScheduleStruct)
var ScheduleCategory int

func ScheduleGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)

	ScheduleStore = make(map[int]ScheduleStruct)
	rows, rowserr := conn.Query(ctxsql, "select * from schedule order by days,hours,position")
	var rowid int
	var days string
	var hours string
	var position string
	var categories string
	var spinstoplay int
	for rows.Next() {

		err := rows.Scan(&rowid, &days, &hours, &position, &categories, &spinstoplay)
		if err != nil {
			log.Println("DB029 Get Schedule row", err)
		}
		ds := ScheduleStruct{}
		ds.Row = rowid
		ds.Days = days
		ds.Hours = hours
		ds.Position = position
		ds.Category = categories
		ds.Spinstoplay = spinstoplay

		ScheduleStore[len(ScheduleStore)] = ds

	}
	if rowserr != nil {
		log.Println("DB030 Get Schedule row error", rowserr)
	}

	conn.Release()
	ctxsqlcan()

}
func ScheduleDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "delete from schedule where rowid =$1", row)

	if rowserr != nil {
		log.Println("DB031 Delete Schedule row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func ScheduleUpdate(row int, days string, hours string, position string, categories string, spinstoplay int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Exec(ctxsql, "update schedule set days =$1, hours = $2, position = $3, categories = $4, spinstoplay = $5 where rowid = $6", days, hours, position, categories, spinstoplay, row)

	if rowserr != nil {
		log.Println("DB032 Update Schedule row error", rowserr)
	}
	ctxsqlcan()
}
func ScheduleAdd(days string, hours string, position string, categories string, spinstoplay int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Query(ctxsql, "insert into  schedule (days,hours, position,categories,spinstoplay) values($1,$2,$3,$4,$5)", days, hours, position, categories, spinstoplay)

	if rowserr != nil {
		log.Println("DB033 Add Schedule row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func ScheduleCopy(dayfrom, dayto string) {
	log.Println("DB034 ScheduleCopy", dayfrom, dayto)
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	conn2, _ := SQL.Pool.Acquire(ctxsql)
	// delete existing dayto
	_, rowserr := conn.Exec(ctxsql, "delete from schedule where days =$1", dayto)

	if rowserr != nil {
		log.Println("DB035 Delete Schedule row error", rowserr)
	}
	// copy dayf  to dayt

	rows, rowserr2 := conn.Query(ctxsql, "select * from schedule where days = $1 order by days,hours,position ", dayfrom)
	var rowid int
	var days string
	var hours string
	var position string
	var categories string
	var spinstoplay int

	for rows.Next() {

		err := rows.Scan(&rowid, &days, &hours, &position, &categories, &spinstoplay)
		if err != nil {
			log.Println("DB036 Copy Schedule rows next ", err)
		}
		if err == nil {
			_, rowserr1 := conn2.Exec(ctxsql, "insert into  schedule (days,hours, position,categories,spinstoplay) values($1,$2,$3,$4,$5)", dayto, hours, position, categories, spinstoplay)

			if rowserr1 != nil {
				log.Println("DB037 Copy Schedule insert row error1", rowserr1)
			}
		}

	}

	if rowserr2 != nil {
		log.Println("DB038 Copy Schedule row error2", rowserr2)
	}
	conn2.Release()
	conn.Release()

	ctxsqlcan()
}

type InventoryStruct struct {
	Row        int    // rowid
	Category   string // category
	Artist     string // artist
	Song       string // song
	Album      string // Album
	Songlength int    // song length
	Rndorder   string // assigned weekly
	Startson   string //
	Expireson  string
	Lastplayed string
	Dateadded  string
	Spinstoday int    // cleared daily at day reset
	Spinsweek  int    // spins weekly at week reset
	Spinstotal int    // total spins
	Sourcelink string // link to relevant source
}

var InventoryStore = make(map[int]InventoryStruct)
var SelectedInventory int

func InventoryGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)

	InventoryStore = make(map[int]InventoryStruct)
	rows, rowserr := conn.Query(ctxsql, "select * from inventory  order by category,artist,song")
	var row int         // rowid
	var category string // category
	var artist string   // artist
	var song string     // song
	var album string    // Album
	var songlength int  // song length
	var rndorder string // assigned weekly
	var startson string
	var expireson string
	var lastplayed string
	var dateadded string
	var spinstoday int    // cleared daily at day reset
	var spinsweek int     // spins weekly at week reset
	var spinstotal int    // total spins
	var sourcelink string // link to source
	for rows.Next() {
		err := rows.Scan(&row, &category, &artist, &song, &album, &songlength, &rndorder, &startson, &expireson, &lastplayed, &dateadded, &spinstoday, &spinsweek, &spinstotal, &sourcelink)
		if err != nil {
			log.Println("DB039 Get Inventory row", err)
		}
		ds := InventoryStruct{}
		ds.Row = row
		ds.Category = category
		ds.Artist = artist
		ds.Song = song
		ds.Album = album
		ds.Songlength = songlength
		ds.Rndorder = rndorder
		ds.Song = song
		ds.Startson = startson
		ds.Expireson = expireson
		ds.Spinstoday = spinstoday
		ds.Spinsweek = spinsweek
		ds.Spinstotal = spinstotal
		ds.Sourcelink = sourcelink
		InventoryStore[len(InventoryStore)] = ds

	}
	if rowserr != nil {
		log.Println("DB040 Get Inventory row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()

}
func InventoryDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)

	_, rowserr := conn.Exec(ctxsql, "delete from inventory where rowid =$1", row)

	if rowserr != nil {
		log.Println("DB041 Delete Inventory row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func InventoryUpdate(row int, category string, artist string, song string, album string, songlength int, rndorder string, startson string, expireson string, lastplayed string, dateadded string, spinstoday int, spinsweek int, spinstotal int, sourcelink string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	_, rowserr := conn.Exec(ctxsql, "update inventory set category =$1, artist = $2, song = $3, album = $4, songlength = $5, rndorder = $6, startson = $7,expireson = $8, lastplayed = $9, dateadded = $10, spinstoday = $11, spinsweek = $12, spinstotal = $13 , sourcelink = $14 where rowid = $15", category, artist, song, album, songlength, rndorder, startson, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink, row)

	if rowserr != nil {
		log.Println("DB042 Update Inventory row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}

func InventoryAdd(category string, artist string, song string, album string, songlength int, rndorder string, startson string, expireson string, lastplayed string, dateadded string, spinstoday int, spinsweek int, spinstotal int, sourcelink string) int {

	var iactxsql context.Context
	var iactxsqlcan context.CancelFunc
	var iadconn *pgxpool.Conn
	var iadrows pgx.Rows
	//var iarows pgx.Rows
	var iadrowserr error
	//var iarowserr error
	var iarows1err error
	var rowcount = 0
	var rowsc = 0
	var rowid = 0
	var iaconn *pgxpool.Conn
	var iaconn1 *pgxpool.Conn
	iactxsql, iactxsqlcan = context.WithTimeout(context.Background(), 1*time.Minute)

	iadconn, _ = SQL.Pool.Acquire(iactxsql)
	iadrows, iadrowserr = iadconn.Query(iactxsql, "select count(*) from inventory  where (category = $1 and artist = $2 and song = $3 and album = $4)", category, artist, song, album)

	if iadrowserr != nil {
		log.Println("DB043 Add Inventory row error query", iadrowserr)
	}
	rowcount = 0
	rowsc = 0
	for iadrows.Next() {
		iadrowserr = iadrows.Scan(&rowsc)
		if iadrowserr != nil {
			log.Println("DB044 Get Inventory row", iadrowserr)
		}
		rowcount++
	}

	if rowcount > 1 {
		iadconn.Release()
		iactxsqlcan()
		return 0

	}
	iadconn.Release()
	iaconn, _ = SQL.Pool.Acquire(iactxsql)
	_, rowserr := iaconn.Exec(iactxsql, "insert into  inventory (category,artist,song,album,songlength,rndorder,startson,expireson,lastplayed,dateadded,spinstoday,spinsweek,spinstotal,sourcelink) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)", category, artist, song, album, songlength, rndorder, startson, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink)

	if rowserr != nil {
		log.Println("DB045 Add Inventory row error insert", rowserr)
	}
	iaconn1, _ = SQL.Pool.Acquire(iactxsql)
	iarows1, iarowserr1 := iaconn1.Query(iactxsql, "select rowid from inventory  where (category = $1 and artist = $2 and song = $3 and album = $4)", category, artist, song, album)

	if iarowserr1 != nil {
		log.Println("DB046 Add Inventory row error query", iarowserr1)
	}

	for iarows1.Next() {
		iarows1err = iarows1.Scan(&rowid)
		if iarows1err != nil {
			log.Println("DB047 Get Inventory row", iarows1err)
		}
	}
	iaconn.Release()
	iaconn1.Release()
	iactxsqlcan()

	return rowid
}

func ToPDF(reportname, stationid string) {
	switch reportname {
	case "SpinsPerDay":
		PDFInventory("Spins Per Day", "TOP40SPD", stationid)
	case "SpinsPerWeek":
		PDFInventory("Spins Per Week", "TOP40SPW", stationid)
	case "SpinsTotal":
		PDFInventory("TOP40 Total", "TOP40ALL", stationid)
	case "InventoryByCategoryFULL":
		PDFInventory("All Categories", "ALL", stationid)
	case "CategoryList":
		PDFCategory("Category List", stationid)
	case "ScheduleList":
		PDFSchedule("Schedule List", stationid)
	case "DaysList":
		PDFDays("Days List", stationid)
	case "HoursList":
		PDFHours("Hours List", stationid)
	case "TrafficReport":
		PDFTraffic("Traffic Report", stationid)
	}
}
func PDFTraffic(rn, stationid string) {
	//log.Println("PDFInventory", rn, stationid)

	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		Build()

	//darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	merr := m.RegisterHeader(PDFTrafficPageHeader(rn + " for " + stationid))
	if merr != nil {
		log.Fatal(merr.Error())
	}

	m.AddRows(PDFTrafficByAlbum()...)
	docpdf, merr := m.Generate()
	if merr != nil {
		log.Fatal(merr.Error())
	}
	merr = docpdf.Save("TrafficReport.pdf")
	if merr != nil {
		log.Fatal("DB069 PDF Save", merr.Error())
	}
	//Send(stationid+" - "+rn, docpdf.GetBytes(),"REPORTS")

}
func PDFTrafficByAlbum() []core.Row {
	var rowsgti []core.Row

	var contentsRow []core.Row
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	var full pgx.Rows
	var fullerr error
	/*   rowid serial primary key,
	     artist text not null,
	     song   text not null,
	     album  text,
	     playedon text */
	if TrafficAlbum == "" {
		full, fullerr = conn.Query(ctxsql, "select * from traffic where playedon >= '"+TrafficStart+"' and playedon <= '"+TrafficEnd+"' order by album,playedon")
	}
	if TrafficAlbum != "" {
		full, fullerr = conn.Query(ctxsql, "select * from traffic where playedon >= '"+TrafficStart+"' and playedon <= '"+TrafficEnd+"' and album = '"+TrafficAlbum+"' order by album,playedon")
	}
	var rowid int
	var artist string
	var song string
	var album string
	var playedon string //

	var ALBUMCHANGE string

	var itemcount int = 0
	var itemcountgrand int = 0
	rowshead := []core.Row{
		row.New(4).Add(
			col.New(1),
			text.NewCol(1, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(2, "Campaign", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(2, "Artist", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(2, "Song", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(2, "Date", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
		),
	}
	contentsRow = append(contentsRow, rowshead...)
	for full.Next() {

		err := full.Scan(&rowid, &artist, &song, &album, &playedon)
		if err != nil {
			log.Println("DB070 PDF Get Traffic row", err)
		}
		//log.Println("Get Inventory row", category)
		//log.Println("CATCHANGE", CATCHANGE, "db", category)
		if len(ALBUMCHANGE) == 0 {
			ALBUMCHANGE = album
		}
		if ALBUMCHANGE != album {
			ALBUMCHANGE = album
			//rowsgti = append(rowsgti, contentsRow...)
			//var sl = avglen / counttotal
			//rowstot []core.Row{
			rowstotals := append(rowsgti, row.New(20).Add(
				col.New(1),
				text.NewCol(1, "Items: ", props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, strconv.Itoa(itemcount), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, "Grand: ", props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, strconv.Itoa(itemcountgrand), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
			))
			contentsRow = append(contentsRow, rowstotals...)
			rowshead2 := []core.Row{
				row.New(4).Add(
					col.New(1),
					text.NewCol(1, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(2, "Campaign", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(2, "Artist", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(2, "Song", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(2, "Date", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
				),
			}
			contentsRow = append(contentsRow, rowshead2...)
			itemcount = 0

		}
		itemcount++
		itemcountgrand++
		rline := row.New(6).Add(
			col.New(1),
			text.NewCol(1, strconv.Itoa(rowid), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, album, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, artist, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, song, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, playedon, props.Text{Size: 8, Align: align.Left}),
		)
		if itemcount%2 == 0 {
			gray := getGrayColor()
			rline.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		contentsRow = append(contentsRow, rline)

		//contentsRow = append(contentsRow, r)
	}
	rowstotals := append(rowsgti, row.New(20).Add(
		col.New(1),
		text.NewCol(1, "Items: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(itemcount), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, "Grand: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(itemcountgrand), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
	))
	contentsRow = append(contentsRow, rowstotals...)

	if fullerr != nil {
		log.Println("DB069 PDF Get Traffic row error", fullerr)
	}
	conn.Release()
	ctxsqlcan()
	return contentsRow
}

func PDFInventory(rn, cat, stationid string) {
	//log.Println("PDFInventory", rn, stationid)

	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		Build()

	//darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	merr := m.RegisterHeader(getPageHeader(rn + " for " + stationid))
	if merr != nil {
		log.Fatal(merr.Error())
	}

	m.AddRows(PDFInventoryByCategory(cat)...)
	docpdf, merr := m.Generate()
	if merr != nil {
		log.Fatal(merr.Error())
	}
	merr = docpdf.Save(cat + "-InventoryByCategory.pdf")
	if merr != nil {
		log.Fatal("DB048 PDF Save", merr.Error())
	}
	//Send(stationid+" - "+rn, docpdf.GetBytes(),"REPORTS")

}
func PDFCategory(rn, stationid string) {

	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		Build()

	//darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	merr := m.RegisterHeader(getPageHeader(rn + " for " + stationid))
	if merr != nil {
		log.Fatal("DB049 PDF Page Header", merr.Error())
	}

	m.AddRows(PDFCategoryByID()...)
	docpdf, merr := m.Generate()
	if merr != nil {
		log.Fatal("DB050 PDF Generate" + merr.Error())
	}
	merr = docpdf.Save(stationid + "-CategoryList.pdf")
	if merr != nil {
		log.Fatal("DB051 PDF Save ", merr.Error())
	}

}
func PDFSchedule(rn, stationid string) {

	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		Build()

	//darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	merr := m.RegisterHeader(getPageHeader(rn + " for " + stationid))
	if merr != nil {
		log.Fatal("DB052 PDF Page Header ", merr.Error())
	}

	m.AddRows(PDFScheduleByDay()...)
	docpdf, merr := m.Generate()
	if merr != nil {
		log.Fatal("DB053 PDF Generate ", merr.Error())
	}
	merr = docpdf.Save(stationid + "-ScheduleList.pdf")
	if merr != nil {
		log.Fatal("DB053 PDF Save", merr.Error())
	}

}
func PDFDays(rn, stationid string) {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		Build()

	//darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	merr := m.RegisterHeader(getPageHeader(rn + " for " + stationid))
	if merr != nil {
		log.Fatal("DB054 PDF Page Header ", merr.Error())
	}

	m.AddRows(PDFDaysByDay()...)
	docpdf, merr := m.Generate()
	if merr != nil {
		log.Fatal("DB055 PDF GeneraTE ", merr.Error())
	}
	merr = docpdf.Save(stationid + "-DaysList.pdf")
	if merr != nil {
		log.Fatal("DB056 PDF sAVE ", merr.Error())
	}

}
func PDFHours(rn, stationid string) {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(10).
		WithRightMargin(10).
		Build()

	//darkGrayColor := getDarkGrayColor()
	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)
	merr := m.RegisterHeader(getPageHeader(rn + " for " + stationid))
	if merr != nil {
		log.Fatal("DB057 PDF Page Header ", merr.Error())
	}

	m.AddRows(PDFHoursByHour()...)
	docpdf, merr := m.Generate()
	if merr != nil {
		log.Fatal("DB058 PDF Generate ", merr.Error())
	}
	merr = docpdf.Save(stationid + "-HoursList.pdf")
	if merr != nil {
		log.Fatal("DB059 PDF Save ", merr.Error())
	}

}

func getPageHeader(rn string) core.Row {

	return row.New(20).Add(

		col.New(9).Add(
			text.New(rn, props.Text{
				Top:   2,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Center,
				Color: getRedColor(),
			}),
			text.New("For: "+strconv.Itoa(time.Now().Year())+"-"+strconv.Itoa(int(time.Now().Month()))+"-"+strconv.Itoa(time.Now().Day()), props.Text{
				Top:   12,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Center,
				Color: getBlueColor(),
			}),
		),
	)

}

func PDFTrafficPageHeader(rn string) core.Row {
	var ta = TrafficAlbum
	if TrafficAlbum == "" {
		ta = "All"
	}
	return row.New(20).Add(

		col.New(9).Add(
			text.New(rn, props.Text{
				Top:   2,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Center,
				Color: getRedColor(),
			}),

			text.New("From: "+TrafficStart+" To: "+TrafficEnd+" For: "+ta, props.Text{
				Top:   12,
				Style: fontstyle.BoldItalic,
				Size:  8,
				Align: align.Center,
				Color: getBlueColor(),
			}),
		),
	)

}

func PDFScheduleByDay() []core.Row {

	var rowsgti []core.Row

	var contentsRow []core.Row
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)

	full, fullerr := conn.Query(ctxsql, "select * from schedule order by days,hours,position")

	var rowid int       // rowid
	var day string      //
	var hour string     //
	var position string //
	var category string //
	var spinstoplay int //

	var itemcount int = 0
	rowshead := []core.Row{
		row.New(4).Add(
			col.New(1),
			text.NewCol(1, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Day", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Hour", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Pos", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(2, "Cat", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Spins", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
		),
	}
	contentsRow = append(contentsRow, rowshead...)
	for full.Next() {

		err := full.Scan(&rowid, &day, &hour, &position, &category, &spinstoplay)
		if err != nil {
			log.Println("DB060 PDF Get Schedule row", err)
		}

		itemcount++

		rline := row.New(4).Add(
			col.New(1),
			text.NewCol(1, strconv.Itoa(rowid), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, day, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, hour, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, position, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, category, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, strconv.Itoa(spinstoplay), props.Text{Size: 8, Align: align.Left}),
		)
		if itemcount%2 == 0 {
			gray := getGrayColor()
			rline.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		contentsRow = append(contentsRow, rline)

		//contentsRow = append(contentsRow, r)
	}
	rowstotals := append(rowsgti, row.New(20).Add(
		col.New(1),
		text.NewCol(1, "Items: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(itemcount), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
	))
	contentsRow = append(contentsRow, rowstotals...)

	if fullerr != nil {
		log.Println("DB061 PDF Get Schedule row error", fullerr)
	}
	conn.Release()
	ctxsqlcan()
	return contentsRow
}
func PDFDaysByDay() []core.Row {
	var rowsgti []core.Row

	var contentsRow []core.Row
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)

	full, fullerr := conn.Query(ctxsql, "select * from days order by dayofweek")

	var rowid int   // rowid
	var id string   //
	var desc string //
	var dow int     //

	var itemcount int = 0
	rowshead := []core.Row{
		row.New(4).Add(
			col.New(1),
			text.NewCol(1, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "ID", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Desc", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "DOW", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
		),
	}
	contentsRow = append(contentsRow, rowshead...)
	for full.Next() {

		err := full.Scan(&rowid, &id, &desc, &dow)
		if err != nil {
			log.Println("DB062 PDF Get Days row", err)
		}

		itemcount++

		rline := row.New(4).Add(
			col.New(1),
			text.NewCol(1, strconv.Itoa(rowid), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, id, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, desc, props.Text{Size: 8, Align: align.Left}),

			text.NewCol(1, strconv.Itoa(dow), props.Text{Size: 8, Align: align.Left}),
		)
		if itemcount%2 == 0 {
			gray := getGrayColor()
			rline.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		contentsRow = append(contentsRow, rline)

		//contentsRow = append(contentsRow, r)
	}
	rowstotals := append(rowsgti, row.New(20).Add(
		col.New(1),
		text.NewCol(1, "Items: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(itemcount), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
	))
	contentsRow = append(contentsRow, rowstotals...)

	if fullerr != nil {
		log.Println("DB063 PDF Get Days row error", fullerr)
	}
	conn.Release()
	ctxsqlcan()
	return contentsRow
}
func PDFHoursByHour() []core.Row {
	var rowsgti []core.Row

	var contentsRow []core.Row
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)

	full, fullerr := conn.Query(ctxsql, "select * from hours order by id")

	var rowid int   // rowid
	var id string   //
	var desc string //

	var itemcount int = 0
	rowshead := []core.Row{
		row.New(4).Add(
			col.New(1),
			text.NewCol(1, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "ID", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(3, "Desc", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
		),
	}
	contentsRow = append(contentsRow, rowshead...)
	for full.Next() {

		err := full.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("DB064 PDF Get Hours row", err)
		}

		itemcount++

		rline := row.New(4).Add(
			col.New(1),
			text.NewCol(1, strconv.Itoa(rowid), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, id, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(3, desc, props.Text{Size: 8, Align: align.Left}),
		)
		if itemcount%2 == 0 {
			gray := getGrayColor()
			rline.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		contentsRow = append(contentsRow, rline)

		//contentsRow = append(contentsRow, r)
	}
	rowstotals := append(rowsgti, row.New(20).Add(
		col.New(1),
		text.NewCol(1, "Items: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(itemcount), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
	))
	contentsRow = append(contentsRow, rowstotals...)

	if fullerr != nil {
		log.Println("DB065 PDF Get Days row error", fullerr)
	}
	conn.Release()
	ctxsqlcan()
	return contentsRow
}
func PDFCategoryByID() []core.Row {
	var rowsgti []core.Row

	var contentsRow []core.Row
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)

	full, fullerr := conn.Query(ctxsql, "select * from categories order by id")

	var rowid int   // rowid
	var id string   // category
	var desc string // artist

	var itemcount int = 0
	rowshead := []core.Row{
		row.New(4).Add(
			col.New(1),
			text.NewCol(2, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(2, "ID", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(2, "Desc", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
		),
	}
	contentsRow = append(contentsRow, rowshead...)
	for full.Next() {

		err := full.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("DB066 PDF Get Category row", err)
		}

		itemcount++

		rline := row.New(8).Add(
			col.New(1),
			text.NewCol(2, strconv.Itoa(rowid), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, id, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(2, desc, props.Text{Size: 8, Align: align.Left}),
		)
		if itemcount%2 == 0 {
			gray := getGrayColor()
			rline.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		contentsRow = append(contentsRow, rline)

		//contentsRow = append(contentsRow, r)
	}
	rowstotals := append(rowsgti, row.New(20).Add(
		col.New(1),
		text.NewCol(1, "Items: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(itemcount), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
	))
	contentsRow = append(contentsRow, rowstotals...)

	if fullerr != nil {
		log.Println("DB067 PDF Get Category row error", fullerr)
	}
	conn.Release()
	ctxsqlcan()
	return contentsRow
}

func PDFInventoryByCategory(cat string) []core.Row {
	var rowsgti []core.Row

	var contentsRow []core.Row
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	var full pgx.Rows
	var fullerr error
	if strings.Contains(cat, "TOP40") {
		full, fullerr = conn.Query(ctxsql, "select * from inventory where category = 'TOP40' order by category,artist,song")
	}
	if strings.Contains(cat, "ALL") {
		full, fullerr = conn.Query(ctxsql, "select * from inventory  order by category,artist,song")
	}
	var rowid int       // rowid
	var category string // category
	var artist string   // artist
	var song string     // song
	var album string    // Album
	var songlength int  // song length
	var rndorder string // assigned weekly
	var startson string
	var expireson string
	var lastplayed string
	var dateadded string
	var spinstoday int    // cleared daily at day reset
	var spinsweek int     // spins weekly at week reset
	var spinstotal int    // total spins
	var sourcelink string // link to source

	var avglen int = 0
	var counttoday int = 0
	var countweek int = 0
	var counttotal int = 0
	var CATCHANGE string

	var itemcount int = 0
	rowshead := []core.Row{
		row.New(4).Add(
			col.New(1),
			text.NewCol(1, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Category", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Artist", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Song", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Album", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Length", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Last Play", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Today", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Week", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
			text.NewCol(1, "Total", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
		),
	}
	contentsRow = append(contentsRow, rowshead...)
	for full.Next() {

		err := full.Scan(&rowid, &category, &artist, &song, &album, &songlength, &rndorder, &startson, &expireson, &lastplayed, &dateadded, &spinstoday, &spinsweek, &spinstotal, &sourcelink)
		if err != nil {
			log.Println("DB068 PDF Get Inventory row", err)
		}
		//log.Println("Get Inventory row", category)
		//log.Println("CATCHANGE", CATCHANGE, "db", category)
		if len(CATCHANGE) == 0 {
			CATCHANGE = category
		}
		if CATCHANGE != category {
			CATCHANGE = category
			//rowsgti = append(rowsgti, contentsRow...)
			//var sl = avglen / counttotal
			//rowstot []core.Row{
			rowstotals := append(rowsgti, row.New(20).Add(
				col.New(1),
				text.NewCol(1, "Items: ", props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, strconv.Itoa(itemcount), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, "Avg Len: ", props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),

				text.NewCol(1, strconv.Itoa(avglen), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
				}),
				text.NewCol(1, "Today: ", props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, strconv.Itoa(counttoday), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
				}),
				text.NewCol(1, "Week: ", props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, strconv.Itoa(countweek), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
				}),
				text.NewCol(1, "Total: ", props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Right,
				}),
				text.NewCol(1, strconv.Itoa(counttotal), props.Text{
					Top:   5,
					Style: fontstyle.Bold,
					Size:  8,
					Align: align.Left,
				}),
			))
			contentsRow = append(contentsRow, rowstotals...)
			rowshead2 := []core.Row{
				row.New(4).Add(
					col.New(1),
					text.NewCol(1, "Row", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Category", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Artist", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Song", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Album", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Length", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Last Play", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Today", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Week", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
					text.NewCol(1, "Total", props.Text{Size: 9, Align: align.Left, Style: fontstyle.Bold}),
				),
			}
			contentsRow = append(contentsRow, rowshead2...)
			itemcount = 0
			avglen = 0
			counttoday = 0
			countweek = 0
			counttotal = 0
		}
		itemcount++
		avglen = avglen / itemcount
		counttoday = counttoday + spinstoday
		countweek = countweek + spinsweek
		counttotal = counttotal + spinstotal
		rline := row.New(6).Add(
			col.New(1),
			text.NewCol(1, strconv.Itoa(rowid), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, category, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, artist, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, song, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, album, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, strconv.Itoa(songlength), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, lastplayed, props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, strconv.Itoa(spinstoday), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, strconv.Itoa(spinsweek), props.Text{Size: 8, Align: align.Left}),
			text.NewCol(1, strconv.Itoa(spinstotal), props.Text{Size: 8, Align: align.Left}),
		)
		if itemcount%2 == 0 {
			gray := getGrayColor()
			rline.WithStyle(&props.Cell{BackgroundColor: gray})
		}
		contentsRow = append(contentsRow, rline)

		//contentsRow = append(contentsRow, r)
	}
	rowstotals := append(rowsgti, row.New(20).Add(
		col.New(1),
		text.NewCol(1, "Items: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(itemcount), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, "Avg Len: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),

		text.NewCol(1, strconv.Itoa(avglen), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
		text.NewCol(1, "Today: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(counttoday), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
		text.NewCol(1, "Week: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(countweek), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
		text.NewCol(1, "Total: ", props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Right,
		}),
		text.NewCol(1, strconv.Itoa(counttotal), props.Text{
			Top:   5,
			Style: fontstyle.Bold,
			Size:  8,
			Align: align.Left,
		}),
	))
	contentsRow = append(contentsRow, rowstotals...)

	if fullerr != nil {
		log.Println("DB069 PDF Get Inventory row error", fullerr)
	}
	conn.Release()
	ctxsqlcan()
	return contentsRow
}

var AlbumArray []string

func AlbumToArray() []string {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := SQL.Pool.Acquire(ctxsql)
	AlbumArray := []string{}
	rows, rowserr := conn.Query(ctxsql, "select distinct album from traffic order by album")
	var album string

	for rows.Next() {
		err := rows.Scan(&album)
		if err != nil {
			log.Println("DB070 Get Album to Array row", err)
		}
		AlbumArray = append(AlbumArray, album)

	}
	if rowserr != nil {
		log.Println("DB071 Get Album to Array row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
	return AlbumArray

}

func getGrayColor() *props.Color {
	return &props.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}
}

func getBlueColor() *props.Color {
	return &props.Color{
		Red:   10,
		Green: 10,
		Blue:  150,
	}
}

func getRedColor() *props.Color {
	return &props.Color{
		Red:   150,
		Green: 10,
		Blue:  10,
	}
}
