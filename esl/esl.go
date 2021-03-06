package esl

import (
	"log"

	"github.com/bob1118/fm/esl/eslclient"
	"github.com/bob1118/fm/esl/eslserver"
)

func init() {}

//Run
//freeswitch inbound connector execute system api,channel application, and more.
//freeswitch outbound connector execute some channel application.
func Run(eslmode string) {
	switch eslmode {
	case "inbound", "Inbound", "INBOUND":
		eslclient.ClientRun()
	case "outbound", "Outbound", "OUTBOUND":
		eslserver.ServerRun()
	default:
		log.Println("known esl mode")
	}
}
