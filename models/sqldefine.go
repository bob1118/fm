package models

const USER_CREATE = "create user fsdba with password 'fsdba'"
const DB_CREATE = "create database freeswitch with owner fsdba"
const DBUSER_AUTH = "grant all privileges on database freeswitch to fsdba"

const CDR_ALEG = `create table cdr_table_a_leg(
	id						serial primary key,
	CallId					uuid not null,
	orig_id					uuid not null,
	term_id					varchar,
	ClientId				uuid,
	IP						inet,
	IPInternal				inet,
	CODEC					varchar,
	directGateway			inet,
	redirectGateway			inet,
	CallerID				varchar,
	TelNumber				varchar,
	TelNumberFull			varchar,
	sip_endpoint_disposition	varchar,
	sip_current_application		varchar
)`
const CDR_BLEG = `create table cdr_table_b_leg(
	id						serial primary key,
	CallId					uuid not null,
	orig_id					uuid not null,
	term_id					varchar,
	ClientId				uuid,
	IP						inet,
	IPInternal				inet,
	CODEC					varchar,
	directGateway			inet,
	redirectGateway			inet,
	CallerID				varchar,
	TelNumber				varchar,
	TelNumberFull			varchar,
	sip_endpoint_disposition	varchar,
	sip_current_application		varchar
)`
const CDR_BOTH = `create table cdr_table_both(
	id 		serial primary key,
	CallId	uuid,
	orig_id uuid,
	Test_id varchar
)`
