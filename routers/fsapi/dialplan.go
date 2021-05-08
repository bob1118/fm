package fsapi

import (
	"fmt"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/gin-gonic/gin"
)

//doDialplan function return xml dialplan.
func doDialplan(c *gin.Context) (b string) {
	body := NOT_FOUND
	dpMode := fmconfig.CFG.Esl.Mode

	switch dpMode {
	case "inbound", "Inbound", "INBOUND":
		body = DialplanFreeswitchInbound
	case "outbound", "Outbound", "OUTBOUND":
		body = fmt.Sprintf(DialplanFreeswitchOutbound, fmconfig.CFG.Esl.ListenAddr)
	}
	return body
}
