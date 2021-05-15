//eslserver is a tcp server for dialplan application socket
//while mod_sofia receive a incoming call, dialplan execute socket(ip:port async full) and socket connect to eslserver.

package eslserver

import (
	"fmt"
	"log"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/esl/eventsocket"
	"github.com/bob1118/fm/esl/run_time"
	"github.com/bob1118/fm/models"
	"github.com/bob1118/fm/utils"
)

type CALL struct {
	eventname         string
	coreuuid          string
	fsipv4            string
	uuid              string
	callid            string
	direction         string
	profile           string
	domain            string
	ani               string
	distinationnumber string
}

func init() {}

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
func eventChannelDefaultAction(c *eventsocket.Connection, e *eventsocket.Event) {
	DialplanExecute(c, e)
}

//DialplanExecute
func DialplanExecute(c *eventsocket.Connection, e *eventsocket.Event) error {

	call := &CALL{
		eventname:         e.Get("Event-Name"),
		coreuuid:          e.Get("Variable_core-uuid"),
		fsipv4:            e.Get("Variable_freeswitch-ipv4"),
		uuid:              e.Get("Variable_uuid"),
		callid:            e.Get("Variable_sip_call_id"),
		direction:         e.Get("Variable_direction"),
		profile:           e.Get("Variable_sofia_profile_name"),
		domain:            e.Get("Variable_domain_name"),
		ani:               e.Get("Caller-Ani"),
		distinationnumber: e.Get("Caller-Destination-Number"),
	}

	//send myevents
	if e, err := c.Send("myevents"); err != nil {
		log.Println(err)
	} else {
		e.LogPrint()
	}

	//bridge
	if utils.IsEqual(call.profile, "internal") { //ua incoming
		if call.CallerIsUa() { //true { //
			if call.CalleeIsUa() { //dial local domain.
				app := "bridge"
				appargv := fmt.Sprintf(`{origination_caller_id_name=%s,origination_caller_id_number=%s}sofia/%s/%s`, "local", call.ani, call.domain, call.distinationnumber)
				if e, err := c.Execute(app, appargv, true); err != nil {
					log.Println(err)
				} else {
					e.LogPrint()
				}
			} else { //dial out
				// key := call.ani + "@" + call.domain
				// gatewayname := run_time.GetUaDefaultE164Number(key)
				// gatewaye164number := run_time.GetUaDefaultE164Number(key)
				// appargv := fmt.Sprintf(`{origination_caller_id_number=%s}sofia/gateway/%s/%s`, gatewaye164number, gatewayname, call.distinationnumber)
				// c.Execute("bridge", appargv, true)
				q := fmt.Sprintf("account_id=%s and account_domain=%s and acce164_isdefault=true limit 1", call.ani, call.domain)
				if acce164s, err := models.GetAcce164s(q); err != nil {
					gatewayname := acce164s[0].Gname
					gatewaye164number := acce164s[0].Enumber
					appargv := fmt.Sprintf(`{origination_caller_id_number=%s}sofia/gateway/%s/%s`, gatewaye164number, gatewayname, call.distinationnumber)
					c.Execute("bridge", appargv, true)
				} else {
					c.Execute("hangup", "INVALID_GATEWAY", true)
				}
			}
		} else {
			c.Execute("hangup", "USER_NOT_REGISTERED", true)
		}
	}
	if utils.IsEqual(call.profile, "external") { //gateway incoming
	}
	return nil
}

func (c *CALL) CallerIsUa() bool {
	mkey := c.ani + "@" + c.domain
	return run_time.IsUa(mkey)
}
func (c *CALL) CalleeIsUa() bool {
	mkey := c.distinationnumber + "@" + c.domain
	return run_time.IsUa(mkey)
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

//backgroundjobAction function.
func backgroundjobAction(c *eventsocket.Connection, e *eventsocket.Event) {
	bgcommand := e.Get("Job-Command")
	if len(bgcommand) > 0 {
		switch bgcommand {
		case "originate", "Originate", "ORIGINATE":
		case "command":
			//todo.
		default:
			//todo.
		}
	}
}

//channelstateAction function.
func channelstateAction(c *eventsocket.Connection, e *eventsocket.Event) {}

//channelcallstateAction function.
func channelcallstateAction(c *eventsocket.Connection, e *eventsocket.Event) {}

//channelhangupAction function.
func channelhangupAction(c *eventsocket.Connection, e *eventsocket.Event) {}

//channelCDRAction function. channel cdr.
func channelCDRAction(c *eventsocket.Connection, e *eventsocket.Event) {}
