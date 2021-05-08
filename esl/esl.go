package esl

import (
	"log"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/esl/eslclient"
	"github.com/bob1118/fm/esl/eslserver"
)

func init() {}

//Run
func Run() {
	var eslmode = fmconfig.CFG.Esl.Mode
	switch eslmode {
	case "inbound", "Inbound", "INBOUND":
		eslclient.ClientRun()
	case "outbound", "Outbound", "OUTBOUND":
		eslserver.ServerRun()
	default:
		log.Println("known esl mode")
	}

}
