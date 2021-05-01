package models

import (
	"fmt"
	"log"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db, pgdb *sqlx.DB

//init function.
func init() {
	InitFreeswitch("user=postgres password=fuckIBM host=127.0.0.1 dbname=postgres sslmode=disable")
}

//openPgsql function.
func openPgsql(s string) (db *sqlx.DB, e error) {
	var err error
	db, err = sqlx.Connect("postgres", s)
	if err != nil {
		log.Println(err.Error())
	}
	return db, err
}

//closePgsql funciton.
func closePgsql(db *sqlx.DB) { db.Close() }

//initFreeswitch function.
func InitFreeswitch(strcon string) {
	var err error

	//freeswitch database.
	pgsqlInitFreeswitch(strcon)

	//freeswitch tables.
	s := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", fmconfig.CFG.Database.User, fmconfig.CFG.Database.Password, fmconfig.CFG.Database.Host, fmconfig.CFG.Database.Name)
	if db, err = openPgsql(s); err != nil {
		pgsqlInitFreeswitchCDR()
		pgsqlInitFreeswitchAccounts()
		pgsqlInitFreeswitchGateways()
		pgsqlInitFreeswitchE164s()
	}
}

func pgsqlInitFreeswitch(strcon string) {
	var err error
	var isFound bool
	if pgdb, err = openPgsql(strcon); err != nil {
		log.Println(err.Error())
	} else {
		if err = pgdb.Get(&isFound, "select count(1)!=0 as isFound from pg_user where usename =$1", "fsdba"); err != nil {
			log.Println(err.Error())
		} else {
			if !isFound { //create user.
				pgdb.MustExec(USER_CREATE)
			}
			if err = pgdb.Get(&isFound, "select count(1)!=0 as isFound from pg_database where datname=$1", "freeswitch"); err != nil {
				log.Println(err.Error())
			} else {
				if !isFound { //create db.
					pgdb.MustExec(DB_CREATE)
					pgdb.MustExec(DBUSER_AUTH)
				}
			}
		}
		closePgsql(pgdb)
	}
}

//pgsqlInitFreeswitchCDR function
func pgsqlInitFreeswitchCDR() {
	var err error
	var isFound bool

	if err = db.Get(&isFound, "select count(1)!=0 as isFound from pg_tables where tablename =$1", "table cdr_table_a_leg"); err != nil {
		log.Println(err.Error())
	} else {
		if !isFound {
			db.MustExec(CDR_ALEG)
		}
	}
	if err = db.Get(&isFound, "select count(1)!=0 as isFound from pg_tables where tablename =$1", "table cdr_table_b_leg"); err != nil {
		log.Println(err.Error())
	} else {
		if !isFound {
			db.MustExec(CDR_BLEG)
		}
	}
	if err = db.Get(&isFound, "select count(1)!=0 as isFound from pg_tables where tablename =$1", "table cdr_table_both"); err != nil {
		log.Println(err.Error())
	} else {
		if !isFound {
			db.MustExec(CDR_BOTH)
		}
	}
}

//pgsqlInitFreeswitchAccounts
func pgsqlInitFreeswitchAccounts() {}

//pgsqlInitFreeswitchGateways
func pgsqlInitFreeswitchGateways() {}

//pgsqlInitFreeswitchE164s
func pgsqlInitFreeswitchE164s() {}
