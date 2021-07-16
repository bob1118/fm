////////////////////////////////////////////
//freeswitch outgoing call
//bgapi originate ...
////////////////////////////////////////////
//https://freeswitch.org/confluence/display/FREESWITCH/mod_commands
//originate <call_url> <exten>|&<application_name>(<app_args>) [<dialplan>] [<context>] [<cid_name>] [<cid_num>] [<timeout_sec>]
//
//originate {origination_caller_id_number=9005551212}sofia/default/whatever@wherever 19005551212 XML default
//originate {ignore_early_media=true,originate_timeout=60}sofia/gateway/name/number &playback(message)
//originate {origination_caller_id_number=2024561000}sofia/gateway/whitehouse.gov/2125551212 &bridge({effective_caller_id_number=7036971379}sofia/gateway/pentagon.gov/3035554499)

package call

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bob1118/fm/ec"
	"github.com/bob1118/fm/esl/eslclient"
	"github.com/gin-gonic/gin"
)

type GATEWAYLEG struct {
	gatewayname    string `json:"gwname"`
	calleridnumber string `json:"gwcaller"`
	calleeidnumber string `json:"gwcallee"`
}

type USERAGENTLEG struct {
	uadomain       string `json:"uadomain"`
	calleridnumber string `json:"uacaller"`
	calleeidnumber string `json:"uacallee"`
}

type CALLGWBRIDGEUA struct {
	aleg GATEWAYLEG
	bleg USERAGENTLEG
}

func (c CALLGWBRIDGEUA) emptyCheck() bool {
	if false ||
		c.aleg.gatewayname == "" ||
		c.aleg.calleridnumber == "" ||
		c.aleg.calleeidnumber == "" ||
		c.bleg.uadomain == "" ||
		c.bleg.calleridnumber == "" ||
		c.bleg.calleeidnumber == "" {
		return true
	} else {
		return false
	}
}

//PostDial function.
func PostDial(c *gin.Context) {}

//PostOriginateExecuteDialplanExtension function.
func PostOriginateExecuteDialplanExtension(c *gin.Context) {}

//PostOriginateExecuteBridge function.
func PostOriginateExecuteApp(c *gin.Context) {}

//PostOriginateExecuteBridge function.
//request: POST /call
//response: job uuid
func PostOriginateExecuteBridge(c *gin.Context) {
	var data string
	code := ec.SUCCESS
	call := CALLGWBRIDGEUA{}

	if err := c.BindJSON(&call); err != nil {
		code = ec.ERROR_BIND_JSON
	} else {
		if call.emptyCheck() {
			code = ec.ERROR_PARAM_NULL
		} else {
			cmd := fmt.Sprintf("originate {origination_caller_id_number=%s,codec_string='PCMU,PCMA',ignore_early_media=true}sofia/gateway/%s/%s &bridge({origination_caller_id_number=%s}sofia/%s/%s)",
				call.aleg.calleridnumber, call.aleg.gatewayname, call.aleg.calleeidnumber, call.bleg.calleridnumber, call.bleg.uadomain, call.bleg.calleeidnumber)
			log.Println(cmd)
			if jobuuid, err := eslclient.SendBgapiCommand(cmd); err != nil {
				code = ec.ERROR_BGAPI_ORIGINATE
				data = err.Error()
			} else {
				data = jobuuid
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
	})
}
