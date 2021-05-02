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
