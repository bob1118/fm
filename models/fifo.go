package models

import "fmt"

type Fifo struct {
	Fuuid       string `db:"fifo_uuid" json:"uuid"`
	Fname       string `db:"fifo_name" json:"name"`
	Fimportance string `db:"fifo_importance" json:"importance"`
	Fannounce   string `db:"fifo_announce" json:"announce"`
	Fholdmusic  string `db:"fifo_holdmusic" json:"holdmusic"`
}

//GetFifos function.
func GetFifos(condition interface{}) (fifos []Fifo, e error) {
	query := fmt.Sprintf("select * from cc_fifos where %s", condition)
	err := db.Select(&fifos, query)
	return fifos, err
}
