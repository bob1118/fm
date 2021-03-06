package fsapi

import (
	"fmt"

	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/acl"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/db"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/event_socket"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/fifo"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/local_stream"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/odbc_cdr"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/sofia"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/switch_main"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder/autoload/voicemail"
	"github.com/gin-gonic/gin"
)

// configureation request when switch boot ...
//1.1
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=odbc_cdr.conf]
//1.2
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A26&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A26%20GMT&Event-Date-Timestamp=1618565366035759&Event-Calling-File=sofia.c&Event-Calling-Function=config_sofia&Event-Calling-Line-Number=4491&Event-Sequence=27]
//
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A27&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A27%20GMT&Event-Date-Timestamp=1618565367391604&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=29&profile=external-ipv6]
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A28&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A28%20GMT&Event-Date-Timestamp=1618565368557666&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=32&profile=external]
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A29&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A29%20GMT&Event-Date-Timestamp=1618565369574163&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=36&profile=internal-ipv6]
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A30&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A30%20GMT&Event-Date-Timestamp=1618565370358404&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=40&profile=internal]
//1.3
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=conference.conf&Event-Name=COMMAND&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A32&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A32%20GMT&Event-Date-Timestamp=1618565372760212&Event-Calling-File=mod_conference.c&Event-Calling-Function=send_presence&Event-Calling-Line-Number=3878&Event-Sequence=214&presence=true]
//1.4
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=fifo.conf]
//1.5
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=hash.conf]
//1.6
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=voicemail.conf]
//1.7
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=opus.conf]
//1.8
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=avformat.conf]
//1.9
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=avcodec.conf]
//1.10
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sndfile.conf]
//1.11
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=local_stream.conf]
//1.12
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=post_load_modules.conf]
//1.13
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=acl.conf]
//1.14
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=event_socket.conf]
//1.15
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=post_load_switch.conf]
//1.16
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=switch.conf]
// reload mod
//2.1
// reloadxml
//3.1 timezones only.

//doConfiguration function return xml config.
func doConfiguration(c *gin.Context) (b string) {
	body := NOT_FOUND

	value := c.PostForm("key_value")
	switch value {
	//switch boot order.
	// case "console.conf":
	// case "logfile.conf":
	// case "enum.conf":
	// case "xml_curl.conf":
	case "odbc_cdr.conf": //1th request.
		if conf, e := odbc_cdr.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "sofia.conf": //2th request(a request per profile).
		if conf, e := sofia.ReadConfiguration(c); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	//case "loopback.conf": //3th
	//case "verto.conf": //4th
	case "conference.conf": //5th
		if conf, e := autoload.ReadDefaultConfiguration("conference.conf.xml"); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "db.conf": //6th
		if conf, e := db.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "fifo.conf": //7th
		if conf, e := fifo.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "hash.conf": //8th
		if conf, e := autoload.ReadDefaultConfiguration("hash.conf.xml"); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "voicemail.conf": //9th
		if conf, e := voicemail.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	//case "httapi.conf": //10th
	//case "spandsp.conf": //11th
	case "opus.conf": //12th
		if conf, e := autoload.ReadDefaultConfiguration("opus.conf.xml"); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "avformat.conf": //13th
		if conf, e := autoload.ReadDefaultConfiguration("av.conf.xml"); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "avcodec.conf": //14th
		if conf, e := autoload.ReadDefaultConfiguration("av.conf.xml"); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "sndfile.conf": //15th
		if conf, e := autoload.ReadDefaultConfiguration("sndfile.conf.xml"); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "local_stream.conf": //16th
		if conf, e := local_stream.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	//case "lua.conf": //17th
	case "post_load_modules.conf": //18th
	case "acl.conf": //19th
		if conf, e := acl.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "event_socket.conf": //20th
		if conf, e := event_socket.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}
	case "post_load_switch.conf": //21th
	case "switch.conf": //22th
		if conf, e := switch_main.ReadConfiguration(); e != nil {
			body = NOT_FOUND
		} else {
			body = fmt.Sprintf(CONFIGURATION, conf)
		}

	//reloadxml
	case "timezones.conf":
		body = NOT_FOUND
	default:
		body = NOT_FOUND
	}
	return body
}
