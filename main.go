package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/routers"
)

func main() {
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
	go func() {
		//
	}()

	//do licence check every 1second.
	for now := range time.Tick(10 * time.Second) {
		fmt.Println(now)
	}
}
