package models

const USER_CREATE = "create user fsdba with password 'fsdba'"
const DB_CREATE = "create database freeswitch with owner fsdba"
const DBUSER_AUTH = "grant all privileges on database freeswitch to fsdba"
const DROP_DATABASE_USER = `drop database if exists freeswitch; drop user if exists fsdba;`

//CREATE TABLE if not exists table_name(...)
const CDR_ALEG = `
CREATE TABLE cdr_table_a_leg (
	id serial NOT NULL,
	callid varchar NOT NULL,
	orig_id varchar NOT NULL,
	term_id varchar NULL,
	clientid varchar NULL,
	ip varchar NULL,
	ipinternal varchar NULL,
	codec varchar NULL,
	directgateway varchar NULL,
	redirectgateway varchar NULL,
	callerid varchar NULL,
	telnumber varchar NULL,
	telnumberfull varchar NULL,
	sip_endpoint_disposition varchar NULL,
	sip_current_application varchar NULL,
	CONSTRAINT cdr_table_a_leg_pkey PRIMARY KEY (id)
);
`
const CDR_BLEG = `
CREATE TABLE cdr_table_b_leg (
	id serial NOT NULL,
	callid varchar NOT NULL,
	orig_id varchar NOT NULL,
	term_id varchar NULL,
	clientid varchar NULL,
	ip varchar NULL,
	ipinternal varchar NULL,
	codec varchar NULL,
	directgateway varchar NULL,
	redirectgateway varchar NULL,
	callerid varchar NULL,
	telnumber varchar NULL,
	telnumberfull varchar NULL,
	sip_endpoint_disposition varchar NULL,
	sip_current_application varchar NULL,
	CONSTRAINT cdr_table_b_leg_pkey PRIMARY KEY (id)
);
`
const CDR_BOTH = `
CREATE TABLE cdr_table_both (
	id serial NOT NULL,
	callid varchar NULL,
	orig_id varchar NULL,
	test_id varchar NULL,
	CONSTRAINT cdr_table_both_pkey PRIMARY KEY (id)
);
`
const CC_ACCOUNTS = `
CREATE TABLE cc_accounts (
	account_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	account_id varchar NOT NULL,
	account_name varchar NULL,
	account_auth varchar NULL,
	account_password varchar NULL,
	account_a1hash varchar NULL,
	account_group varchar NULL,
	account_domain varchar NULL,
	account_proxy varchar NULL,
	account_cacheable varchar NULL,
	CONSTRAINT cc_accounts_pkey PRIMARY KEY (account_uuid),
	CONSTRAINT cc_accounts_un UNIQUE (account_id, account_domain)
);
`
const CC_GATEWAYS = `
CREATE TABLE cc_gateways (
	gateway_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	gateway_name varchar NOT NULL,
	gateway_username varchar NULL,
	gateway_realm varchar NULL,
	gateway_fromuser varchar NULL,
	gateway_fromdomain varchar NULL,
	gateway_password varchar NULL,
	gateway_extension varchar NULL,
	gateway_proxy varchar NULL,
	gateway_registerproxy varchar NULL,
	gateway_expire varchar NULL,
	gateway_register varchar NULL,
	gateway_calleridinfrom varchar NULL,
	gateway_extensionincontact varchar NULL,
	gateway_optionping varchar NULL,
	CONSTRAINT cc_gateways_pkey PRIMARY KEY (gateway_uuid),
	CONSTRAINT cc_gateways_un UNIQUE (gateway_name)
);
`
const CC_E164S = `
CREATE TABLE cc_e164s (
	e164_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	gateway_name varchar NOT NULL DEFAULT '',
	e164_number varchar NOT NULL,
	e164_enable bool NULL DEFAULT true,
	e164_lockin bool NULL DEFAULT false,
	e164_lockout bool NULL DEFAULT false,
	CONSTRAINT cc_e164s_pkey PRIMARY KEY (e164_uuid),
	CONSTRAINT cc_e164s_un UNIQUE (e164_number)
);
`
const CC_ACCE164 = `
CREATE TABLE cc_acce164 (
	acce164_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	account_id varchar NOT NULL,
	account_domain varchar NOT NULL,
	gateway_name varchar NOT NULL,
	e164_number varchar NOT NULL,
	acce164_isdefault bool NOT NULL DEFAULT false
);
ALTER TABLE cc_acce164 ADD CONSTRAINT cc_acce164_fk FOREIGN KEY (account_id,account_domain) REFERENCES cc_accounts(account_id,account_domain) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE cc_acce164 ADD CONSTRAINT cc_acce164_fk_1 FOREIGN KEY (gateway_name) REFERENCES cc_gateways(gateway_name) ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE cc_acce164 ADD CONSTRAINT cc_acce164_fk_2 FOREIGN KEY (e164_number) REFERENCES cc_e164s(e164_number) ON DELETE CASCADE ON UPDATE CASCADE;
`
const CC_FIFOS = `
CREATE TABLE cc_fifos (
	fifo_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	fifo_name varchar NOT NULL,
	fifo_importance varchar NULL DEFAULT 0,
	fifo_announce varchar NULL DEFAULT '',
	fifo_holdmusic varchar NULL DEFAULT '',
	CONSTRAINT cc_fifos_un UNIQUE (fifo_name)
);
`
const CC_FIFOMEMBER = `
CREATE TABLE cc_fifomember (
	fifomember_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	fifo_name varchar NOT NULL,
	member_string varchar NOT NULL,
	member_simo varchar NULL DEFAULT 1,
	member_timeout varchar NULL DEFAULT 10,
	member_lag varchar NULL DEFAULT 10
);
ALTER TABLE public.cc_fifomember ADD CONSTRAINT cc_fifomember_fk FOREIGN KEY (fifo_name) REFERENCES cc_fifos(fifo_name) ON DELETE CASCADE ON UPDATE CASCADE;
`
const DEFAULT_ACCOUNTS = `
insert into cc_accounts(account_id,account_name,account_auth,account_password,account_a1hash,account_group,account_domain,account_proxy,account_cacheable) values
('8000','8000','8000','8000','','default','10.10.10.250','10.10.10.250',''),
('8001','8001','8001','8001','','default','10.10.10.250','10.10.10.250',''),
('8002','8002','8002','8002','','default','10.10.10.250','10.10.10.250',''),
('8003','8003','8003','8003','','default','10.10.10.250','10.10.10.250',''),
('8004','8004','8004','8004','','default','10.10.10.250','10.10.10.250',''),
('8005','8005','8005','8005','','default','10.10.10.250','10.10.10.250',''),
('8006','8006','8006','8006','','default','10.10.10.250','10.10.10.250',''),
('8007','8007','8007','8007','','default','10.10.10.250','10.10.10.250',''),
('8008','8008','8008','8008','','default','10.10.10.250','10.10.10.250',''),
('8009','8009','8009','8009','','default','10.10.10.250','10.10.10.250',''),
('8000','8000','8000','8000','','default','1.domain','10.10.10.250',''),
('8001','8001','8001','8001','','default','1.domain','10.10.10.250',''),
('8002','8002','8002','8002','','default','1.domain','10.10.10.250',''),
('8003','8003','8003','8003','','default','1.domain','10.10.10.250',''),
('8004','8004','8004','8004','','default','1.domain','10.10.10.250',''),
('8000','8000','8000','8000','','default','2.domain','10.10.10.250',''),
('8001','8001','8001','8001','','default','2.domain','10.10.10.250',''),
('8002','8002','8002','8002','','default','2.domain','10.10.10.250',''),
('8003','8003','8003','8003','','default','2.domain','10.10.10.250',''),
('8004','8004','8004','8004','','default','2.domain','10.10.10.250','')
`
const DEFAULT_GATEWAYS = `
insert into cc_gateways(gateway_name,gateway_username,gateway_realm,gateway_fromuser,gateway_fromdomain,gateway_password,gateway_extension,gateway_proxy,gateway_registerproxy,gateway_expire,gateway_register,gateway_calleridinfrom,gateway_extensionincontact,gateway_optionping) values
('p2pgatewayname','','p2p.ip','','','','','','','','false','true','',''),
('myfsgateway','1000','10.10.10.200','1000','10.10.10.200','1234','1000','10.10.10.200','10.10.10.200','3600','true','true','true',''),
('vos_in','username','vos.ip','','','password','','','','','false','true','true',''),
('vos_out','username','vos.ip','','','password','','','','','false','true','true','')
`
const DEFAULT_E164S = `
insert into cc_e164s (e164_number)values
('4001234567'),
('1000'),
('1001')
`
const DEFAULT_ACCE164 = `
insert into cc_acce164(account_id,account_domain,gateway_name,e164_number, acce164_isdefault) values
('8000','1.domain','myfsgateway','1000',true),
('8001','1.domain','myfsgateway','1000',true),
('8002','1.domain','myfsgateway','1000',true),
('8003','1.domain','myfsgateway','1000',true),
('8004','1.domain','myfsgateway','1000',true)
`
const DEFAULT_FIFOS = `
insert into cc_fifos(fifo_name)values
('fifos@fifomember'),
('fifos@fifoconsumer')
`
const DEFAULT_FIFOMEMBER = `
insert into cc_fifomember(fifo_name,member_string)values
('fifos@fifomember','user/8000@1.domain'),
('fifos@fifomember','user/8001@1.domain'),
('fifos@fifomember','user/8002@1.domain'),
('fifos@fifomember','user/8003@1.domain'),
('fifos@fifomember','user/8004@1.domain')
`
