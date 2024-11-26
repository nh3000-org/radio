package config

import (
	"context"
	"log"

	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

/*
* posgresql support only
 */
var conn *pgx.Conn
var DBaddress = "localhost:5432/radio"
var DBuser = "postgres"
var DBpassword = "postgres"

//var dburlpropocol = "postgresql://localhost:5432/radio?user=root&password=password"

type sqlconn struct {
	SQLConnect           pgx.Conn
	Pool                 pgxpool.Pool
	Ctx                  context.Context
	Ctxcan               context.CancelFunc
	daysget              *pgconn.StatementDescription
	hoursget             *pgconn.StatementDescription
	categoriesget        *pgconn.StatementDescription
	scheduleget          *pgconn.StatementDescription
	inventoryreset       *pgconn.StatementDescription
	inventorygetschedule *pgconn.StatementDescription
	inventoryget         *pgconn.StatementDescription
	inventoryadd         *pgconn.StatementDescription
}

var myerror error

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	var thedb = "postgresql://" + DBuser + ":" + DBpassword + "@" + DBaddress + "?user=sslmode=verify-ca&pool_max_conns=10&pool_max_conn_lifetime=1h30m"
	// Your own Database URL
	//const DATABASE_URL string = "postgres://postgres:12345678@localhost:5432/postgres?"

	dbConfig, err := pgxpool.ParseConfig(thedb)
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
func NewPGSQL() (*sqlconn, error) {

	var d = new(sqlconn)
	ctxsql, ctxsqlcan := context.WithTimeout(context.Background(), 2048*time.Hour)

	d.Ctxcan = ctxsqlcan
	d.Ctx = ctxsql

	/*	var thedb = DBaddress + DBname + "?user=" + DBuser + "&password=" + DBpassword

		conn, myerror = pgx.Connect(ctxsql, thedb)
		if myerror != nil {
			fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", myerror)
			os.Exit(1)
		} */
	d.SQLConnect = *conn
	Daysget, _ := conn.Prepare(ctxsql, "daysget", "select * from days order by dayofweek")
	d.daysget = Daysget
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
	return d, nil

}
