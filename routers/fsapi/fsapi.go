package fsapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//PostFromXmlCurl function response all xml config.
func PostFromXmlCurl(c *gin.Context) {

	responseBody := NOT_FOUND
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
