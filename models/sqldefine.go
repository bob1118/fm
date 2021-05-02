package models

const USER_CREATE = "create user fsdba with password 'fsdba'"
const DB_CREATE = "create database freeswitch with owner fsdba"
const DBUSER_AUTH = "grant all privileges on database freeswitch to fsdba"

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
