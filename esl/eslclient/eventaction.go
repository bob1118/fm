package eslclient

import (
	"fmt"

	"github.com/bob1118/fm/esl/eventsocket"
	"github.com/bob1118/fm/esl/run_time"
)

//eventAction function.
func eventAction(e *eventsocket.Event) {
	eventName := e.Get("Event-Name")
	if len(eventName) > 0 {
		switch eventName {
		case "API":
			apiAction(e)
		case "CUSTOM":
			customAction(e)
		case "BACKGROUND_JOB":
			backgroundjobAction(e)
		case "CHANNEL_STATE":
			channelstateAction(e)
		case "CHANNEL_CALLSTATE":
			channelcallstateAction(e)
		case "CHANNEL_HANGUP":
			channelhangupAction(e)
		case "CHANNEL_DESTROY":
			channelCDRAction(e)
		default:
			//nothing todo.
		}
	}
}

func apiAction(e *eventsocket.Event) {}

func customAction(e *eventsocket.Event) {
	user := e.Get("User_Name")
	domain := e.Get("Domain_Name")
	eventsubclass := e.Get("Event-Subclass")

	if len(eventsubclass) > 0 {
		switch eventsubclass {
		case "sofia::pre_register", "sofia::register_attempt", "sofia::register_failure": //sofia_reg_handle_register_token
		case "sofia::register": //sofia_reg_handle_register_token
			if len(user) > 0 && len(domain) > 0 {
				originate_string := fmt.Sprintf(`sofia/%s/%s`, domain, user)
				run_time.SetUaOnline(e)
				run_time.FifoMemberManage(ClientCon, originate_string, true)
			}
		case "sofia::unregister": //sofia_reg_handle_register_token
			if len(user) > 0 && len(domain) > 0 {
				originate_string := fmt.Sprintf(`sofia/%s/%s`, domain, user)
				run_time.SetUaOffline(e)
				run_time.FifoMemberManage(ClientCon, originate_string, false)
			}
		case "sofia::expire": //sofia_reg_del_call_back
			if len(user) > 0 && len(domain) > 0 {
				originate_string := fmt.Sprintf(`sofia/%s/%s`, domain, user)
				run_time.SetUaOffline(e)
				run_time.FifoMemberManage(ClientCon, originate_string, false)
			}
		case "sofia::gateway_state": //sofia_reg_fire_custom_gateway_state_event
			run_time.SetGatewayState(e)
		default:
		}
	}
}

//backgroundjobAction function.
func backgroundjobAction(e *eventsocket.Event) {
	bgcommand := e.Get("Job-Command")
	if len(bgcommand) > 0 {
		switch bgcommand {
		case "originate", "Originate", "ORIGINATE":
			// in := models.BackgroundJob{}
			// in.BgjobUUID = e.Header["Job-Uuid"].(string)
			// in.BgjobCommand = bgcommand
			// in.BgjobCmdArg = e.Header["Job-Command-Arg"].(string)
			// //+OK 9fbc526c-80c2-49c8-bc2d-9735872dfa53//-ERR UNALLOCATED_NUMBER
			// body := strings.Split(e.Body, " ")
			// in.BgjobResult = strings.TrimSpace(body[1])
			// switch body[0] {
			// case "+OK":
			// 	in.BgjobReturn = true
			// case "-ERR":
			// 	in.BgjobReturn = false
			// }
			// if err := models.CreateBgjob(&in); err != nil {
			// 	log.Println(err)
			// }
		case "command":
			//todo.
		default:
			//todo.
		}
	}
}

//channelstateAction function.
func channelstateAction(e *eventsocket.Event) {
}

//channelcallstateAction function.
func channelcallstateAction(e *eventsocket.Event) {}

//channelhangupAction function.
func channelhangupAction(e *eventsocket.Event) {
	//	if uuid, ok = e.Header["variable_uuid"].(string); ok {
	//	}
}

//channelCDRAction function. channel cdr.
func channelCDRAction(e *eventsocket.Event) {

}
