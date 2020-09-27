package fmconfig

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"time"
)

//FMhttpserver struct.
type FMhttpserver struct {
	Address      string
	Readtimeout  time.Duration
	Writetimeout time.Duration
}

//FMdatabase struct {
type FMdatabase struct {
	Type        string
	Name        string
	Address     string
	User        string
	Password    string
	Tableprefix string
}

//FMruntime struct {
type FMruntime struct {
	Runmode    string
	Jwtsecret  string
	Pagesize   uint
	Enablehash bool
}

//FMesl struct {
type FMesl struct {
	Type     string
	Address  string
	Password string
	Timeout  time.Duration
	Retries  uint
}

//FMconfig struct {
type FMconfig struct {
	Server   FMhttpserver
	Database FMdatabase
	Runtime  FMruntime
	Esl      FMesl
}

//CFGFILE file.
const CFGFILE = "conf.xml"

//NewFmconfig return default fmconfig .
func NewFmconfig() *FMconfig {
	c := FMconfig{
		Server: FMhttpserver{
			Address:      "127.0.0.1:8021",
			Readtimeout:  4 * time.Second,
			Writetimeout: 4 * time.Second,
		},
		Database: FMdatabase{
			Type:        "postgres",
			Name:        "freeswitch",
			Address:     "127.0.0.1:5432",
			User:        "fsdba",
			Password:    "fsdba",
			Tableprefix: "",
		},
		Runtime: FMruntime{
			Runmode:    "debug",
			Jwtsecret:  "abcdefghijklmnopqrstuvwxyz",
			Pagesize:   20,
			Enablehash: true,
		},
		Esl: FMesl{
			Type:    "inbound",
			Address: "127.0.0.1:8021",
			Timeout: 4 * time.Second,
			Retries: 0,
		},
	}
	return &c
}

//Read config from file.
func (p *FMconfig) Read(file string) error {
	var e error

	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	e = xml.Unmarshal(data, p)
	if e != nil {
		return e
	}
	return nil
}

//Write config to file.
func (p *FMconfig) Write(file string) error {
	var e error

	//f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.Create(file)
	defer f.Close()
	if err != nil {
		return err
	}
	f.Sync()
	return e
}
