package models

import (
	"fmt"
)

//acce164 struct
type ACCE164 struct {
	AEuuid    string `db:"acce164_uuid" json:"uuid"`
	Aid       string `db:"account_id" json:"id"`
	Adomain   string `db:"account_domain" json:"domain"`
	Gname     string `db:"gateway_name" json:"name"`
	Enumber   string `db:"e164_number" json:"number"`
	Isdefault bool   `db:"acce164_isdefault" json:"isdefault"`
}

//GetAcce164s
func GetAcce164s(condition string) (acce164s []ACCE164, e error) {
	query := fmt.Sprintf("select * from cc_acce164 where %s", condition)
	err := db.Select(&acce164s, query)
	return acce164s, err
}
