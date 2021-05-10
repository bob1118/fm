package sofia

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	"github.com/bob1118/fm/models"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
	"github.com/bob1118/fm/utils"
	"github.com/gin-gonic/gin"
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

var defaultConfname, defaultConffile, defaultData string

func init() { defaultConfname = "sofia.conf.xml" }

//MakeDefaultConfiguration.
func MakeDefaultConfiguration() {}

//ReadConfiguration from file.
func ReadConfiguration(c *gin.Context) (b string, e error) {
	var err error

	event_calling_function := c.PostForm("Event-Calling-Function")
	switch event_calling_function {
	case "config_sofia":
		profile := c.PostForm("profile")
		reconfig := c.PostForm("reconfig")
		switch profile {
		case "": //call config_sofia, while freeswitch load mod_sofia.
			defaultConffile = xmlbuilder.GetDefaultDirectory() + `autoload_configs/` + defaultConfname
			if data, e := os.ReadFile(defaultConffile); e != nil {
				err = e
			} else {
				data = bytes.ReplaceAll(data,
					[]byte(`<X-PRE-PROCESS cmd="include" data="../sip_profiles/*.xml"/>`),
					[]byte(`<X-PRE-PROCESS cmd="include" data="./sip_profiles/*.xml"/>`))
				defaultData = string(data)
			}
		case "internal":
			if utils.IsEqual(reconfig, "true") { //profile internal reconfig
				defaultConffile = xmlbuilder.GetDefaultDirectory() + `sip_profiles/internal.xml`
				if internal, e := os.ReadFile(defaultConffile); e != nil {
					err = e
				} else { //internal settings rewrite.
					internal = bytes.ReplaceAll(internal,
						[]byte(`<param name="__param_name" value="value"/>`),
						[]byte(`<!--<param name="__param_name" value="value"/>-->`))
					defaultData = fmt.Sprintf(PROFILE, string(internal))
				}
			}
		case "internal-ipv6":
			if utils.IsEqual(reconfig, "true") { //profile internal-ipv6 reconfig
				defaultConffile = xmlbuilder.GetDefaultDirectory() + "sip_profiles/internal-ipv6.xml"
				if internalipv6, e := os.ReadFile(defaultConffile); e != nil {
					err = e
				} else { //internal-ipv6 settings rewrite.
					internalipv6 = bytes.ReplaceAll(internalipv6,
						[]byte(`<param name="__param_name" value="value"/>`),
						[]byte(`<!--<param name="__param_name" value="value"/>-->`))
					defaultData = fmt.Sprintf(PROFILE, string(internalipv6))
				}
			}
		case "external":
			if utils.IsEqual(reconfig, "true") { //profile external reconfig
				defaultConffile = xmlbuilder.GetDefaultDirectory() + "sip_profiles/external.xml"
				if external, e := os.ReadFile(defaultConffile); e != nil {
					err = e
				} else { //external settings rewrite.
					if xmlGateways, count := buildXmlGateways(); count > 0 { //external <gateways>xmlGateways</gateways>
						external = bytes.ReplaceAll(external, []byte(`    <X-PRE-PROCESS cmd="include" data="external/*.xml"/>`), []byte(xmlGateways))
					} else {
						external = bytes.ReplaceAll(external,
							[]byte(`<X-PRE-PROCESS cmd="include" data="external/*.xml"/>`),
							[]byte(`<X-PRE-PROCESS cmd="include" data="./sip_profiles/external/*.xml"/>`))
					}
					defaultData = fmt.Sprintf(PROFILE, string(external))
				}
			}
		case "external-ipv6":
			if utils.IsEqual(reconfig, "true") { //profile external-ipv6 reconfig
				defaultConffile = xmlbuilder.GetDefaultDirectory() + "sip_profiles/external-ipv6.xml"
				if externalipv6, e := os.ReadFile(defaultConffile); e != nil {
					err = e
				} else { //external-ipv6 settings rewrite.
					externalipv6 = bytes.ReplaceAll(externalipv6,
						[]byte(`<X-PRE-PROCESS cmd="include" data="external-ipv6/*.xml"/>`),
						[]byte(`<X-PRE-PROCESS cmd="include" data="./sip_profiles/external-ipv6/*.xml"/>`))
					defaultData = fmt.Sprintf(PROFILE, string(externalipv6))
				}
			}
		default:
			err = errors.New("Event-Calling-Function:config_sofia unknow profile ")
		}
	case "launch_sofia_worker_thread":
		profile := c.PostForm("profile")
		switch profile {
		case "internal": //./sip_profiles/internal.xml
			defaultConffile = xmlbuilder.GetDefaultDirectory() + `sip_profiles/internal.xml`
			if internal, e := os.ReadFile(defaultConffile); e != nil {
				err = e
			} else { //internal settings rewrite.
				internal = bytes.ReplaceAll(internal,
					[]byte(`<param name="__param_name" value="value"/>`),
					[]byte(`<!--<param name="__param_name" value="value"/>-->`))
				defaultData = fmt.Sprintf(PROFILE, string(internal))
			}
		case "internal-ipv6": //./sip_profiles/internal-ipv6.xml
			defaultConffile = xmlbuilder.GetDefaultDirectory() + "sip_profiles/internal-ipv6.xml"
			if internalipv6, e := os.ReadFile(defaultConffile); e != nil {
				err = e
			} else { //internal-ipv6 settings rewrite.
				internalipv6 = bytes.ReplaceAll(internalipv6,
					[]byte(`<param name="__param_name" value="value"/>`),
					[]byte(`<!--<param name="__param_name" value="value"/>-->`))
				defaultData = fmt.Sprintf(PROFILE, string(internalipv6))
			}
		case "external": //./sip_profiles/external.xml
			defaultConffile = xmlbuilder.GetDefaultDirectory() + "sip_profiles/external.xml"
			if external, e := os.ReadFile(defaultConffile); e != nil {
				err = e
			} else { //external settings rewrite.
				if xmlGateways, count := buildXmlGateways(); count > 0 { //external <gateways>xmlGateways</gateways>
					external = bytes.ReplaceAll(external, []byte(`    <X-PRE-PROCESS cmd="include" data="external/*.xml"/>`), []byte(xmlGateways))
				} else {
					external = bytes.ReplaceAll(external,
						[]byte(`<X-PRE-PROCESS cmd="include" data="external/*.xml"/>`),
						[]byte(`<X-PRE-PROCESS cmd="include" data="./sip_profiles/external/*.xml"/>`))
				}
				defaultData = fmt.Sprintf(PROFILE, string(external))
			}
		case "external-ipv6": //./sip_profiles/external-ipv6.xml
			defaultConffile = xmlbuilder.GetDefaultDirectory() + "sip_profiles/external-ipv6.xml"
			if externalipv6, e := os.ReadFile(defaultConffile); e != nil {
				err = e
			} else { //external-ipv6 settings rewrite.
				externalipv6 = bytes.ReplaceAll(externalipv6,
					[]byte(`<X-PRE-PROCESS cmd="include" data="external-ipv6/*.xml"/>`),
					[]byte(`<X-PRE-PROCESS cmd="include" data="./sip_profiles/external-ipv6/*.xml"/>`))
				defaultData = fmt.Sprintf(PROFILE, string(externalipv6))
			}
		default:
			err = errors.New(`sofia profle unknown`)
		}
	default:
		err = errors.New("Event-Calling-Function unknows")
	}

	return defaultData, err
}

//write configuration into db.
func WriteConfiguration(s string) {}

//buildXmlGateways
func buildXmlGateways() (s string, c uint) {
	var allgateway string
	count := models.GetGatewaysCount("true")
	if count > 0 {
		mygateways := models.GetGateways("true")
		for _, gw := range mygateways {
			allgateway += fmt.Sprintf(GATEWAY, gw.Gname, gw.Gusername, gw.Grealm, gw.Gfromuser, gw.Gfromdomain, gw.Gpassword, gw.Gextension, gw.Gproxy, gw.Gregisterproxy, gw.Gexpire, gw.Gregister, gw.Gcalleridinfrom, gw.Gextensionincontact, gw.Goptionping)
		}
	}
	return allgateway, count
}
