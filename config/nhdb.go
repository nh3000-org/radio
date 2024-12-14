package config

import (
	"context"
	"log"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
* posgresql support only
* hostssl  all  all  0.0.0.0/0  cert in pg_hba.conf
* psql "sslmode=require host=localhost dbname=test"
 */
var conn *pgx.Conn

// var DBaddress = "localhost:5432/radio"
var DBaddress = "db.newhorizons3000.org:5432/radio?sslmode=verify-ca"
var DBuser = "postgres"
var DBpassword = "postgres"

//var dburlpropocol = "postgresql://localhost:5432/radio?user=root&password=password"

type SQLconn struct {
	conn   pgx.Conn
	Pool   pgxpool.Pool
	Ctx    context.Context
	Ctxcan context.CancelFunc
}

var myerror error

var DBClientcert = "-----BEGIN CERTIFICATE-----\nMIICyTCCAbECFCJlOZ058bh90IyWT+Z+VS+3K42pMA0GCSqGSIb3DQEBCwUAMCEx\nHzAdBgNVBAMMFmRiLm5ld2hvcml6b25zMzAwMC5vcmcwHhcNMjQxMjA2MTI1MjA1\nWhcNMzQxMjA0MTI1MjA1WjAhMR8wHQYDVQQDDBZkYi5uZXdob3Jpem9uczMwMDAu\nb3JnMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsuUB/1oGzIHbZ3Jj\nRdzqZ3mU2phrwuWrwW02DSbrSqhdrWkhwEDZqcNPclAgU45cP9I0ZS9gMYY+pcv6\naqC5XVZ32kHdLSKW0UGWliyDCU3X7fQEighvutCQqktph0EpY7578WrA7uC8Zap6\nT/zQI9hMf1+YkepfH8oB9m9ekXA/wc0Bf9dHNFzlCl6+UieEL/jUTs4TcOfa/YGK\nOyODoyNmmTTLlf9nr8HwZ5MEHflh2v92AXMG0oB0dD+iNmeOvj9Zq/Hw5QzrkFGZ\nSCp7jxiTJkH/9h+63hECk35ttOFqa0v+ccmTPdrFKOeMc1thqXjQ/k4W/AgFpvuD\nJq6kNQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQAaU7ZUC+bscXNnWLjG3YMoz9RW\n/OpgK6i634PXOUyMGmT6985PBTOPaVsA4FNXDVAjhrbtCZhavzuD+NYluTW9IR5c\nZKWFILTLpzkoEGSphlWIm6I7m60mbYugl1I7bZZX3vGM7IZkvu9OtGXjuESWUSwm\nYmYUSJG0bQ2gMa7DkZUXtLbVT3ZGpQrYW3gVA9LqAvCCRnC9YlN/wmWZHjyNPp8s\nDnSyV9HEtRdk3jTQR7ocUOOX5vVXTEG5K84cgk8DlwwQlIZ/WQVFHsV7eqDh4Lfo\nQLuVg4VCt45Rbo4mVkyY/FX/13guFV2tM04wGU+2WBgIFQvZYWtqh2O4N0zB\n-----END CERTIFICATE-----\n"
var DBClientkey = "-----BEGIN PRIVATE KEY-----\nMIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCy5QH/WgbMgdtn\ncmNF3OpneZTamGvC5avBbTYNJutKqF2taSHAQNmpw09yUCBTjlw/0jRlL2Axhj6l\ny/pqoLldVnfaQd0tIpbRQZaWLIMJTdft9ASKCG+60JCqS2mHQSljvnvxasDu4Lxl\nqnpP/NAj2Ex/X5iR6l8fygH2b16RcD/BzQF/10c0XOUKXr5SJ4Qv+NROzhNw59r9\ngYo7I4OjI2aZNMuV/2evwfBnkwQd+WHa/3YBcwbSgHR0P6I2Z46+P1mr8fDlDOuQ\nUZlIKnuPGJMmQf/2H7reEQKTfm204WprS/5xyZM92sUo54xzW2GpeND+Thb8CAWm\n+4MmrqQ1AgMBAAECggEAAmq5z0QdWuZ3OyDOarYhxuzw1U+MBymjW304/COrjUo/\nMmr88n69zL9vg5fFtAifQkUkrD2gCFNBZevepyb4PM9WAQbKRkQ/yDXMDcfYq7Js\n9Cag6BIrRkQalukRaisidBnxopDtaId0ltAJ5SpBJUwpuRUmzR6Jk0xJ9f/KWzR5\n/gofzpJmlLl4Ymfx4/JAhJFyfUOfq1/1BIIXBEUxhCFDX8fJAog1eJD4i7N2DfrY\nIj2PK15WOizweKnigRGXFghh3dnwLNytm18QW5a/CdGOjCuf7uPPuNfnoW7rZok8\nr7geQh+FeXt3Wbqg3Kxc9Da+IRRix7R8VLAvqNWEcQKBgQDapzQCUkopOBSEIz1/\nQeuvfJLlpDqXtYZ7/KwRPcaBzZfpq17X7SpJG/HDaVm9+RRJxF1f9ft5SBxiGet4\nb1h0uCG3NbTUJXTAk1PmlyxuAQ9Xzd+YlLI+U0ih468Q4tpZrd4kUgtc97uzN7t0\nR0RbeeU5ZGcEfsOefVjK4HNe0QKBgQDRc1OgQz+JvxGhGevjdwa2tRhI0EZaaQOg\nvfc4ppBRLdE3F3j2AFPIBigWPyfHsO9fn5nAYxJIHzI6B7o/6M/JJw/C69/cQg7p\np5/VV73m10dcbtzR9/5O1sl+WGHOdgtvM5hP+IiGsvKRgx3N7F5B/UJkZ4R4scOD\nX/BYzr7wJQKBgApTcSJW7oepzVY8L9BNtaqw8GMF8XpuqS47zYh26WQB6JWxcSYz\nXhbbyfwXgpR1Kd8d9ebtP/YHUMfVP4iNgZjphTYYxDRsnGnny0ONihyb0jSsVU3o\nX86PslNq5D6g5/zqOB5w/XZjgKrDDAg+wVyskgW21yKgNe7LLqFOHkSxAoGBAKWa\nQq9/HDikCqNO5HRXwsYpD0da7ZVEXKr2KAbxoz+cM0QU2f3fKl8HhyB31NMNsWXw\nwdccPfMqP0Mkov0u7UMFEHA0oS38SOAzOausESj4Y6LQwOV+5+Kb7npoFQTxzn6g\n07e/MOsXh7THb4RGdAxG2vyZ4GKxYn14GIePB+bFAoGAKzqQzkb5VP5ML5QwGAmM\n2vn8SU1brebCnnRDSaiBvagLZNeA0X0vVqUWSJ29uKO52x2QpOnBPATJ0th+mHOZ\n/APTv+8hrBkARuu5Kplvfl+MLXbZH+iYWTPbD74kVaJk+Km8tePglR919LvZ8Ysl\nL/QCVUYaP4f/REVsnoShzpg=\n-----END PRIVATE KEY-----\n"

func NewPGSQL() (*SQLconn, error) {
	var d = new(SQLconn)
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 2048*time.Hour)

	d.Ctxcan = ctxsqlcan
	d.Ctx = ctxsql
	var TheDB = "postgresql://" + DBuser + ":" + DBpassword + "@" + DBaddress
	//var thedb = DBaddress + DBname + "?user=" + DBuser + "&password=" + DBpassword

	conn, myerror = pgx.Connect(ctxsql, TheDB)
	if myerror != nil {
		log.Println("Unable to connect to database: ", myerror)

	}
	d.conn = *conn

	return d, nil
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
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("GetDays", dberr)
	}
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")
	DaysStore = make(map[int]DaysStruct)
	rows, rowserr := db.conn.Query(db.Ctx, "select * from days order by dayofweek")
	var rowid int
	var day string
	var desc string
	var dow int
	for rows.Next() {

		err := rows.Scan(&rowid, &day, &desc, &dow)
		if err != nil {
			log.Println("GetDays row", err)
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
		log.Println("GetDays row error", rowserr)
	}
	db.Ctxcan()

}
func DaysDelete(row int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Delete Days", dberr)
	}
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")

	_, rowserr := db.conn.Query(db.Ctx, "delete from days where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Days row error", rowserr)
	}
	db.Ctxcan()
}
func DaysUpdate(row int, day string, desc string, dow int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Update Days", dberr)
	}
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")

	_, rowserr := db.conn.Exec(db.Ctx, "update days set id =$1, description = $2, dayofweek = $3 where rowid = $4", day, desc, dow, row)

	if rowserr != nil {
		log.Println("Update Days row error", rowserr)
	}
	db.Ctxcan()
}
func DaysAdd(day string, desc string, dow int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Add Days", dberr)
	}
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")

	_, rowserr := db.conn.Query(db.Ctx, "insert into  days (id, description, dayofweek) values($1,$2,$3)", day, desc, dow)

	if rowserr != nil {
		log.Println("Add Days row error", rowserr)
	}
	db.Ctxcan()
}

type HoursStruct struct {
	Row  int    // rowid
	Id   string // hours id
	Desc string // alias
}

var HoursStore = make(map[int]HoursStruct)
var SelectedHour int

func HoursGet() {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("GetHours", dberr)
	}

	HoursStore = make(map[int]HoursStruct)
	rows, rowserr := db.conn.Query(db.Ctx, "select * from hours order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("GetHours row", err)
		}
		ds := HoursStruct{}
		ds.Row = rowid
		ds.Id = id
		ds.Desc = desc

		HoursStore[len(HoursStore)] = ds

	}
	if rowserr != nil {
		log.Println("Gethours row error", rowserr)
	}
	db.Ctxcan()

}
func HoursDelete(row int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Delete Hours", dberr)
	}

	_, rowserr := db.conn.Query(db.Ctx, "delete from hours where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Hours row error", rowserr)
	}
	db.Ctxcan()
}
func HoursUpdate(row int, id string, desc string) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Update Hours", dberr)
	}
	_, rowserr := db.conn.Exec(db.Ctx, "update hours set id =$1, description = $2 where rowid = $3", id, desc, row)

	if rowserr != nil {
		log.Println("Update Hours row error", rowserr)
	}
	db.Ctxcan()
}
func HoursAdd(id string, desc string) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Add Hours", dberr)
	}
	_, rowserr := db.conn.Query(db.Ctx, "insert into  hours (id, description) values($1,$2)", id, desc)

	if rowserr != nil {
		log.Println("Add Hours row error", rowserr)
	}
	db.Ctxcan()
}

type CategoriesStruct struct {
	Row  int    // rowid
	Id   string // hours id
	Desc string // alias
}

var CategoriesStore = make(map[int]CategoriesStruct)
var SelectedCategory int

func CategoriesGet() {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("GetCategories", dberr)
	}

	CategoriesStore = make(map[int]CategoriesStruct)
	rows, rowserr := db.conn.Query(db.Ctx, "select * from categories order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("Get Categories row", err)
		}
		ds := CategoriesStruct{}
		ds.Row = rowid
		ds.Id = id
		ds.Desc = desc

		CategoriesStore[len(CategoriesStore)] = ds

	}
	if rowserr != nil {
		log.Println("Get Categories row error", rowserr)
	}
	db.Ctxcan()

}

var CategoryArray []string

func CategoriesToArray() []string {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Get Categories to Array", dberr)
	}

	CategoryArray = []string{}
	rows, rowserr := db.conn.Query(db.Ctx, "select * from categories order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("Get Categories to Array row", err)
		}
		CategoryArray = append(CategoryArray, id)

	}
	if rowserr != nil {
		log.Println("Get Categories to Array row error", rowserr)
	}
	db.Ctxcan()
	return CategoryArray

}

func CategoriesDelete(row int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Delete Catagories", dberr)
	}

	_, rowserr := db.conn.Query(db.Ctx, "delete from categories where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Categories row error", rowserr)
	}
	db.Ctxcan()
}
func CategoriesUpdate(row int, id string, desc string) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Update Categories", dberr)
	}
	_, rowserr := db.conn.Exec(db.Ctx, "update categories set id =$1, description = $2 where rowid = $3", id, desc, row)

	if rowserr != nil {
		log.Println("Update Categories row error", rowserr)
	}
	db.Ctxcan()
}
func CategoriesAdd(id string, desc string) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Add Categories", dberr)
	}
	_, rowserr := db.conn.Query(db.Ctx, "insert into  categories (id, description) values($1,$2)", id, desc)

	if rowserr != nil {
		log.Println("Add Categories row error", rowserr)
	}
	db.Ctxcan()
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
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Get Schedule", dberr)
	}

	ScheduleStore = make(map[int]ScheduleStruct)
	rows, rowserr := db.conn.Query(db.Ctx, "select * from schedule order by days,hours,position")
	var rowid int
	var days string
	var hours string
	var position string
	var categories string
	var spinstoplay int
	for rows.Next() {

		err := rows.Scan(&rowid, &days, &hours, &position, &categories, &spinstoplay)
		if err != nil {
			log.Println("Get Schedule row", err)
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
		log.Println("Get Schedule row error", rowserr)
	}

	db.Ctxcan()

}
func ScheduleDelete(row int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Delete Schedule", dberr)
	}

	_, rowserr := db.conn.Query(db.Ctx, "delete from schedule where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Schedule row error", rowserr)
	}
	ScheduleGet()
	db.Ctxcan()
}
func ScheduleUpdate(row int, days string, hours string, position string, categories string, spinstoplay int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Update Schedule", dberr)
	}
	_, rowserr := db.conn.Exec(db.Ctx, "update schedule set days =$1, hours = $2, position = $3, categories = $4, spinstoplay = $5 where rowid = $6", days, hours, position, categories, spinstoplay, row)

	if rowserr != nil {
		log.Println("Update Schedule row error", rowserr)
	}
	ScheduleGet()
	db.Ctxcan()
}
func ScheduleAdd(days string, hours string, position string, categories string, spinstoplay int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Add Schedule", dberr)
	}
	_, rowserr := db.conn.Query(db.Ctx, "insert into  schedule (days,hours, position,categories,spinstoplay) values($1,$2,$3,$4,$5)", days, hours, position, categories, spinstoplay)

	if rowserr != nil {
		log.Println("Add Schedule row error", rowserr)
	}
	ScheduleGet()
	db.Ctxcan()
}
func ScheduleCopy(dayfrom, dayto string) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Copy Schedule", dberr)
	}
	// delete existing dayto
	_, rowserr := db.conn.Exec(db.Ctx, "delete from schedule where days =$1", dayto)

	if rowserr != nil {
		log.Println("Delete Schedule row error", rowserr)
	}
	// copy dayf  to dayt

	rows, rowserr2 := db.conn.Query(db.Ctx, "select * from schedule where days = $1 order by days,hours,position ", dayfrom)
	var rowid int
	var days string
	var hours string
	var position string
	var categories string
	var spinstoplay int
	db2, dberr2 := NewPGSQL()
	if dberr2 != nil {
		log.Println("Insert Schedule", dberr)
	}
	for rows.Next() {

		err := rows.Scan(&rowid, &days, &hours, &position, &categories, &spinstoplay)
		if err != nil {
			log.Println("Copy Schedule rows next ", err)
		}
		if err == nil {

			_, rowserr1 := db2.conn.Exec(db.Ctx, "insert into  schedule (days,hours, position,categories,spinstoplay) values($1,$2,$3,$4,$5)", dayto, hours, position, categories, spinstoplay)

			if rowserr1 != nil {
				log.Println("Copy Schedule insert row error1", rowserr1)
				db2.Ctxcan()
			}
		}

	}
	if rowserr2 != nil {
		log.Println("Copy Schedule row error2", rowserr2)
	}

	/* 	_, rowserr := db.conn.Query(db.Ctx, "insert into  schedule (days,hours, position,categories,spinstoplay) values($1,$2,$3,$4,$5)", days, hours, position, categories, spinstoplay)

	   	if rowserr != nil {
	   		log.Println("Add Categories row error", rowserr)
	   	} */

	db2.Ctxcan()
	db.Ctxcan()
	ScheduleGet()
}

type InventoryStruct struct {
	Row        int    // rowid
	Category   string // category
	Artist     string // artist
	Song       string // song
	Album      string // Album
	Songlength int    // song length
	Rndorder   string // assigned weekly
	Expireson  time.Time
	Lastplayed time.Time
	Dateadded  time.Time
	Spinstoday int    // cleared daily at day reset
	Spinsweek  int    // spins weekly at week reset
	Spinstotal int    // total spins
	Sourcelink string // link to relevant source
}

var InventoryStore = make(map[int]InventoryStruct)
var SelectedInventory int

func InventoryGet() {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Get Inventory", dberr)
	}

	InventoryStore = make(map[int]InventoryStruct)
	rows, rowserr := db.conn.Query(db.Ctx, "select * from inventory  order by category,artist,song")
	var row int         // rowid
	var category string // category
	var artist string   // artist
	var song string     // song
	var album string    // Album
	var songlength int  // song length
	var rndorder string // assigned weekly
	var expireson time.Time
	var lastplayed time.Time
	var dateadded time.Time
	var spinstoday int    // cleared daily at day reset
	var spinsweek int     // spins weekly at week reset
	var spinstotal int    // total spins
	var sourcelink string // link to source
	for rows.Next() {
		err := rows.Scan(&row, &category, &artist, &song, &album, &songlength, &rndorder, &expireson, &lastplayed, &dateadded, &spinstoday, &spinsweek, &spinstotal, &sourcelink)
		if err != nil {
			log.Println("Get Inventory row", err)
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
		ds.Expireson = expireson
		ds.Spinstoday = spinstoday
		ds.Spinsweek = spinsweek
		ds.Spinstotal = spinstotal
		ds.Sourcelink = sourcelink
		InventoryStore[len(InventoryStore)] = ds

	}
	if rowserr != nil {
		log.Println("Get Inventory row error", rowserr)
	}
	db.Ctxcan()

}

func InventoryDelete(row int) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Delete Inventory", dberr)
	}

	_, rowserr := db.conn.Exec(db.Ctx, "delete from inventory where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Inventory row error", rowserr)
	}
	InventoryGet()
	db.Ctxcan()
}
func InventoryUpdate(row int, category string, artist string, song string, album string, songlength int, rndorder string, expireson time.Time, lastplayed time.Time, dateadded time.Time, spinstoday int, spinsweek int, spinstotal int, sourcelink string) {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Update Inventory", dberr)
	}
	_, rowserr := db.conn.Exec(db.Ctx, "update inventory set category =$1, artist = $2, song = $3, album = $4, songlength = $5, rndorder = $6, expireson = $7, lastplayed = $8, dateadded = $9, spinstoday = $10, spinsweek = $11, spinstotal = $12 , sourcelink = $13 where rowid = $14", category, artist, song, album, songlength, rndorder, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink, row)

	if rowserr != nil {
		log.Println("Update Inventory row error", rowserr)
	}
	db.Ctxcan()
}
func InventoryAdd(category string, artist string, song string, album string, songlength int, rndorder string, expireson time.Time, lastplayed time.Time, dateadded time.Time, spinstoday int, spinsweek int, spinstotal int, sourcelink string) int {
	db, dberr := NewPGSQL()
	if dberr != nil {
		log.Println("Add Inventory", dberr)
	}
	_, rowserr := db.conn.Exec(db.Ctx, "insert into  inventory (category,artist,song,album,songlength,rndorder,expireson,lastplayed,dateadded,spinstoday,spinsweek,spinstotal,sourcelink) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)", category, artist, song, album, songlength, rndorder, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink)

	if rowserr != nil {
		log.Println("Add Inventory row error insert", rowserr)
	}

	rows, rowserr1 := db.conn.Query(db.Ctx, "select row from inventory  where category = $1,artist = #2,song = $3, album = $4", category, artist, song, album)

	if rowserr1 != nil {
		log.Println("Add Inventory row error query", rowserr1)
	}
	var row int // rowid
	for rows.Next() {
		err := rows.Scan(&row)
		if err != nil {
			log.Println("Get Inventory row", err)
		}
	}
	InventoryGet()
	db.Ctxcan()
	return row
}
