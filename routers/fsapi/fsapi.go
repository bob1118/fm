package fsapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostFromXmlCurl function
func PostFromXmlCurl(c *gin.Context) {

	responseBody := Notfound
	value := c.PostForm("section")
	switch value {
	case "configuration":
		responseBody = doConfiguration(c)
	case "dialplan":
		responseBody = doDialplan(c)
	case "directory":
		responseBody = doDirectory(c)
	case "phrases":
		responseBody = doPhrases(c)
	}
	c.String(http.StatusOK, responseBody)
}
