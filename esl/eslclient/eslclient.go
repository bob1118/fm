//eslclient is a tcp client connect to mod_evnet_socket.
//while mod_sofia receive a incoming call, dialplan execute app park.
//now, do what you want to before received park execute complete event.

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
	var e error
	alwaysrun := true
	//	if isrun, ok := <-CHfsisrun; ok {
	//		if isrun {
	for alwaysrun {
		log.Println(time.Now(), "->start reconnect.")
		c, err := eventsocket.Dial(fmconfig.CFG.Esl.ServerAddr, fmconfig.CFG.Esl.Password)
		if err != nil {
			if errors.Is(err, syscall.WSAECONNRESET+7) { //syscall.Errno=10061 (No connection could be made because the target machine actively refused it)
				log.Println(time.Now(), err)
				e = err
			}
		} else {
			ClientCon = c
			if eventSubscribe("plain") &&
				eventUnsubscribe("plain", "RE_SCHEDULE", "HEARTBEAT", "MESSAGE_WAITING", "MESSAGE_QUERY") { // nixevent RE_SCHEDULE HEARTBEAT MESSAGE_WAITING MESSAGE_QUERY
				if err := eventReadLoop(); err != nil {
					e = err
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
	}
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

//eslclient send api command return api response body.
func SendApiCommand(cmd string) string {
	// api version

	// Content-Type: api/response
	// Content-Length: 85

	// FreeSWITCH Version 1.10.6-release-18-1ff9d0a60e~64bit (-release-18-1ff9d0a60e 64bit)
	var response string
	if len(cmd) > 0 {
		apicommand := "api" + " " + cmd
		if ev, err := ClientCon.Send(apicommand); err != nil {
			response = err.Error()
		} else {
			ev.LogPrint()
			response = ev.Body
		}
	}
	return response
}

//eslclient send bgapi command return command reply Job-UUID.
func SendBgapiCommand(cmd string) (uuid string, e error) {
	// 	bgapi version

	// Content-Type: command/reply
	// Reply-Text: +OK Job-UUID: 2b5563b2-d465-4d90-8abd-52a032d0933f
	// Job-UUID: 2b5563b2-d465-4d90-8abd-52a032d0933f

	var jobuuid string
	if len(cmd) > 0 {
		bgapicommand := "bgapi" + " " + cmd
		if ev, err := ClientCon.Send(bgapicommand); err != nil {
			return "", err
		} else {
			jobuuid = ev.Get("Job-Uuid")
		}
	}
	return jobuuid, nil
}
