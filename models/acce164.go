package models

import "fmt"

// -- public.cc_acce164 definition

// -- Drop table

// -- DROP TABLE cc_acce164;

// CREATE TABLE cc_acce164 (
// 	acce164_uuid uuid NOT NULL,
// 	account_id varchar NOT NULL,
// 	account_domain varchar NOT NULL,
// 	gateway_name varchar NOT NULL,
// 	e164_number varchar NOT NULL
// 	acce164_isdefault bool NOT NULL DEFAULT false
// );
// COMMENT ON TABLE public.cc_acce164 IS 'useragent''s account dial out with gateway';

// -- Permissions

// ALTER TABLE public.cc_acce164 OWNER TO fsdba;
// GRANT ALL ON TABLE public.cc_acce164 TO fsdba;

// -- public.cc_acce164 foreign keys

// ALTER TABLE public.cc_acce164 ADD CONSTRAINT cc_acce164_fk FOREIGN KEY (account_id,account_domain) REFERENCES cc_accounts(account_id,account_domain) ON DELETE CASCADE ON UPDATE CASCADE;
// ALTER TABLE public.cc_acce164 ADD CONSTRAINT cc_acce164_fk_1 FOREIGN KEY (gateway_name) REFERENCES cc_gateways(gateway_name) ON DELETE CASCADE ON UPDATE CASCADE;
// ALTER TABLE public.cc_acce164 ADD CONSTRAINT cc_acce164_fk_2 FOREIGN KEY (e164_number) REFERENCES cc_e164s(e164_number) ON DELETE CASCADE ON UPDATE CASCADE;

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
