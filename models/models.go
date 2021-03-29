package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"

	"github.com/jmoiron/sqlx"
	"githug.com/bob118/fm/config/fmconfig"
)

var db *sqlx.DB

//init function.
func init() {
	var err error
	strcon := fmt.Sprintf("user=%s password=%s host=%s dbname=%s", fmconfig.CFG.Database.User, fmconfig.CFG.Database.Password, fmconfig.CFG.Database.Address, fmconfig.CFG.Database.Name)
	db, err = sqlx.Connect(fmconfig.CFG.Database.Type, strcon)
	if err != nil {
		log.Println(err)
	}
}

//CloseDB funciton.
func CloseDB() {
	defer db.Close()
}

//MakeA1Hash function.
func MakeA1Hash(in string) (s string) {
	h := md5.New()
	io.WriteString(h, in)
	return fmt.Sprintf("%x", h.Sum(nil))
}
