package run_time

import (
	"log"
	"sync"

	"github.com/bob1118/fm/esl/eventsocket"
)

var MUseragent sync.Map
var MChannel sync.Map

//useragent
type rtUseragent struct {
	coreuuid    string
	hostipv4    string
	hostipv6    string
	profilename string
	callid      string
	expires     string
	user        string
	domain      string
}

//channels
type rtChannel struct{}

func Setuaonline(e *eventsocket.Event) {
	ua := rtUseragent{
		coreuuid:    e.Get("Core-Uuid"),
		hostipv4:    e.Get("Freeswitch-Ipv4"),
		hostipv6:    e.Get("Freeswitch-Ipv6"),
		profilename: e.Get("Profile-Name"),
		callid:      e.Get("Call-Id"),
		expires:     e.Get("Expires"),
		user:        e.Get("User_Name"),
		domain:      e.Get("Domain_Name"),
	}
	if _, isloaded := MUseragent.LoadOrStore(ua.callid, &ua); isloaded {
		MUseragent.Delete(ua.callid)
		MUseragent.Store(ua.callid, &ua)
	} else {
		log.Println(ua)
	}
}

func Setuaoffline(e *eventsocket.Event) {
	mycallid := e.Get("Call-Id")
	if ua, isloaded := MUseragent.LoadAndDelete(mycallid); isloaded {
		log.Println(ua)
	}
}
