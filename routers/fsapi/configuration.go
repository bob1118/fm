package fsapi

import (
	"github.com/gin-gonic/gin"
)

func doConfiguration(c *gin.Context) (b string) {
	body := Notfound
	value := c.PostForm("key_value")
	switch value {
	//switch boot order.
	// case "console.conf":
	// case "logfile.conf":
	// case "enum.conf":
	// case "xml_curl.conf":
	case "odbc_cdr.conf": //1th request.
		//body = odbc_cdr.ReadConfiguration(c)
	case "sofia.conf": //2th request(a request per profile).
	//case "loopback.conf": //3th
	//case "verto.conf": //4th
	case "conference.conf": //5th
	//case "db.conf": //6th
	case "fifo.conf": //7th
	case "hash.conf": //8th
	case "voicemail.conf": //9th
	//case "httapi.conf": //10th
	//case "spandsp.conf": //11th
	case "opus.conf": //12th
	case "avformat.conf": //13th
	case "avcodec.conf": //14th
	case "sndfile.conf": //15th
	case "local_stream.conf": //16th
	//case "lua.conf": //17th
	case "post_load_modules.conf": //18th
	case "acl.conf": //19th
	case "event_socket.conf": //20th
	case "post_load_switch.conf": //21th
	case "switch.conf": //22th

	//reloadxml
	case "timezones.conf":
	}

	return body
}
