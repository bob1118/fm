package models

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"githug.com/bob118/fm/config/fmconfig"
)

var db *sqlx.DB

//init function.
func init() {
	var err error
	strcon := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable", fmconfig.CFG.Database.User, fmconfig.CFG.Database.Password, fmconfig.CFG.Database.Host, fmconfig.CFG.Database.Name)
	db, err = sqlx.Connect(fmconfig.CFG.Database.Type, strcon)
	if err != nil {
		log.Println(err)
	}
}

//CloseDB funciton.
func CloseDB() {
	defer db.Close()
}
