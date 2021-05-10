package eslclient

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
	"syscall"
	"time"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/esl/eventsocket"
)

var CHfsisrun chan bool

var ClientCon *eventsocket.Connection

func init() { CHfsisrun = make(chan bool) }

//clientRun
func ClientRun() {
	if err := clientReconnect(); err != nil {
		log.Println(err)
	}
}

//clientReconnect
func clientReconnect() error {
	//var e error
	//	if isrun, ok := <-CHfsisrun; ok {
	//		if isrun {
	log.Println(time.Now(), "->start reconnect.")
	c, err := eventsocket.Dial(fmconfig.CFG.Esl.ServerAddr, fmconfig.CFG.Esl.Password)
	if err != nil {
		if errors.Is(err, syscall.WSAECONNRESET+7) { //syscall.Errno=10061 (No connection could be made because the target machine actively refused it)
			log.Println(time.Now(), err)
		}
	} else {
		ClientCon = c
		if eventSubscribe("plain") &&
			eventUnsubscribe("plain", "RE_SCHEDULE", "HEARTBEAT", "MESSAGE_WAITING", "MESSAGE_QUERY") { // RE_SCHEDULE HEARTBEAT MESSAGE_WAITING MESSAGE_QUERY
			if err := eventReadLoop(); err != nil {
				if errors.Is(err, io.EOF) {
					log.Println(time.Now(), err)
				}
				if errors.Is(err, syscall.WSAECONNRESET) { //windows
					log.Println(time.Now(), err)
				}
				if errors.Is(err, syscall.ECONNRESET) { //linux
					log.Println(time.Now(), err)
				}
			}
		}
	}
	//		}
	//	}
	<-time.After(8 * time.Second)
	e := clientReconnect()
	return e
}

//EventLoop function.
func eventReadLoop() error {
	isLoop := true
	for isLoop {
		if e, err := ClientCon.ReadEvent(); err != nil {
			return err
		} else {
			eventAction(e)
		}
	}
	return nil
}

//eventSubscribe function.
func eventSubscribe(format string, enames ...string) bool {
	var isOK bool
	var command string

	command = fmt.Sprintf("event %s", format)
	if len(enames) == 0 {
		command += " all"
	} else {
		for _, ename := range enames {
			command += fmt.Sprintf(" %s", ename)
		}
	}

	if event, err := ClientCon.Send(command); err != nil {
		isOK = false
		log.Println(err)
	} else {
		reply := event.Header["Reply-Text"]
		if strings.Contains(reply.(string), "+OK") {
			isOK = true
		}
	}
	return isOK
}

//eventUnsubscribe function.
func eventUnsubscribe(format string, enames ...string) bool {
	var isOK bool
	var command string

	command = fmt.Sprintf("nixevent %s", format)
	if len(enames) == 0 {
		command = "noevents"
	} else {
		for _, ename := range enames {
			command += fmt.Sprintf(" %s", ename)
		}
	}

	if event, err := ClientCon.Send(command); err != nil {
		isOK = false
		log.Println(err.Error())
	} else {
		reply := event.Header["Reply-Text"]
		if strings.Contains(reply.(string), "+OK") {
			isOK = true
		}
	}
	return isOK
}

//eventAction function.
func eventAction(e *eventsocket.Event) {
	e.LogPrint()
	if eventName, ok := e.Header["Event-Name"].(string); ok {
		switch eventName {
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

//backgroundjobAction function.
func backgroundjobAction(e *eventsocket.Event) {
	if bgcommand, ok := e.Header["Job-Command"].(string); ok {
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
