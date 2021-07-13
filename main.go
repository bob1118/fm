package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/esl"
	"github.com/bob1118/fm/models"
	"github.com/bob1118/fm/routers"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
)

func main() {

	models.Init()
	xmlbuilder.BuildPersonalConf()

	//do http api gateway
	go func() {
		h := routers.NewRouter()
		s := &http.Server{
			Addr:           fmconfig.CFG.Server.Address,
			Handler:        h,
			ReadTimeout:    fmconfig.CFG.Server.Readtimeout,
			WriteTimeout:   fmconfig.CFG.Server.Writetimeout,
			MaxHeaderBytes: 1 << 20,
		}
		s.ListenAndServe()
	}()

	//do freeswitch eventsocket msg analyze.
	go esl.Run("inbound")
	go esl.Run("outbound")

	for now := range time.Tick(8 * time.Second) {
		fmt.Println(now)
	}
}
