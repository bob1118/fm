package sofia

import "github.com/gin-gonic/gin"

//freeswitch mod_sofia configuration.
//default configuration is file sofia.conf.xml

//1，request
//http://10.10.10.250/fsapi
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=sofia.conf&Event-Name=REQUEST_PARAMS&Core-UUID=1f1f888d-760b-486c-826d-e5b801fef0b8&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2014%3A31%3A04&Event-Date-GMT=Mon,%2029%20Mar%202021%2006%3A31%3A04%20GMT&Event-Date-Timestamp=1616999464247261&Event-Calling-File=sofia.c&Event-Calling-Function=config_sofia&Event-Calling-Line-Number=4491&Event-Sequence=27]

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
func ReadConfiguration(c *gin.Context) (b string) { return "" }

//write configuration into db.
func WriteConfiguration(s string) {}
