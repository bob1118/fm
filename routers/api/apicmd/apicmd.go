package apicmd

import (
	"github.com/bob1118/fm/esl/eslclient"
	"github.com/gin-gonic/gin"
)

//some freeswitch api/cmd response.
//request: Get /apicmd?c=xxx
//response: string
func Get(c *gin.Context) {
	var cmd, result string
	cmd = c.Query("c")
	result = eslclient.SendApiCommand(cmd)
	c.String(200, result)
}
