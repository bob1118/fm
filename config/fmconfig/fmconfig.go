package fmconfig

import (
	"time"
)

//FMhttpserver struct.
type FMhttpserver struct {
	address      string
	readtimeout  time.Duration
	writetimeout time.Duration
}

//FMdatabase struct {
type FMdatabase struct {
	dbtype        string
	dbname        string
	dbaddress     string
	dbuser        string
	dbpassword    string
	dbtableprefix string
}

//FMruntime struct {
type FMruntime struct {
	runmode    string
	jwtsecret  string
	pagesize   uint
	enablehash bool
}

//FMesl struct {
type FMesl struct {
	esltype     string
	esladdress  string
	eslpassword string
	esltimeout  time.Duration
	eslretries  uint
}

//FMconfig struct {
type FMconfig struct {
	server   FMhttpserver
	database FMdatabase
	esl      FMesl
	runtime  FMruntime
}

//CFGFILE file.
const CFGFILE = "conf.xml"

//NewFmconfig return default fmconfig .
func NewFmconfig() *FMconfig {
	cfg := FMconfig{
		server: FMhttpserver{
			address:      "127.0.0.1:8021",
			readtimeout:  4 * time.Second,
			writetimeout: 4 * time.Second,
		},
		database: FMdatabase{
			dbtype:        "postgres",
			dbname:        "freeswitch",
			dbaddress:     "127.0.0.1:5432",
			dbuser:        "fsdba",
			dbpassword:    "fsdba",
			dbtableprefix: "",
		},
		runtime: FMruntime{
			runmode:    "debug",
			jwtsecret:  "abcdefghijklmnopqrstuvwxyz",
			pagesize:   20,
			enablehash: true,
		},
		esl: FMesl{
			esltype:    "inbound",
			esladdress: "127.0.0.1:8021",
			esltimeout: 4 * time.Second,
			eslretries: 0,
		},
	}
	return &cfg
}

//Read config from file.
func (p *FMconfig) Read(file string) error {
	return nil
}

//Write config to file.
func (p *FMconfig) Write(file string) error {
	return nil
}
