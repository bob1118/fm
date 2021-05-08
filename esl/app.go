//Package esl app.go sendmsg(execute mod_dptools) and wait return ok or not, try execute mod_dptool while sendmsg return ok.
package esl

import (
	"strings"

	"github.com/bob1118/fm/esl/eventsocket"
)

//APPDptoolsPark function.
func APPDptoolsPark(c *eventsocket.Connection, uuid string) bool {
	var isOK bool
	if event, err := c.SendMsg(eventsocket.MSG{
		"call-command":     "execute",
		"execute-app-name": "park",
	}, uuid, ""); err != nil {
		isOK = false
	} else {
		reply := event.Header["Reply-Text"]
		if strings.Contains(reply.(string), "+OK") {
			isOK = true
		}
	}
	return isOK
}

// APPDptoolsSet function.
func APPDptoolsSet(c *eventsocket.Connection, uuid string, appdata string) bool {
	var isOK bool
	if event, err := c.SendMsg(eventsocket.MSG{
		"call-command":     "execute",
		"execute-app-name": "set",
		"execute-app-arg":  appdata,
	}, uuid, ""); err != nil {
		isOK = false
	} else {
		reply := event.Header["Reply-Text"]
		if strings.Contains(reply.(string), "+OK") {
			isOK = true
		}
	}
	return isOK
}

//APPDptoolsRecord function.
func APPDptoolsRecord(c *eventsocket.Connection, uuid string, appdata string) bool {
	var isOK bool
	if event, err := c.SendMsg(eventsocket.MSG{
		"call-command":     "execute",
		"execute-app-name": "record",
		"execute-app-arg":  appdata,
	}, uuid, ""); err != nil {
		isOK = false
	} else {
		reply := event.Header["Reply-Text"]
		if strings.Contains(reply.(string), "+OK") {
			isOK = true
		}
	}
	return isOK
}

//APPDptoolsRecordSession function.
func APPDptoolsRecordSession(c *eventsocket.Connection, uuid string, appdata string) bool {
	var isOK bool
	if event, err := c.SendMsg(eventsocket.MSG{
		"call-command":     "execute",
		"execute-app-name": "record_session",
		"execute-app-arg":  appdata,
	}, uuid, ""); err != nil {
		isOK = false
	} else {
		reply := event.Header["Reply-Text"]
		if strings.Contains(reply.(string), "+OK") {
			isOK = true
		}
	}
	return isOK
}

//APPDptoolsBridge function.
func APPDptoolsBridge(c *eventsocket.Connection, uuid string, bleg string) bool {
	var isOK bool
	return isOK
}
