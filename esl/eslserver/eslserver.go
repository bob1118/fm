//eslserver is a tcp server for dialplan application socket
//while mod_sofia receive a incoming call, dialplan execute socket(ip:port async full) and socket connect to eslserver.

package eslserver

import (
	"errors"
	"log"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/esl/eventsocket"
	"github.com/bob1118/fm/esl/run_time"
)

type CALL struct {
	//CALL BASIC INFO
	coreuuid          string //Core-Uuid
	fsipv4            string //Freeswitch-Ipv4
	eventname         string //Event-Name
	uuid              string //Variable_uuid
	callid            string //Variable_call_uuid
	direction         string //Variable_direction
	profile           string //Variable_sofia_profile_name
	domain            string //Variable_domain_name
	gateway           string //Variable_sip_gateway
	ani               string //Caller-Ani
	distinationnumber string //Caller-Destination-Number
}

func init() {}

func (c *CALL) CallerIsUa() bool {
	mkey := c.ani + "@" + c.domain
	return run_time.IsUa(mkey)
}
func (c *CALL) CalleeIsUa() bool {
	mkey := c.distinationnumber + "@" + c.domain
	return run_time.IsUa(mkey)
}

func ServerRun() {
	if err := eventsocket.ListenAndServe(fmconfig.CFG.Esl.ListenAddr, handler); err != nil {
		log.Println(err)
	}
}

func ServerRestart() {}

func handler(c *eventsocket.Connection) {
	log.Println("new client:", c, "from:", c.RemoteAddr())
	if e, err := c.Send("connect"); err != nil {
		log.Println(err)
	} else {
		eventChannelDefaultAction(c, e)
		eventReadAChannelLoop(c)
	}
}

//incoming call CHANNEL_DATA
func eventChannelDefaultAction(c *eventsocket.Connection, e *eventsocket.Event) (err error) {
	return DialplanExecute(c, e)
}

//DialplanExecute
func DialplanExecute(c *eventsocket.Connection, e *eventsocket.Event) error {
	var myerr error

	e.LogPrint()
	call := &CALL{
		coreuuid:          e.Get("Core-Uuid"),
		fsipv4:            e.Get("Freeswitch-Ipv4"),
		eventname:         e.Get("Event-Name"),
		uuid:              e.Get("Variable_uuid"),
		callid:            e.Get("Variable_call_uuid"),
		direction:         e.Get("Variable_direction"),
		profile:           e.Get("Variable_sofia_profile_name"),
		domain:            e.Get("Variable_domain_name"),
		gateway:           e.Get("Variable_sip_gateway"),
		ani:               e.Get("Caller-Ani"),
		distinationnumber: e.Get("Caller-Destination-Number"),
	}

	//send myevents
	if e, err := c.Send("myevents"); err != nil {
		myerr = err
		log.Println(err)
	} else {
		e.LogPrint()
	}

	switch call.profile {
	case "internal", "internal-ipv6": //internal ua incoming
		myerr = channelInternalProc(c, call)
	case "external", "external-ipv6": //external gateway incoming
		myerr = channelExternalProc(c, call)
	default:
		myerr = errors.New("CHANNEL_DATA:known profile")
	}
	return myerr
}

//eventReadAChannelLoop
func eventReadAChannelLoop(c *eventsocket.Connection) error {
	isLoop := true
	for isLoop {
		if e, err := c.ReadEvent(); err != nil {
			return err
		} else {
			eventChannelAction(c, e)
		}
	}
	return nil
}

func eventChannelAction(c *eventsocket.Connection, e *eventsocket.Event) {
	e.LogPrint()
	eventName := e.Get("Event-Name")
	if len(eventName) > 0 {
		switch eventName {
		case "BACKGROUND_JOB":
			backgroundjobAction(c, e)
		case "CHANNEL_STATE":
			channelstateAction(c, e)
		case "CHANNEL_CALLSTATE":
			channelcallstateAction(c, e)
		case "CHANNEL_HANGUP":
			channelhangupAction(c, e)
		case "CHANNEL_DESTROY":
			channelCDRAction(c, e)
		default:
			//nothing todo.
		}
	}
}
