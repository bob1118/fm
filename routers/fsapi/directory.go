package fsapi

import (
	"github.com/gin-gonic/gin"
	"githug.com/bob118/fm/utils"
)

// 1,request:
// 1.1 switch boot
// http://10.10.10.250/fsapi
// data: [hostname=bob-office&section=directory&tag_name=&key_name=&key_value=&Event-Name=REQUEST_PARAMS&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2016%3A25%3A59&Event-Date-GMT=Mon,%2029%20Mar%202021%2008%3A25%3A59%20GMT&Event-Date-Timestamp=1617006359600673&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3097&Event-Sequence=44&purpose=gateways&profile=internal]
// http://10.10.10.250/fsapi
// data: [hostname=bob-office&section=directory&tag_name=&key_name=&key_value=&Event-Name=REQUEST_PARAMS&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2016%3A25%3A57&Event-Date-GMT=Mon,%2029%20Mar%202021%2008%3A25%3A57%20GMT&Event-Date-Timestamp=1617006357602928&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3097&Event-Sequence=40&purpose=gateways&profile=external]
// http://10.10.10.250/fsapi
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=GENERAL&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2016%3A26%3A26&Event-Date-GMT=Mon,%2029%20Mar%202021%2008%3A26%3A26%20GMT&Event-Date-Timestamp=1617006386830260&Event-Calling-File=switch_core.c&Event-Calling-Function=switch_load_network_lists&Event-Calling-Line-Number=1623&Event-Sequence=480&domain=10.10.10.250&purpose=network-list]
// 1.2 ua reg
// REGISTER
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=REQUEST_PARAMS&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2017%3A04%3A02&Event-Date-GMT=Mon,%2029%20Mar%202021%2009%3A04%3A02%20GMT&Event-Date-Timestamp=1617008642271157&Event-Calling-File=sofia_reg.c&Event-Calling-Function=sofia_reg_parse_auth&Event-Calling-Line-Number=2846&Event-Sequence=766&action=sip_auth&sip_profile=internal&sip_user_agent=eyeBeam%20AudioOnly%20release%203015c%20stamp%2027106&sip_auth_username=1000&sip_auth_realm=10.10.10.250&sip_auth_nonce=a212e670-2d31-441b-9a48-04b1d1091131&sip_auth_uri=sip%3A10.10.10.250&sip_contact_user=1000&sip_contact_host=10.10.10.250&sip_to_user=1000&sip_to_host=10.10.10.250&sip_via_protocol=udp&sip_from_user=1000&sip_from_host=10.10.10.250&sip_call_id=fb29f5460346c530%40Ym9iLW9mZmljZQ..&sip_request_host=10.10.10.250&sip_auth_qop=auth&sip_auth_cnonce=39645b121d34ea15&sip_auth_nc=00000001&sip_auth_response=14048e801caa7eead5ca3d62ad911c7d&sip_auth_method=REGISTER&client_port=10554&key=id&user=1000&domain=10.10.10.250&ip=10.10.10.250]
// message-count
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=GENERAL&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2017%3A04%3A04&Event-Date-GMT=Mon,%2029%20Mar%202021%2009%3A04%3A04%20GMT&Event-Date-Timestamp=1617008644541207&Event-Calling-File=mod_voicemail.c&Event-Calling-Function=resolve_id&Event-Calling-Line-Number=1363&Event-Sequence=770&action=message-count&key=id&user=1000&domain=10.10.10.250]
// SUBSCRIBE
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=REQUEST_PARAMS&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2017%3A04%3A04&Event-Date-GMT=Mon,%2029%20Mar%202021%2009%3A04%3A04%20GMT&Event-Date-Timestamp=1617008644670862&Event-Calling-File=sofia_reg.c&Event-Calling-Function=sofia_reg_parse_auth&Event-Calling-Line-Number=2846&Event-Sequence=772&action=sip_auth&sip_profile=internal&sip_user_agent=eyeBeam%20AudioOnly%20release%203015c%20stamp%2027106&sip_auth_username=1000&sip_auth_realm=10.10.10.250&sip_auth_nonce=d40c5337-b679-4378-8087-50d95f49bee4&sip_auth_uri=sip%3A1000%4010.10.10.250&sip_contact_user=1000&sip_contact_host=10.10.10.250&sip_to_user=1000&sip_to_host=10.10.10.250&sip_via_protocol=udp&sip_from_user=1000&sip_from_host=10.10.10.250&sip_call_id=2023c7471369a769%40Ym9iLW9mZmljZQ..&sip_request_user=1000&sip_request_host=10.10.10.250&sip_auth_qop=auth&sip_auth_cnonce=16c31bc30f2672a7&sip_auth_nc=00000001&sip_auth_response=f71168e3a742860ef28ce2d5e90ae540&sip_auth_method=SUBSCRIBE&client_port=10554&key=id&user=1000&domain=10.10.10.250&ip=10.10.10.250]
// message-count
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=GENERAL&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2017%3A04%3A06&Event-Date-GMT=Mon,%2029%20Mar%202021%2009%3A04%3A06%20GMT&Event-Date-Timestamp=1617008646961133&Event-Calling-File=mod_voicemail.c&Event-Calling-Function=resolve_id&Event-Calling-Line-Number=1363&Event-Sequence=775&action=message-count&key=id&user=1000&domain=10.10.10.250]
// 1.3 ua unreg
// REGISTER
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=REQUEST_PARAMS&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2016%3A56%3A52&Event-Date-GMT=Mon,%2029%20Mar%202021%2008%3A56%3A52%20GMT&Event-Date-Timestamp=1617008212241740&Event-Calling-File=sofia_reg.c&Event-Calling-Function=sofia_reg_parse_auth&Event-Calling-Line-Number=2846&Event-Sequence=710&action=sip_auth&sip_profile=internal&sip_user_agent=eyeBeam%20AudioOnly%20release%203015c%20stamp%2027106&sip_auth_username=1000&sip_auth_realm=10.10.10.250&sip_auth_nonce=87288fa8-1eaa-49b5-ac36-46c0d2c04aaf&sip_auth_uri=sip%3A10.10.10.250&sip_contact_user=1000&sip_contact_host=10.10.10.250&sip_to_user=1000&sip_to_host=10.10.10.250&sip_via_protocol=udp&sip_from_user=1000&sip_from_host=10.10.10.250&sip_call_id=a06d155c9a204955%40Ym9iLW9mZmljZQ..&sip_request_host=10.10.10.250&sip_auth_qop=auth&sip_auth_cnonce=2132fe4c7b346244&sip_auth_nc=00000002&sip_auth_response=83b08c69a3edbc9d12eb39bb5e860e59&sip_auth_method=REGISTER&client_port=10554&key=id&user=1000&domain=10.10.10.250&ip=10.10.10.250]
// 1.4 ua invite
// INVITE
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=REQUEST_PARAMS&Core-UUID=3369c8b1-2336-4435-a13c-5516a745ed75&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=2001%3A0%3A2851%3Ab9f0%3Ac5a%3Ac6e1%3Afeaa%3A107d&Event-Date-Local=2021-03-29%2017%3A10%3A26&Event-Date-GMT=Mon,%2029%20Mar%202021%2009%3A10%3A26%20GMT&Event-Date-Timestamp=1617009026320999&Event-Calling-File=sofia_reg.c&Event-Calling-Function=sofia_reg_parse_auth&Event-Calling-Line-Number=2846&Event-Sequence=825&action=sip_auth&sip_profile=internal&sip_user_agent=eyeBeam%20AudioOnly%20release%203015c%20stamp%2027106&sip_auth_username=1000&sip_auth_realm=10.10.10.250&sip_auth_nonce=bafb3087-7896-4b82-bb87-5b84fe92759a&sip_auth_uri=sip%3A9664%4010.10.10.250&sip_contact_user=1000&sip_contact_host=10.10.10.250&sip_to_user=9664&sip_to_host=10.10.10.250&sip_via_protocol=udp&sip_from_user=1000&sip_from_host=10.10.10.250&sip_call_id=ef4eea5592213660%40Ym9iLW9mZmljZQ..&sip_request_user=9664&sip_request_host=10.10.10.250&sip_auth_qop=auth&sip_auth_cnonce=71da7875fcc4a6c3&sip_auth_nc=00000001&sip_auth_response=3a3f199a4493e6b1a3b336d8ed71b6f4&sip_auth_method=INVITE&client_port=10554&key=id&user=1000&domain=10.10.10.250&ip=10.10.10.250]

//doDirectory read useragents@domains from db.
func doDirectory(c *gin.Context) (b string) {
	body := Notfound

	//useragent reg, subscribe invite.
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("action"), "sip_auth") {
		method := c.PostForm("sip_auth_method")
		switch method {
		case "REGISTER", "SUBSCRIBE", "INVITE":
			//do auth
			useragentAuth()
		}
	}

	//reverse auth
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("action"), "reverse-auth-lookup") {
		useragentAuth()
	}

	//message-count.
	if utils.IsEqual(c.PostForm("Event-Name"), "GENERAL") && utils.IsEqual(c.PostForm("action"), "message-count") {
	}

	//multi tenant, sofia profile internal rescan/restart.
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "gateways") && utils.IsEqual(c.PostForm("profile"), "internal") {
		//
	}
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "gateways") && utils.IsEqual(c.PostForm("profile"), "external") {
	}

	//
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "network-list") {
	}
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "publish-vm") {
	}
	return body
}

func useragentAuth() {}
