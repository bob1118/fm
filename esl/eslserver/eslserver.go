//eslserver is a tcp server for dialplan application socket
//while mod_sofia receive a incoming call, dialplan execute socket(ip:port async full) and socket connect to eslserver.

package eslserver

import (
	"log"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/esl/eventsocket"
)

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
	return ChannelDefaultAction(c, e)
}

//eventReadAChannelLoop ChannelAction
func eventReadAChannelLoop(c *eventsocket.Connection) error {
	isLoop := true
	for isLoop {
		if e, err := c.ReadEvent(); err != nil {
			return err
		} else {
			ChannelAction(c, e)
		}
	}
	return nil
}
