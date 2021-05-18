package models

import "fmt"

type Fifomember struct {
	FMuuid   string `db:"fifomember_uuid" json:"uuid"`
	Fname    string `db:"fifo_name" json:"name"`
	Mstring  string `db:"member_string" json:"string"`
	Msimo    string `db:"member_simo" json:"simo"`
	Mtimeout string `db:"member_timeout" json:"timeout"`
	Mlag     string `db:"member_lag" json:"lag"`
}

//GetFifos function.
func GetFifomembers(condition interface{}) (fifomembers []Fifomember, e error) {
	query := fmt.Sprintf("select * from cc_fifomember where %s", condition)
	err := db.Select(&fifomembers, query)
	return fifomembers, err
}
