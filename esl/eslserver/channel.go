package eslserver

import (
	"errors"
	"fmt"
	"log"

	"github.com/bob1118/fm/esl/eventsocket"
	"github.com/bob1118/fm/models"
	"github.com/bob1118/fm/utils"
)

////////////////////first event CHANNEL_DATA action///////////////////////////

//DefaultChannelAction
func ChannelDefaultAction(c *eventsocket.Connection, e *eventsocket.Event) error {
	var myerr error

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
	if myevent, err := c.Send("myevents"); err != nil {
		myerr = err
		log.Println(err)
	} else {
		myevent.LogPrint()
	}

	if utils.IsEqual(call.direction, "inbound") {
		switch call.profile {
		case "internal", "internal-ipv6": //internal ua incoming
			myerr = channelInternalProc(c, call)
		case "external", "external-ipv6": //external gateway incoming
			myerr = channelExternalProc(c, call)
		default:
			myerr = errors.New("CHANNEL_DATA:known profile")
		}
	} else {
		//outgoing hit socket?
		e.LogPrint()
	}
	return myerr
}

//channelInternalProc
func channelInternalProc(c *eventsocket.Connection, call *CALL) (err error) {
	var myerr error

	if call.CallerIsUa() {
		if call.CalleeIsUa() { //ua dial ua
			app := "bridge"
			appargv := fmt.Sprintf(`{origination_caller_id_name=%s,origination_caller_id_number=%s}sofia/%s/%s`, "local", call.ani, call.domain, call.distinationnumber)
			c.Execute("set", `hangup_after_bridge=true`, true)
			c.Execute(app, appargv, true)
		} else { //ua dial out through gateway.
			q := fmt.Sprintf(`account_id='%s' and account_domain='%s' and acce164_isdefault=true limit 1`, call.ani, call.domain)
			if acce164s, err := models.GetAcce164s(q); err != nil {
				myerr = err
				c.Execute("hangup", "INVALID_GATEWAY", true)
			} else {
				gatewayname := acce164s[0].Gname
				gatewaye164number := acce164s[0].Enumber
				appargv := fmt.Sprintf(`{origination_caller_id_number=%s,ignore_early_media=true,codec_string='PCMU,PCMA'}sofia/gateway/%s/%s`, gatewaye164number, gatewayname, call.distinationnumber)
				c.Execute("set", `hangup_after_bridge=true`, true)
				c.Execute("bridge", appargv, true)
			}
		}
	} else {
		c.Execute("hangup", "USER_NOT_REGISTERED", true)
	}
	return myerr
}

//channelExternalProc
func channelExternalProc(c *eventsocket.Connection, call *CALL) (err error) {

	if !call.CallFilterPassed() {
		c.Execute("hangup", "CALL_REJECT", true)
		return errors.New("function CallFilterPassed fail, Call Reject")
	} else {
		return channelExternalExecuteFifo(c)
	}
}

func channelExternalExecuteFifo(c *eventsocket.Connection) error {
	//Put a caller into a FIFO queue
	//<action application="fifo" data="myqueue in /tmp/exit-message.wav /tmp/music-on-hold.wav"/>
	argv := `fifomember@fifos in`
	if e, err := c.Execute(`fifo`, argv, true); err != nil {
		log.Println(err)
	} else {
		e.LogPrint()
	}
	return nil
}

////////////////////////channel event action////////////////////

func ChannelAction(c *eventsocket.Connection, e *eventsocket.Event) {
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
