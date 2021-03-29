package fsapi

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const Notfound = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<document type="freeswitch/xml">
	<section name="result">
		<result status="not found"/>
	</section>
</document>`

//PostFromXmlCurl function
func PostFromXmlCurl(c *gin.Context) {

	responseBody := ""
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

//string compare case-insensitivity
func equal(s string, d string) (b bool) {
	return strings.EqualFold(s, d)
}
