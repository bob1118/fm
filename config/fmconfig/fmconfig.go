package fmconfig

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"time"
)

//CFGFILE file.
const CFGFILE = "fm.conf.xml"

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
	Host        string
	User        string
	Password    string
	Tableprefix string
}

//FMruntime struct {
type FMruntime struct {
	Runmode       string
	Jwtsecret     string
	Pagesize      uint
	Enablehash    bool
	ConfDirectory string
}

//FMesl struct {
type FMesl struct {
	Mode       string
	ServerAddr string
	Password   string
	Timeout    time.Duration
	Retries    uint
	ListenAddr string
}

//FMconfig struct {
type FMconfig struct {
	File     string `xml:"file,attr"`
	Server   FMhttpserver
	Database FMdatabase
	Runtime  FMruntime
	Esl      FMesl
}

var CFG *FMconfig

func init() {
	CFG = NewFmconfig()
	if CFG.Read() != nil {
		CFG.Write()
	}
}

//NewFmconfig return default fmconfig .
func NewFmconfig() *FMconfig {
	c := FMconfig{
		File: CFGFILE,
		Server: FMhttpserver{
			Address:      "127.0.0.1:80",
			Readtimeout:  4 * time.Second,
			Writetimeout: 4 * time.Second,
		},
		Database: FMdatabase{
			Type:        "postgres",
			Name:        "freeswitch",
			Host:        "127.0.0.1",
			User:        "fsdba",
			Password:    "fsdba",
			Tableprefix: "",
		},
		Runtime: FMruntime{
			Runmode:       "debug",
			Jwtsecret:     "abcdefghijklmnopqrstuvwxyz",
			Pagesize:      20,
			Enablehash:    true,
			ConfDirectory: "/etc/freeswitch/",
		},
		Esl: FMesl{
			Mode:       "inbound",
			ServerAddr: ":::8021",
			Password:   "ClueCon",
			Timeout:    4 * time.Second,
			Retries:    0,
			ListenAddr: "127.0.0.1:12345",
		},
	}
	return &c
}

//Read config from file.
func (p *FMconfig) Read() error {
	var e error

	f, err := os.Open(CFGFILE)
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
func (p *FMconfig) Write() error {
	var e error

	//f, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.Create(CFGFILE)
	defer f.Close()
	if err != nil {
		return err
	}

	//data, err := xml.Marshal(p)
	data, err := xml.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}

	f.WriteString(xml.Header)
	f.Write(data)
	f.Sync()
	return e
}
