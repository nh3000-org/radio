package config

import (
	"context"
	"log"
	"os"

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
	//conn   pgx.Conn
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
	pooltime, pterr := time.ParseDuration("4096h")
	if pterr != nil {
		Send("messages.postgresql", "Connection Pool  failed to parse time "+pterr.Error(), "postgres")
		log.Fatal("Failed to parse time: ", pterr)
	}
	var TheDB = "postgresql://" + DBuser + ":" + DBpassword + "@" + DBaddress
	//var thedb = DBaddress + DBname + "?user=" + DBuser + "&password=" + DBpassword
	mydb, mydberr := pgxpool.ParseConfig(TheDB)
	mydb.MaxConnIdleTime = pooltime
	mydb.MaxConns = 50
	mydb.MaxConnLifetime = pooltime
	if mydberr != nil {
		log.Println("Unable to connect to parse config database: ", mydberr)

	}
	/* 	conn, myerror = pgx.Connect(ctxsql, mydb)
	   	if myerror != nil {
	   		log.Println("Unable to connect to database: ", myerror)

	   	} */
	mypool, mypoolerr := pgxpool.NewWithConfig(ctxsql, mydb)
	if mypoolerr != nil {
		log.Println("Unable to create connection pool: ", myerror)

	}
	d.Pool = *mypool

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

func (s *SQLconn) DaysGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	DaysStore = make(map[int]DaysStruct)
	conn, _ := s.Pool.Acquire(s.Ctx)
	rows, rowserr := conn.Query(ctxsql, "select * from days order by dayofweek")
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
	conn.Release()
	ctxsqlcan()

}
func (s *SQLconn) DaysDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "delete from days where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Days row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) DaysUpdate(row int, day string, desc string, dow int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Exec(ctxsql, "update days set id =$1, description = $2, dayofweek = $3 where rowid = $4", day, desc, dow, row)

	if rowserr != nil {
		log.Println("Update Days row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) DaysAdd(day string, desc string, dow int) {

	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "insert into  days (id, description, dayofweek) values($1,$2,$3)", day, desc, dow)

	if rowserr != nil {
		log.Println("Add Days row error", rowserr)
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

func (s *SQLconn) HoursGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	HoursStore = make(map[int]HoursStruct)
	rows, rowserr := conn.Query(ctxsql, "select * from hours order by id")
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
	conn.Release()
	ctxsqlcan()

}
func (s *SQLconn) HoursDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "delete from hours where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Hours row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) HoursUpdate(row int, id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Exec(ctxsql, "update hours set id =$1, description = $2 where rowid = $3", id, desc, row)

	if rowserr != nil {
		log.Println("Update Hours row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) HoursAdd(id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "insert into  hours (id, description) values($1,$2)", id, desc)

	if rowserr != nil {
		log.Println("Add Hours row error", rowserr)
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

func (s *SQLconn) CategoriesGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	CategoriesStore = make(map[int]CategoriesStruct)
	rows, rowserr := conn.Query(ctxsql, "select * from categories order by id")
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
	conn.Release()
	ctxsqlcan()

}

var instructions = "Radio Stub Instructions\nBrowse to this file to initiate import\nSongs are identified by ARTIST-SONG-ALBUM.mp3 and ARTIST-SONG-ALBUM-INTRO.mp3 and ARTIST-SONG-ALBUM-OUTRO.mp3 where INTRO and OUTRO are for TOP40 anouncements in the following categories\nADDS, ADDSDRIVETIME and ADDSTOH are used to add advertising to system.\nFILLTOTOH is a phantom category used internally\nIMAGINGID is used to hold artist station plugs\nLIVE is phantom category to indicate live segments and suspend player for an hour\nMUSIC is the music category\nNEXT is phantom category\nROOTS is accompanying music category\nSTATIONID is ids for sprinkling\nTOP40 is currect hits\nNWS is News Weather Sports and will play once then delete"

func (s *SQLconn) CategoriesWriteStub() {
	userHome, usherr := os.UserHomeDir()
	if usherr != nil {
		log.Println("Write Categories User Home", usherr)
	}
	log.Println("User Home", userHome)
	/* 	db, dberr := NewPGSQL()
	   	if dberr != nil {
	   		log.Println("WriteCategories", dberr)
	   	} */
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	log.Println("Writing Categories to Stub ")
	CategoriesStore = make(map[int]CategoriesStruct)
	err4 := os.RemoveAll(userHome + "/radio/stub")
	if err4 != nil {
		log.Println("Remove Stub", err4)
	}

	err3 := os.MkdirAll(userHome+"/radio/stub/", os.ModePerm)
	if err3 != nil {
		log.Println("Get Categories row for Stub", err3)
	}
	os.WriteFile(userHome+"/radio/stub/README.txt", []byte(instructions), os.ModePerm)
	rows, rowserr := conn.Query(ctxsql, "select * from categories order by id")
	var rowid int
	var id string
	var desc string
	for rows.Next() {
		err := rows.Scan(&rowid, &id, &desc)
		if err != nil {
			log.Println("Get Categories row for Stub", err)
		}
		log.Println("Writing Stub", userHome+"/radio/stub/"+id)
		err2 := os.Mkdir(userHome+"/radio/stub/"+id, os.ModePerm)
		if err2 != nil {
			log.Println("Get Categories row for Stub", err2)
		}
	}
	if rowserr != nil {
		log.Println("Get Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()

}

var CategoryArray []string

func (s *SQLconn) CategoriesToArray() []string {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	CategoryArray = []string{}
	rows, rowserr := conn.Query(ctxsql, "select * from categories order by id")
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
	conn.Release()
	ctxsqlcan()
	return CategoryArray

}

func (s *SQLconn) CategoriesDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "delete from categories where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) CategoriesUpdate(row int, id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Exec(ctxsql, "update categories set id =$1, description = $2 where rowid = $3", id, desc, row)

	if rowserr != nil {
		log.Println("Update Categories row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) CategoriesAdd(id string, desc string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "insert into  categories (id, description) values($1,$2)", id, desc)

	if rowserr != nil {
		log.Println("Add Categories row error", rowserr)
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

func (s *SQLconn) ScheduleGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)

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
	conn.Release()
	ctxsqlcan()

}
func (s *SQLconn) ScheduleDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "delete from schedule where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Schedule row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) ScheduleUpdate(row int, days string, hours string, position string, categories string, spinstoplay int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Exec(ctxsql, "update schedule set days =$1, hours = $2, position = $3, categories = $4, spinstoplay = $5 where rowid = $6", days, hours, position, categories, spinstoplay, row)

	if rowserr != nil {
		log.Println("Update Schedule row error", rowserr)
	}
	ctxsqlcan()
}
func (s *SQLconn) ScheduleAdd(days string, hours string, position string, categories string, spinstoplay int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Query(ctxsql, "insert into  schedule (days,hours, position,categories,spinstoplay) values($1,$2,$3,$4,$5)", days, hours, position, categories, spinstoplay)

	if rowserr != nil {
		log.Println("Add Schedule row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) ScheduleCopy(dayfrom, dayto string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	// delete existing dayto
	_, rowserr := conn.Exec(ctxsql, "delete from schedule where days =$1", dayto)

	if rowserr != nil {
		log.Println("Delete Schedule row error", rowserr)
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
			log.Println("Copy Schedule rows next ", err)
		}
		if err == nil {
			conn2, _ := s.Pool.Acquire(s.Ctx)
			_, rowserr1 := conn2.Exec(ctxsql, "insert into  schedule (days,hours, position,categories,spinstoplay) values($1,$2,$3,$4,$5)", dayto, hours, position, categories, spinstoplay)

			if rowserr1 != nil {
				log.Println("Copy Schedule insert row error1", rowserr1)

			}
		}

	}
	if rowserr2 != nil {
		log.Println("Copy Schedule row error2", rowserr2)
	}
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

func (s *SQLconn) InventoryGet() {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)

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
		ds.Startson = startson
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
	conn.Release()
	ctxsqlcan()

}

func (s *SQLconn) InventoryDelete(row int) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)

	_, rowserr := conn.Exec(ctxsql, "delete from inventory where rowid =$1", row)

	if rowserr != nil {
		log.Println("Delete Inventory row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}
func (s *SQLconn) InventoryUpdate(row int, category string, artist string, song string, album string, songlength int, rndorder string, startson string, expireson string, lastplayed string, dateadded string, spinstoday int, spinsweek int, spinstotal int, sourcelink string) {
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 1*time.Minute)
	conn, _ := s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Exec(ctxsql, "update inventory set category =$1, artist = $2, song = $3, album = $4, songlength = $5, rndorder = $6, startson = $7,expireson = $8, lastplayed = $9, dateadded = $10, spinstoday = $11, spinsweek = $12, spinstotal = $13 , sourcelink = $14 where rowid = $15", category, artist, song, album, songlength, rndorder, startson, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink, row)

	if rowserr != nil {
		log.Println("Update Inventory row error", rowserr)
	}
	conn.Release()
	ctxsqlcan()
}

var iactxsql context.Context
var iactxsqlcan context.CancelFunc
var iadconn *pgxpool.Conn
var iadrows pgx.Rows
var iarows pgx.Rows
var iadrowserr error
var iarowserr error
var iarows1err error
var rowcount = 0
var rowsc = 0
var row = 0
var iaconn *pgxpool.Conn
var iaconn1 *pgxpool.Conn

func (s *SQLconn) InventoryAdd(category string, artist string, song string, album string, songlength int, rndorder string, startson string, expireson string, lastplayed string, dateadded string, spinstoday int, spinsweek int, spinstotal int, sourcelink string) int {
	iactxsql, iactxsqlcan = context.WithTimeout(context.Background(), 1*time.Minute)

	iadconn, _ = s.Pool.Acquire(s.Ctx)
	iadrows, iadrowserr = iadconn.Query(iactxsql, "select count(*) from inventory  where (category = $1 and artist = $2 and song = $3 and album = $4)", category, artist, song, album)

	if iadrowserr != nil {
		log.Println("Add Inventory row error query", iadrowserr)
	}
	rowcount = 0
	rowsc = 0
	for iarows.Next() {
		iarowserr = iarows.Scan(&rowsc)
		if iarowserr != nil {
			log.Println("Get Inventory row", iarowserr)
		}
		rowcount++
	}

	if rowcount > 1 {
		iadconn.Release()
		iactxsqlcan()
		return 0

	}
	iadconn.Release()
	iaconn, _ = s.Pool.Acquire(s.Ctx)
	_, rowserr := conn.Exec(iactxsql, "insert into  inventory (category,artist,song,album,songlength,rndorder,startson,expireson,lastplayed,dateadded,spinstoday,spinsweek,spinstotal,sourcelink) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)", category, artist, song, album, songlength, rndorder, startson, expireson, lastplayed, dateadded, spinstoday, spinsweek, spinstotal, sourcelink)

	if rowserr != nil {
		log.Println("Add Inventory row error insert", rowserr)
	}
	iaconn1, _ = s.Pool.Acquire(s.Ctx)
	iarows1, iarowserr1 := iaconn1.Query(iactxsql, "select rowid from inventory  where (category = $1 and artist = $2 and song = $3 and album = $4)", category, artist, song, album)

	if iarowserr1 != nil {
		log.Println("Add Inventory row error query", iarowserr1)
	}

	for iarows1.Next() {
		iarows1err = iarows1.Scan(&row)
		if iarows1err != nil {
			log.Println("Get Inventory row", iarows1err)
		}
	}
	iaconn.Release()
	iaconn1.Release()
	iactxsqlcan()

	return row
}
