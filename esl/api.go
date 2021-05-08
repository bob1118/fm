//Package esl send api command and wait return.
package esl

import "github.com/bob1118/fm/esl/eventsocket"

//APICreateUUID function. send "api create_uuid" and wait return uuid.
func APICreateUUID(c *eventsocket.Connection) (string, error) {
	var (
		err  error
		uuid string
	)
	if event, e := c.Send("api create_uuid"); e != nil {
		err = e
	} else {
		uuid = event.Body
	}
	return uuid, err
}
