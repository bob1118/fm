package eslserver

import (
	"log"
	"sync"

	"github.com/bob1118/fm/esl/eventsocket"
)

var ServerCon sync.Map

func init() {}

func ServerRun() {
	eventsocket.ListenAndServe("", handler)
}

func ServerRestart() {}

func handler(c *eventsocket.Connection) {
	log.Println("new client:", c.RemoteAddr())
	if e, err := c.Send("connect"); err != nil {
		log.Println(err)
	} else {
		log.Println(e)
	}

}
