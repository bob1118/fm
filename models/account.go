package models

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/bob118/fm/utils"
)

// -- public.cc_accounts definition

// -- Drop table

// -- DROP TABLE public.cc_accounts;

// CREATE TABLE public.cc_accounts (
// 	account_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
// 	account_id varchar NOT NULL,
// 	account_name varchar NULL,
// 	account_auth varchar NULL,
// 	account_password varchar NULL,
// 	account_a1hash varchar NULL,
// 	account_group varchar NULL,
// 	account_domain varchar NULL,
// 	account_proxy varchar NULL,
// 	account_cacheable varchar NULL,
// 	CONSTRAINT cc_accounts_pkey PRIMARY KEY (account_uuid)
// );
// COMMENT ON TABLE public.cc_accounts IS 'freeswitch mod_sofia profile internal account@domain.';

// -- Permissions

// ALTER TABLE public.cc_accounts OWNER TO fsdba;
// GRANT ALL ON TABLE public.cc_accounts TO fsdba;

type Account struct {
	Auuid      string `db:"account_uuid" json:"uuid"`
	Aid        string `db:"account_id" json:"id"`
	Aname      string `db:"account_name" json:"name"`
	Aauth      string `db:"account_auth" json:"auth"`
	Apassword  string `db:"account_password" json:"password"`
	Aa1hash    string `db:"account_a1hash" json:"a1hash"`
	Agroup     string `db:"account_group" json:"group:"`
	Adomain    string `db:"account_domain" json:"domain"`
	Aproxy     string `db:"account_proxy" json:"proxy"`
	Acacheable string `db:"account_cacheable" json:"cacheable"`
}

//GetAccountsCount
func GetAccountsCount(condition interface{}) (count int) {
	q := fmt.Sprintf("select count(1) from cc_accounts where %s", condition)
	db.Get(&count, q)
	return count
}

//GetAccounts
func GetAccounts(condition interface{}) (accounts []Account) {
	q := fmt.Sprintf("select * from cc_accounts where %s", condition)
	db.Select(&accounts, q)
	return accounts
}

//IsExistByiddomain
func IsExistByiddomain(new Account) (isExist bool, old Account) {
	var is bool
	ua := Account{}
	q := fmt.Sprintf("select * from cc_accounts where  account_id='%s' and account_domain='%s' limit 1", new.Aid, new.Adomain)
	if err := db.Get(&ua, q); err != nil {
		if err == sql.ErrNoRows {
			is = false
		}
	} else {
		is = true
	}
	return is, ua
}

//CreateAccount
func CreateAccount(in *Account) (e error) {
	var err error

	ua := in
	q := fmt.Sprintf("insert into cc_accounts(account_id,account_name,account_auth,account_password,account_a1hash,account_group,account_domain,account_proxy,account_cacheable)values('%s','%s','%s','%s','%s','%s','%s','%s','%s')",
		ua.Aid, ua.Aname, ua.Aauth, ua.Apassword, ua.Aa1hash, ua.Agroup, ua.Adomain, ua.Aproxy, ua.Acacheable)
	db.MustExec(q)

	return err
}

//CreateAccountEx
//gen safe uuid
func CreateAccountEx(in Account) (e error) { return nil }

//IsExistByuuid
func IsExistByuuid(uuid string) (b bool, out Account) {
	var isExist bool
	ua := Account{}

	q := fmt.Sprintf("select * from cc_accounts where true and account_uuid='%s' limit 1", uuid)
	if err := db.Get(&ua, q); err != nil {
		if err == sql.ErrNoRows {
			isExist = false
		}
	} else {
		isExist = true
	}
	return isExist, ua
}

//ModifyAccount
func ModifyAccount(old Account, new *Account) (e error) {

	isNewHash := 0
	q := "update cc_accounts set "
	if new.Aid == "" {
		new.Aid = old.Aid
	} else {
		q = fmt.Sprintf("%s account_id='%s',", q, new.Aid)
		isNewHash++
	}
	if new.Aname == "" {
		new.Aname = old.Aname
	} else {
		q = fmt.Sprintf("%s account_name='%s',", q, new.Aname)
	}
	if new.Aauth == "" {
		new.Aauth = old.Aauth
	} else {
		q = fmt.Sprintf("%s account_auth='%s',", q, new.Aauth)
	}
	if new.Apassword == "" {
		new.Apassword = old.Apassword
	} else {
		q = fmt.Sprintf("%s account_password='%s',", q, new.Apassword)
		isNewHash++
	}
	if new.Agroup == "" {
		new.Agroup = old.Agroup
	} else {
		q = fmt.Sprintf("%s account_group='%s',", q, new.Agroup)
	}
	if new.Adomain == "" {
		new.Adomain = old.Adomain
	} else {
		q = fmt.Sprintf("%s account_domain='%s',", q, new.Adomain)
		isNewHash++
	}
	if new.Aproxy == "" {
		new.Aproxy = old.Aproxy
	} else {
		q = fmt.Sprintf("%s account_proxy='%s',", q, new.Aproxy)
	}
	if new.Acacheable == "" {
		new.Acacheable = old.Acacheable
	} else {
		q = fmt.Sprintf("%s account_cacheable='%s',", q, new.Acacheable)
	}

	if isNewHash > 0 { //make a1hash, MakeA1Hash(user:domain:password)
		newHash := utils.MakeA1Hash(fmt.Sprintf("%s:%s:%s", new.Aid, new.Adomain, new.Apassword))
		q = fmt.Sprintf("%s account_a1hash='%s',", q, newHash)
		new.Aa1hash = newHash
	}

	//remove ',' from tail.
	q = strings.TrimSuffix(q, ",")
	q = fmt.Sprintf("%s where account_uuid='%s'", q, new.Auuid)
	db.MustExec(q)
	return nil
}

//DeleteAccount function.
func DeleteAccount(uuid string) {
	q := fmt.Sprintf("delete from cc_accounts where account_uuid='%s'", uuid)
	db.MustExec(q)
}

//DistinctAccountDomains function.
func DistinctAccountDomains() (s []string, e error) {
	var err error
	var domains []string
	q := "select distinct(account_domain) from cc_accounts"
	if dberror := db.Select(&domains, q); dberror != nil {
		err = dberror
	}
	return domains, err
}
