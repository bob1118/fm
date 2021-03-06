package fsapi

import (
	"fmt"

	"github.com/bob1118/fm/config/fmconfig"
	"github.com/bob1118/fm/models"
	"github.com/bob1118/fm/utils"
	"github.com/gin-gonic/gin"
)

// 1,request:
// 1.1 switch boot
// data: [hostname=bob-office&section=directory&tag_name=&key_name=&key_value=&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A29&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A29%20GMT&Event-Date-Timestamp=1618565369759906&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3097&Event-Sequence=37&purpose=gateways&profile=external]
// data: [hostname=bob-office&section=directory&tag_name=&key_name=&key_value=&Event-Name=REQUEST_PARAMS&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A31&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A31%20GMT&Event-Date-Timestamp=1618565371560583&Event-Calling-File=sofia.c&Event-Calling-Function=launch_sofia_worker_thread&Event-Calling-Line-Number=3097&Event-Sequence=43&purpose=gateways&profile=internal]
// data: [hostname=bob-office&section=directory&tag_name=domain&key_name=name&key_value=10.10.10.250&Event-Name=GENERAL&Core-UUID=c8eb6d34-b0f7-4d67-b70a-e6693d45cc01&FreeSWITCH-Hostname=bob-office&FreeSWITCH-Switchname=bob-office&FreeSWITCH-IPv4=10.10.10.250&FreeSWITCH-IPv6=%3A%3A1&Event-Date-Local=2021-04-16%2017%3A29%3A47&Event-Date-GMT=Fri,%2016%20Apr%202021%2009%3A29%3A47%20GMT&Event-Date-Timestamp=1618565387896133&Event-Calling-File=switch_core.c&Event-Calling-Function=switch_load_network_lists&Event-Calling-Line-Number=1623&Event-Sequence=478&domain=10.10.10.250&purpose=network-list]
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

//doDirectory function return xml direcotry.
func doDirectory(c *gin.Context) (b string) {
	body := NOT_FOUND

	//useragent reg, subscribe invite.
	if false ||
		(utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("action"), "sip_auth") && utils.IsEqual(c.PostForm("sip_auth_method"), "REGISTER")) ||
		(utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("action"), "sip_auth") && utils.IsEqual(c.PostForm("sip_auth_method"), "SUBSCRIBE")) ||
		(utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("action"), "sip_auth") && utils.IsEqual(c.PostForm("sip_auth_method"), "INVITE")) {
		ua := models.Account{}
		ua.Aid = c.PostForm("user")
		ua.Adomain = c.PostForm("domain")
		body = useragentAuth(ua)
	}

	//voicemail need lookup a user id, response like auth.
	if utils.IsEqual(c.PostForm("Event-Name"), "GENERAL") && utils.IsEqual(c.PostForm("action"), "message-count") {
		ua := models.Account{}
		ua.Aid = c.PostForm("user")
		ua.Adomain = c.PostForm("domain")
		body = useragentAuth(ua)
	}

	//an endpoint requests reverse authentication for a request, using reverse-auth-lookup;
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("action"), "reverse-auth-lookup") {
		ua := models.Account{}
		ua.Aid = c.PostForm("user")
		ua.Adomain = c.PostForm("domain")
		body = useragentReverseAuth(ua)
	}

	//multi tenant, sofia profile internal rescan/restart.
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "gateways") && utils.IsEqual(c.PostForm("profile"), "internal") {
		if models.GetAccountsCount("true") > 0 {
			if domains, err := models.DistinctAccountDomains(); err != nil {
				body = NOT_FOUND
			} else {
				var alldomain string
				for _, domain := range domains {
					alldomain += fmt.Sprintf(DOMAIN, domain)
				}
				body = fmt.Sprintf(DOMAINS, alldomain)
			}
		} else {
			body = NOT_FOUND
		}
	}

	// user's gateways ?? like conf/direcotry/default/brian.xml or conf/direcotry/default/example.com.xml
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "gateways") && utils.IsEqual(c.PostForm("profile"), "external") {
		body = NOT_FOUND
	}

	//switch_load_network_lists
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "network-list") {
		body = NOT_FOUND
	}

	//voicemail ?
	if utils.IsEqual(c.PostForm("Event-Name"), "REQUEST_PARAMS") && utils.IsEqual(c.PostForm("purpose"), "publish-vm") {
		body = NOT_FOUND
	}
	return body
}

func useragentAuth(u models.Account) (b string) {
	body := NOT_FOUND
	if isExist, ua := models.IsExistByiddomain(u); isExist {
		if !fmconfig.CFG.Runtime.Enablehash {
			body = fmt.Sprintf(USERAGENT, ua.Adomain, ua.Agroup, ua.Aid, ua.Acacheable, ua.Apassword)
		} else {
			body = fmt.Sprintf(USERAGENT_A1HASH, ua.Adomain, ua.Agroup, ua.Aid, ua.Acacheable, ua.Aa1hash)
		}
	}
	return body
}

func useragentReverseAuth(u models.Account) (b string) {
	body := NOT_FOUND
	if isExist, ua := models.IsExistByiddomain(u); isExist {
		body = fmt.Sprintf(USERAGENT_REVERSE, ua.Adomain, ua.Agroup, ua.Aid, ua.Acacheable, ua.Aid, ua.Apassword)
	}
	return body
}
