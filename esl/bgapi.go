package esl

import "github.com/bob1118/fm/esl/eventsocket"

//BGAPIOriginateAndBridge function. send bgapi originate a and bridge b.
func BGAPIOriginateAndBridge(c *eventsocket.Connection, cmd string) (string, error) {
	return BGAPIOriginate(c, cmd)
}

//BGAPIOriginate function. send bgapi originate and wait return job-uuid.
func BGAPIOriginate(c *eventsocket.Connection, cmd string) (string, error) {
	var (
		jobuuid string
		err     error
	)
	if event, e := c.Send(cmd); e != nil {
		err = e
	} else {
		jobuuid = event.Get("Job-Uuid")
	}
	return jobuuid, err
}
