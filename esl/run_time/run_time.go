package run_time

import (
	"log"
	"sync"

	"github.com/bob1118/fm/esl/eventsocket"
)

var uamap sync.Map
var chmap sync.Map

//runtime useragent
type rtua struct {
	coreuuid    string
	hostipv4    string
	hostipv6    string
	profilename string
	callid      string
	expires     string
	user        string
	domain      string
}

//runtime channel
type rtch struct{}

func SetUaOnline(e *eventsocket.Event) {
	ua := rtua{
		coreuuid:    e.Get("Core-Uuid"),
		hostipv4:    e.Get("Freeswitch-Ipv4"),
		hostipv6:    e.Get("Freeswitch-Ipv6"),
		profilename: e.Get("Profile-Name"),
		callid:      e.Get("Call-Id"),
		expires:     e.Get("Expires"),
		user:        e.Get("User_Name"),
		domain:      e.Get("Domain_Name"),
	}

	k := ua.user + "@" + ua.domain
	if _, isloaded := uamap.LoadOrStore(k, &ua); isloaded {
		uamap.Delete(k)
		uamap.Store(k, &ua)
		log.Println("SetUaOnline: ua is loaded, delete and store")
	} else {
		log.Println(ua)
	}
}

func SetUaOffline(e *eventsocket.Event) {
	k := e.Get("User_Name") + "@" + e.Get("Domain_Name")
	if ua, isloaded := uamap.LoadAndDelete(k); isloaded {
		log.Println(ua)
	}
}

func IsUa(k interface{}) bool {
	_, ok := uamap.Load(k)
	return ok
}

func uamapclean() {
	uamap.Range(func(k, v interface{}) bool { uamap.Delete(k); return true })
}
func chmapclean() {
	chmap.Range(func(k, v interface{}) bool { chmap.Delete(k); return true })
}
