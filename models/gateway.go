package models

import (
	"database/sql"
	"fmt"
	"strings"
)

type Gateway struct {
	Guuid               string `db:"gateway_uuid" json:"uuid"`
	Gname               string `db:"gateway_name" json:"name"`
	Gusername           string `db:"gateway_username" json:"username"`
	Grealm              string `db:"gateway_realm" json:"realm"`
	Gfromuser           string `db:"gateway_fromuser" json:"fromuser"`
	Gfromdomain         string `db:"gateway_fromdomain" json:"fromdomain"`
	Gpassword           string `db:"gateway_password" json:"password"`
	Gextension          string `db:"gateway_extension" json:"extension"`
	Gproxy              string `db:"gateway_proxy" json:"proxy"`
	Gregisterproxy      string `db:"gateway_registerproxy" json:"registerproxy"`
	Gexpire             string `db:"gateway_expire" json:"expire"`
	Gregister           string `db:"gateway_register" json:"register"`
	Gcalleridinfrom     string `db:"gateway_calleridinfrom" json:"calleridinfrom"`
	Gextensionincontact string `db:"gateway_extensionincontact" json:"extensionincontact"`
	Goptionping         string `db:"gateway_optionping" json:"optionping"`
}

//GetGatewaysCount function.
func GetGatewaysCount(condition interface{}) (count uint) {
	query := fmt.Sprintf("select count(1) from cc_gateways where %s", condition)
	db.Get(&count, query)
	return count
}

//GetGateways function.
func GetGateways(condition interface{}) (gateways []Gateway) {
	query := fmt.Sprintf("select * from cc_gateways where %s", condition)
	db.Select(&gateways, query)
	return gateways
}

//IsExistGatewayByname function.
func IsExistGatewayByname(new Gateway) (b bool, old Gateway) {
	var is bool
	gw := Gateway{}
	query := fmt.Sprintf("select * from cc_gateways where true and gateway_name='%s' limit 1", new.Gname)
	if err := db.Get(&gw, query); err != nil {
		if err == sql.ErrNoRows {
			is = false
		}
	} else {
		is = true
	}
	return is, gw
}

//CreateGateway function.
func CreateGateway(in *Gateway) (e error) {
	var err error

	gw := in
	q := fmt.Sprintf("insert into cc_gateways(gateway_name,gateway_username,gateway_realm,gateway_fromuser,gateway_fromdomain,gateway_password,gateway_extension,gateway_proxy,gateway_registerproxy,gateway_expire,gateway_register,gateway_calleridinfrom,gateway_extensionincontact,gateway_optionping)values('%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s','%s')",
		gw.Gname, gw.Gusername, gw.Grealm, gw.Gfromuser, gw.Gfromdomain, gw.Gpassword, gw.Gextension, gw.Gproxy, gw.Gregisterproxy, gw.Gexpire, gw.Gregister, gw.Gcalleridinfrom, gw.Gextensionincontact, gw.Goptionping)
	db.MustExec(q)

	return err
}

//IsExistGatewayByuuid function.
func IsExistGatewayByuuid(u string) (b bool, g Gateway) {
	var ret bool
	gateway := Gateway{}
	query := fmt.Sprintf("select * from cc_gateways where true and gateway_uuid='%s' limit 1", u)
	if err := db.Get(&gateway, query); err != nil {
		if err == sql.ErrNoRows {
			ret = false
		}
	} else {
		ret = true
	}
	return ret, gateway
}

//ModifyGateway function.
func ModifyGateway(old Gateway, new *Gateway) (e error) {
	var err error
	query := "update cc_gateways set "

	if new.Gname == "" {
		new.Gname = old.Gname
	} else {
		query = fmt.Sprintf("%s gateway_name='%s',", query, new.Gname)
	}
	if new.Gusername == "" {
		new.Gusername = old.Gusername
	} else {
		query = fmt.Sprintf("%s gateway_username='%s',", query, new.Gusername)
	}
	if new.Grealm == "" {
		new.Grealm = old.Grealm
	} else {
		query = fmt.Sprintf("%s gateway_realm='%s',", query, new.Grealm)
	}
	if new.Gfromuser == "" {
		new.Gfromuser = old.Gfromuser
	} else {
		query = fmt.Sprintf("%s gateway_fromuser='%s',", query, new.Gfromuser)
	}
	if new.Gfromdomain == "" {
		new.Gfromdomain = old.Gfromdomain
	} else {
		query = fmt.Sprintf("%s gateway_fromdomain='%s',", query, new.Gfromdomain)
	}
	if new.Gpassword == "" {
		new.Gpassword = old.Gpassword
	} else {
		query = fmt.Sprintf("%s gateway_password='%s',", query, new.Gpassword)
	}
	if new.Gextension == "" {
		new.Gextension = old.Gextension
	} else {
		query = fmt.Sprintf("%s gateway_extension='%s',", query, new.Gextension)
	}
	if new.Gproxy == "" {
		new.Gproxy = old.Gproxy
	} else {
		query = fmt.Sprintf("%s gateway_proxy='%s',", query, new.Gproxy)
	}
	if new.Gregisterproxy == "" {
		new.Gregisterproxy = old.Gregisterproxy
	} else {
		query = fmt.Sprintf("%s gateway_registerproxy='%s',", query, new.Gregisterproxy)
	}
	if new.Gexpire == "" {
		new.Gexpire = old.Gexpire
	} else {
		query = fmt.Sprintf("%s gateway_expire='%s',", query, new.Gexpire)
	}
	if new.Gregister == "" {
		new.Gregister = old.Gregister
	} else {
		query = fmt.Sprintf("%s gateway_register='%s',", query, new.Gregister)
	}
	if new.Gcalleridinfrom == "" {
		new.Gcalleridinfrom = old.Gcalleridinfrom
	} else {
		query = fmt.Sprintf("%s gateway_calleridinfrom='%s',", query, new.Gcalleridinfrom)
	}
	if new.Gextensionincontact == "" {
		new.Gextensionincontact = old.Gextensionincontact
	} else {
		query = fmt.Sprintf("%s gateway_extensionincontact='%s',", query, new.Gextensionincontact)
	}
	if new.Goptionping == "" {
		new.Goptionping = old.Goptionping
	} else {
		query = fmt.Sprintf("%s gateway_optionping='%s',", query, new.Goptionping)
	}
	//remove ',' from tail.
	query = strings.TrimSuffix(query, ",")
	query = fmt.Sprintf("%s where gateway_uuid='%s'", query, new.Guuid)
	db.MustExec(query)
	err = nil
	return err
}

//DeleteGateway funcrtion.
func DeleteGateway(u string) {
	query := fmt.Sprintf("delete from cc_gateways where gateway_uuid='%s'", u)
	db.MustExec(query)
}
