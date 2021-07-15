package eslclient

import (
	"fmt"
	"log"

	"github.com/bob1118/fm/esl/eventsocket"
	"github.com/bob1118/fm/esl/run_time"
	"github.com/bob1118/fm/models"
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
	job := &models.Bgjob{
		Juuid:    e.Get("Job-Uuid"),
		Jcmd:     e.Get("Job-Command"),
		Jcmdarg:  e.Get("Job-Command-Arg"),
		Jcontent: e.Body,
	}
	if false ||
		len(job.Juuid) == 0 ||
		len(job.Jcmd) == 0 ||
		len(job.Jcontent) == 0 {
		log.Println(e)
	} else {
		models.CreateBgjob(job)
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
