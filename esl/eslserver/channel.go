package eslserver

import (
	"fmt"

	"github.com/bob1118/fm/esl/eventsocket"
	"github.com/bob1118/fm/models"
)

////////////////////first event CHANNEL_DATA action///////////////////////////

//channelInternalProc
func channelInternalProc(c *eventsocket.Connection, call *CALL) (err error) {
	var myerr error

	if call.CallerIsUa() {
		if call.CalleeIsUa() { //ua dial ua
			app := "bridge"
			appargv := fmt.Sprintf(`{origination_caller_id_name=%s,origination_caller_id_number=%s}sofia/%s/%s`, "local", call.ani, call.domain, call.distinationnumber)
			c.Execute("set", `hangup_after_bridge=true`, true)
			c.Execute(app, appargv, true)
		} else { //ua dial out with gateway.
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
	//default fifo init.
	//c.Execute("set", `continue_on_fail=true`, true)
	//c.Execute("set", `hangup_after_bridge=true`, true)

	return nil
}

////////////////////////channel event action////////////////////

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
