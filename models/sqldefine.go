package models

const USER_CREATE = "create user fsdba with password 'fsdba'"
const DB_CREATE = "create database freeswitch with owner fsdba"
const DBUSER_AUTH = "grant all privileges on database freeswitch to fsdba"
const DROP_DATABASE_USER = `drop database if exists freeswitch; drop user if exists fsdba;`

//CREATE TABLE if not exists table_name(...)
const CDR_ALEG = `
CREATE TABLE cdr_table_a_leg (
	id serial NOT NULL,
	callid uuid NOT NULL,
	orig_id uuid NOT NULL,
	term_id varchar NULL,
	clientid uuid NULL,
	ip inet NULL,
	ipinternal inet NULL,
	codec varchar NULL,
	directgateway inet NULL,
	redirectgateway inet NULL,
	callerid varchar NULL,
	telnumber varchar NULL,
	telnumberfull varchar NULL,
	sip_endpoint_disposition varchar NULL,
	sip_current_application varchar NULL,
	CONSTRAINT cdr_table_a_leg_pkey PRIMARY KEY (id)
)`
const CDR_BLEG = `
CREATE TABLE cdr_table_b_leg (
	id serial NOT NULL,
	callid uuid NOT NULL,
	orig_id uuid NOT NULL,
	term_id varchar NULL,
	clientid uuid NULL,
	ip inet NULL,
	ipinternal inet NULL,
	codec varchar NULL,
	directgateway inet NULL,
	redirectgateway inet NULL,
	callerid varchar NULL,
	telnumber varchar NULL,
	telnumberfull varchar NULL,
	sip_endpoint_disposition varchar NULL,
	sip_current_application varchar NULL,
	CONSTRAINT cdr_table_b_leg_pkey PRIMARY KEY (id)
)`
const CDR_BOTH = `
CREATE TABLE cdr_table_both (
	id serial NOT NULL,
	callid uuid NULL,
	orig_id uuid NULL,
	test_id varchar NULL,
	CONSTRAINT cdr_table_both_pkey PRIMARY KEY (id)
)`
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
	CONSTRAINT cc_accounts_pkey PRIMARY KEY (account_uuid)
)`
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
	CONSTRAINT cc_gateways_pkey PRIMARY KEY (gateway_uuid)
)`
const CC_E164S = `
CREATE TABLE cc_e164s (
	e164_uuid uuid NOT NULL DEFAULT gen_random_uuid(),
	gateway_uuid uuid NULL,
	e164_number varchar NOT NULL,
	e164_enable bool NULL DEFAULT true,
	e164_lockin bool NULL DEFAULT false,
	e164_lockout bool NULL DEFAULT false,
	CONSTRAINT cc_e164s_pkey PRIMARY KEY (e164_uuid)
)`

const DEFAULT_ACCOUNTS = `
insert into cc_accounts(account_id,account_name,account_auth,account_password,account_a1hash,account_group,account_domain,account_proxy,account_cacheable) values
('8000','8000','8000','8000','','default','$${domain}','$${domain}',''),
('8001','8001','8001','8001','','default','$${domain}','$${domain}',''),
('8002','8002','8002','8002','','default','$${domain}','$${domain}',''),
('8003','8003','8003','8003','','default','$${domain}','$${domain}',''),
('8004','8004','8004','8004','','default','$${domain}','$${domain}',''),
('8005','8005','8005','8005','','default','$${domain}','$${domain}',''),
('8006','8006','8006','8006','','default','$${domain}','$${domain}',''),
('8007','8007','8007','8007','','default','$${domain}','$${domain}',''),
('8008','8008','8008','8008','','default','$${domain}','$${domain}',''),
('8009','8009','8009','8009','','default','$${domain}','$${domain}',''),
('8000','8000','8000','8000','','default','1.domain','10.10.10.250',''),
('8001','8001','8001','8001','','default','1.domain','10.10.10.250',''),
('8002','8002','8002','8002','','default','1.domain','10.10.10.250',''),
('8003','8003','8003','8003','','default','1.domain','10.10.10.250',''),
('8088','8088','8088','8088','','default','1.domain','10.10.10.250',''),
('8000','8000','8000','8000','','default','2.domain','10.10.10.250',''),
('8001','8001','8001','8001','','default','2.domain','10.10.10.250',''),
('8002','8002','8002','8002','','default','2.domain','10.10.10.250',''),
('8003','8003','8003','8003','','default','2.domain','10.10.10.250',''),
('8088','8088','8088','8088','','default','2.domain','10.10.10.250','')
`
const DEFAULT_GATEWAYS = `
insert into cc_gateways(gateway_name,gateway_username,gateway_realm,gateway_fromuser,gateway_fromdomain,gateway_password,gateway_extension,gateway_proxy,gateway_registerproxy,gateway_expire,gateway_register,gateway_calleridinfrom,gateway_extensionincontact,gateway_optionping) values
('p2pgateway.com','','','','','','','','','','false','true','',''),
('p2pgatewayname','','p2pgateway.com','','','','','','','','false','true','',''),
('myfsgateway','1000','10.10.10.200','1000','10.10.10.200','1234','1000','10.10.10.200','10.10.10.200','3600','true','true','true',''),
('vos_inbound','username','1.1.1.1','','','password','','','','','true','true','true',''),
('vos_outbound','username','1.1.1.1','','','password','','','','','true','true','true','')
`
const DEFAULT_E164S = `
insert into cc_e164s (e164_number)values
('8001234567'),
('4001234567'),
('01012345678'),
('02012345678'),
('03012345678')
`
