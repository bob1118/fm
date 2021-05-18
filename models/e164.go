package models

import (
	"database/sql"
	"fmt"
	"strings"
)

//E164 struct.
type E164 struct {
	Euuid    string `db:"e164_uuid" json:"uuid"`
	Gwuuid   string `db:"gateway_uuid" json:"guuid"`
	Enumber  string `db:"e164_number" json:"number"`
	Eenable  bool   `db:"e164_enable" json:"enalbe"`
	Elockin  bool   `db:"e164_lockin" json:"lockin"`
	Elockout bool   `db:"e164_lockout" json:"lockout"`
}

//GetE164sCount function.
func GetE164sCount(condition interface{}) (count int) {
	query := fmt.Sprintf("select count(1) from cc_e164s where %s", condition)
	db.Get(&count, query)
	return count
}

//GetE164s function.
func GetE164s(condition interface{}) (e164s []E164) {
	query := fmt.Sprintf("select * from cc_e164s where %s", condition)
	db.Select(&e164s, query)
	return e164s
}

//IsExistE164Bynumber function.
func IsExistE164Bynumber(new E164) (b bool, old E164) {
	var is bool
	e164 := E164{}
	query := fmt.Sprintf("select * from cc_e164s where e164_number='%s' limit 1", new.Enumber)
	if err := db.Get(&e164, query); err != nil {
		if err == sql.ErrNoRows {
			is = false
		}
	} else {
		is = true
	}
	return is, e164
}

//CreateE164 function.
func CreateE164(in *E164) (e error) {
	var err error

	e164 := in
	query := fmt.Sprintf("insert into cc_e164s(e164_number)values('%s')", e164.Enumber)
	db.MustExec(query)

	return err
}

//IsExistE164Byuuid function.
func IsExistE164Byuuid(uuid string) (exist bool, out E164) {
	var ret bool
	e164 := E164{}
	query := fmt.Sprintf("select * from cc_e164s where true and e164_uuid='%s' limit 1", uuid)
	if err := db.Get(&e164, query); err != nil {
		if err == sql.ErrNoRows {
			ret = false
		}
	} else {
		ret = true
	}
	return ret, e164
}

//ModifyE164 function.
func ModifyE164(old E164, new *E164) (e error) {
	var err error
	query := "update cc_e164s set "
	if new.Gwuuid == "" {
		new.Gwuuid = old.Gwuuid
	} else {
		query = fmt.Sprintf("%s gateway_uuid='%s',", query, new.Gwuuid)
	}
	if new.Enumber == "" {
		new.Enumber = old.Enumber
	} else {
		query = fmt.Sprintf("%s e164_number='%s',", query, new.Enumber)
	}
	if new.Eenable != old.Eenable {
		if new.Eenable {
			query = fmt.Sprintf("%s e164_enable=TRUE,", query)
		} else {
			query = fmt.Sprintf("%s e164_enable=FALSE,", query)
		}
	}
	if new.Elockin != old.Elockin {
		if new.Elockin {
			query = fmt.Sprintf("%s e164_enable=TRUE,", query)
		} else {
			query = fmt.Sprintf("%s e164_enable=FALSE,", query)
		}
	}
	if new.Elockout != old.Elockout {
		if new.Elockout {
			query = fmt.Sprintf("%s e164_enable=TRUE,", query)
		} else {
			query = fmt.Sprintf("%s e164_enable=FALSE,", query)
		}
	}
	//remove ',' from tail.
	query = strings.TrimSuffix(query, ",")
	query = fmt.Sprintf("%s where e164_uuid='%s'", query, new.Euuid)
	db.MustExec(query)
	return err
}

//DeleteE164 funcrtion.
func DeleteE164(uuid string) {
	query := fmt.Sprintf("delete from cc_e164s where e164_uuid='%s'", uuid)
	db.MustExec(query)
}
