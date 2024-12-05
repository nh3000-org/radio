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
var DBaddress = "localhost:5432/radio"
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

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	var TheDB = "postgresql://" + DBuser + ":" + DBpassword + "@" + DBaddress + "?user=sslmode=verify-ca&pool_max_conns=10&pool_max_conn_lifetime=1h30m"
	// Your own Database URL
	//const DATABASE_URL string = "postgres://postgres:12345678@localhost:5432/postgres?"

	dbConfig, err := pgxpool.ParseConfig(TheDB)
	if err != nil {
		log.Fatal("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
		log.Println("Before acquiring the connection pool to the database!!")
		return true
	}

	dbConfig.AfterRelease = func(c *pgx.Conn) bool {
		log.Println("After releasing the connection pool to the database!!")
		return true
	}
	dbConfig.BeforeClose = func(c *pgx.Conn) {
		log.Println("Closed the connection pool to the database!!")
	}

	return dbConfig
}
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

	/* 	Hoursget, _ := conn.Prepare(ctxsql, "hoursget", "select * from hours order by id")
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
	   	d.inventoryadd = Inventoryadd */
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

func GetDays() {
	db, dberr := NewPGSQL()
 	if dberr != nil {
		log.Println("GetDays", dberr)
	} 
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")

	rows, rowserr := db.conn.Query(db.Ctx, "select * from days order by dayofweek")
	for rows.Next() {
		var rowid int
		var day string
		var desc string
		var dow int
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
func DeleteDays(row int) {
	db, dberr := NewPGSQL()
 	if dberr != nil {
		log.Println("Delete Days", dberr)
	} 
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")

	_, rowserr := db.conn.Query(db.Ctx, "delete from days where rowid =$1",row)

	if rowserr != nil {
		log.Println("Delete Days row error", rowserr)
	}
	db.Ctxcan()
}
func UpdateDays(row int, day string,desc string,dow int) {
	db, dberr := NewPGSQL()
 	if dberr != nil {
		log.Println("Update Days", dberr)
	} 
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")

	_, rowserr := db.conn.Query(db.Ctx, "update days set day =$1, desc = $2, dow = $3 where rowid = $4",day,desc,dow,row)

	if rowserr != nil {
		log.Println("Delete Days row error", rowserr)
	}
	db.Ctxcan()
}
func AddDays(day string,desc string,dow int) {
	db, dberr := NewPGSQL()
 	if dberr != nil {
		log.Println("Add Days", dberr)
	} 
	//db.conn.Prepare(db.Ctx, "daysget", "select * from days order by dayofweek")

	_, rowserr := db.conn.Query(db.Ctx, "insert into  days (day, desc, dow) values($1,$2.$3)",day,desc,dow)

	if rowserr != nil {
		log.Println("Add Days row error", rowserr)
	}
	db.Ctxcan()
}