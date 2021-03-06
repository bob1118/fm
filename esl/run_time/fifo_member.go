package run_time

import (
	"fmt"
	"log"

	"github.com/bob1118/fm/esl/eventsocket"
	"github.com/bob1118/fm/models"
)

//mod_fifo default queue is cool_fifo@$${domain}
//define queue fifomember@fifos and fifoconsumer@fifos;
//fifo member manage function, fifo_member add/fifo_member del

//FifoMemberManage
func FifoMemberManage(c *eventsocket.Connection, originate string, is bool) (e error) {
	var myerr error
	var apicmd string
	var op string

	condition := fmt.Sprintf("member_string='%s'", originate)
	if fifomembers, err := models.GetFifomembers(condition); err != nil {
		log.Println(err)
		return err
	} else {
		for _, fifomember := range fifomembers {
			if is {
				op = "fifo_member add"
			} else {
				op = "fifo_member del"
			}
			apicmd = fmt.Sprintf("api %s %s %s %s %s %s", op, fifomember.Fname, fifomember.Mstring, fifomember.Msimo, fifomember.Mtimeout, fifomember.Mlag)
			if _, err := c.Send(apicmd); err != nil {
				log.Println(err)
				myerr = err
			}
		}
	}
	return myerr
}
