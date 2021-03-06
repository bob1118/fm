//freeswitch run bgapi command do a backgroud job.
//job down while receive BACKGROUND_JOB event.
package bgapicmd

import (
	"github.com/bob1118/fm/esl/eslclient"
	"github.com/gin-gonic/gin"
)

//some freeswitch bgapi/cmd reply.
//request: Get /bgapi?cmd=xxx
//response: job uuid.
func Get(c *gin.Context) {
	var cmd, result string
	cmd = c.Query("cmd")
	if jobuuid, err := eslclient.SendBgapiCommand(cmd); err != nil {
		result = err.Error()
	} else {
		result = jobuuid
	}
	c.String(200, result)
}
