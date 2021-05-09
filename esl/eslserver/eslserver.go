package eslserver

import (
	"log"
	"sync"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/esl/eventsocket"
)

var ServerCon sync.Map

func init() {}

func ServerRun() {
	eventsocket.ListenAndServe(fmconfig.CFG.Esl.ListenAddr, handler)
}

func ServerRestart() {}

func handler(c *eventsocket.Connection) {
	log.Println("new client:", c, "from:", c.RemoteAddr())
	if e, err := c.Send("connect"); err != nil {
		log.Println(err)
	} else {
		log.Println(e)
		eventReadAChannelLoop(c)
	}
}

//eventReadAChannelLoop
func eventReadAChannelLoop(c *eventsocket.Connection) error {
	isLoop := true
	for isLoop {
		if e, err := c.ReadEvent(); err != nil {
			log.Println(err)
			return err
		} else {
			eventChannelAction(c, e)
		}
	}
	return nil
}

func eventChannelAction(c *eventsocket.Connection, e *eventsocket.Event) {
	if eventName, ok := e.Header["Event-Name"].(string); ok {
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
func channelstateAction(c *eventsocket.Connection, e *eventsocket.Event) {}

//channelcallstateAction function.
func channelcallstateAction(c *eventsocket.Connection, e *eventsocket.Event) {}

//channelhangupAction function.
func channelhangupAction(c *eventsocket.Connection, e *eventsocket.Event) {}

//channelCDRAction function. channel cdr.
func channelCDRAction(c *eventsocket.Connection, e *eventsocket.Event) {}
