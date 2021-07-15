//freeswitch outgoing call
//bgapi originate arg &app arg
package call

import "github.com/gin-gonic/gin"

type DIAL struct {
	//DIAL BASIC INFO
	cmd string //originate

}

//PostDial function.
//request: POST /call
//response: job uuid
func PostDial(c *gin.Context) {}
