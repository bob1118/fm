package sofia

import (
	"errors"

	"github.com/gin-gonic/gin"
	"githug.com/bob118/fm/models"
	"githug.com/bob118/fm/utils"
)

//freeswitch mod_sofia configuration.
//default configuration is file sofia.conf.xml

//1，request
//http://10.10.10.250/fsapi
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A26&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A26%20GMT&Event-Date-Timestamp=1618565366035759&Event-Calling-File=sofia.c&Event-Calling-Function=config_sofia&Event-Calling-Line-Number=4491&Event-Sequence=27]
//
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A27&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A27%20GMT&Event-Date-Timestamp=1618565367391604&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=29&profile=external-ipv6]
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A28&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A28%20GMT&Event-Date-Timestamp=1618565368557666&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=32&profile=external]
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A29&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A29%20GMT&Event-Date-Timestamp=1618565369574163&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=36&profile=internal-ipv6]
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A30&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A30%20GMT&Event-Date-Timestamp=1618565370358404&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3079&Event-Sequence=40&profile=internal]

//2，response
// <document type="freeswitch/xml">
//   <section name="configuration">
////     <configuration name="sofia.conf" description="sofia endpoint">

// <global_settings>
// <param name="log-level" value="0"/>
// <!-- <param name="abort-on-empty-external-ip" value="true"/> -->
// <!-- <param name="auto-restart" value="false"/> -->
// <param name="debug-presence" value="0"/>
// <!-- <param name="capture-server" value="udp:homer.domain.com:5060"/> -->

// <!--
// 	the new format for HEPv2/v3 and capture ID

// protocol:host:port;hep=2;capture_id=200;

// -->

// <!-- <param name="capture-server" value="udp:homer.domain.com:5060;hep=3;capture_id=100"/> -->
// </global_settings>

// <!--
//   The rabbit hole goes deep.  This includes all the
//   profiles in the sip_profiles directory that is up
//   one level from this directory.
// -->
// <profiles>
// <X-PRE-PROCESS cmd="include" data="../sip_profiles/*.xml"/>
// </profiles>

////     </configuration>
//   </section>
// </document>

//read configuration from file, and then write into db.
func MakeDefaultConfiguration() {}

//read configuration from db.
func ReadConfiguration(c *gin.Context) (e error, b string) {
	var err error
	var body string

	profile := c.PostForm("profile")
	switch profile {
	case "internal":
		if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("reconfig"), "true") {
		}
	case "external":
		if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") { //&& utils.IsEqual(c.PostForm("reconfig"), "true") {
			if models.GetGatewaysCount(true) == 0 {
				err = errors.New("sofia profile external gateway count return 0")
			} else {
				// var allgateway string
				// gws := models.GetGateways(true)
				// for _, gw := range gws {
				// 	allgateway += fmt.Sprintf(GATEWAY,
				// 		gw.Gname,
				// 		gw.Gusername,
				// 		gw.Grealm,
				// 		gw.Gfromuser,
				// 		gw.Gfromdomain,
				// 		gw.Gpassword,
				// 		gw.Gextension,
				// 		gw.Gproxy,
				// 		gw.Gregisterproxy,
				// 		gw.Gexpire,
				// 		gw.Gregister,
				// 		gw.Gcalleridinfrom,
				// 		gw.Gextensionincontact,
				// 		gw.Goptionping)
				// }
				// body = fmt.Sprintf(GATEWAYS, allgateway)
			}
		}
	}
	return err, body
}

//write configuration into db.
func WriteConfiguration(s string) {}
