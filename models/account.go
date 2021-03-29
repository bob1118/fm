package models

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
