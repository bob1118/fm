package xmlbuilder

import (
	"bytes"
	"errors"
	"log"
	"os"
	"runtime"
	"strings"

	"github.com/bob1118/fm/config/fmconfig"
)

var defaultDirectory string

func init() {
	SetDefaultDirectory(fmconfig.CFG.Runtime.ConfDirectory)
}

func SetDefaultDirectory(dir string) {
	if _, e := os.Stat(dir); e == nil {
		if strings.Contains(dir, `\`) {
			dir = strings.ReplaceAll(dir, `\`, `/`)
		}
		if strings.HasSuffix(dir, `/`) {
			defaultDirectory = dir
		} else {
			defaultDirectory = dir + `/`
		}
	} else {
		sysType := runtime.GOOS
		switch sysType {
		case "linux":
			defaultDirectory = "/etc/freeswitch/"
		case "windows":
			defaultDirectory = "C:/Program Files/FreeSWITCH/conf/"
		}
	}
}

func GetDefaultDirectory() (dir string) {
	return defaultDirectory
}

//BuildPersonalConf build personal *.conf.xml files, set origination *.conf.xml named *.conf.xml.default
func BuildPersonalConf() {

	//./*.xml
	makePersonalXml("freeswitch.xml")
	makePersonalXml("vars.xml")

	//./autoload_configs/*.xml
	makePersonalXml("switch.conf.xml")
	makePersonalXml("modules.conf.xml")
	makePersonalXml("xml_curl.conf.xml")

	//./sip_profiles/*.xml
	makePersonalXml("internal.xml")
	makePersonalXml("internal-ipv6.xml")
	makePersonalXml("external.xml")
	makePersonalXml("external-ipv6.xml")
}

func BuildDefaultConf() {}

//makePersonalXml file.
func makePersonalXml(name string) (e error) {
	var err error

	switch name {
	case "freeswitch.xml":
	case "vars.xml": //freeswitch global var define.
		var newvars []byte
		filepath := defaultDirectory + "vars.xml"
		defaultfilepath := defaultDirectory + "vars.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if vars, e := os.ReadFile(filepath); e != nil {
				if os.IsNotExist(e) {
					log.Println("vars.xml is missing ...")
				}
				err = e
			} else {
				os.WriteFile(defaultfilepath, vars, 0660)
				var strmultiline string
				if strings.EqualFold(runtime.GOOS, "windows") {
					strmultiline = `  <X-PRE-PROCESS cmd="set" data="default_password=D_e_f_a_u_l_t_P_a_s_s_w_o_r_d"/>` + "\r\n" +
						`  <X-PRE-PROCESS cmd="set" data="local_ip_v4=10.10.10.250"/>` + "\r\n" +
						`  <X-PRE-PROCESS cmd="set" data="json_db_handle=$${pg_handle}"/>` + "\r\n" +
						`  <X-PRE-PROCESS cmd="set" data="pg_handle=pgsql://hostaddr=127.0.0.1 dbname=freeswitch user=fsdba password=fsdba"/>`
				} else {
					strmultiline = `  <X-PRE-PROCESS cmd="set" data="default_password=D_e_f_a_u_l_t_P_a_s_s_w_o_r_d"/>` + "\n" +
						`  <X-PRE-PROCESS cmd="set" data="local_ip_v4=10.10.10.250"/>` + "\n" +
						`  <X-PRE-PROCESS cmd="set" data="json_db_handle=$${pg_handle}"/>` + "\n" +
						`  <X-PRE-PROCESS cmd="set" data="pg_handle=pgsql://hostaddr=127.0.0.1 dbname=freeswitch user=fsdba password=fsdba"/>`
				}
				// <X-PRE-PROCESS cmd="set" data="default_password=1234"/>
				newvars = Update(vars, `  <X-PRE-PROCESS cmd="set" data="default_password=1234"/>`, strmultiline)
				//  <X-PRE-PROCESS cmd="stun-set" data="external_sip_ip=stun:stun.freeswitch.org"/>
				newvars = Update(newvars, `  <X-PRE-PROCESS cmd="stun-set" data="external_sip_ip=stun:stun.freeswitch.org"/>`,
					`  <X-PRE-PROCESS cmd="stun-set" data="external_sip_ip=$${local_ip_v4}"/>`)
				//  <X-PRE-PROCESS cmd="stun-set" data="external_rtp_ip=stun:stun.freeswitch.org"/>
				newvars = Update(newvars, `  <X-PRE-PROCESS cmd="stun-set" data="external_rtp_ip=stun:stun.freeswitch.org"/>`,
					`  <X-PRE-PROCESS cmd="stun-set" data="external_rtp_ip=$${local_ip_v4}"/>`)
				os.WriteFile(filepath, newvars, 0660)
			}
		}

	case "switch.conf.xml": //freeswitch autoload switch.conf.xml
		var newswitch []byte
		filepath := defaultDirectory + "autoload_configs/switch.conf.xml"
		defaultfilepath := defaultDirectory + "autoload_configs/switch.conf.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if switch_main, e := os.ReadFile(filepath); e != nil {
				err = e
			} else {
				os.WriteFile(defaultfilepath, switch_main, 0660)
				newswitch = Update(switch_main,
					`<!-- <param name="core-db-dsn" value="dsn:username:password" /> -->`,
					`<param name="core-db-dsn" value="$${pg_handle}"/>`)
				os.WriteFile(filepath, newswitch, 0600)
			}
		}
	case "modules.conf.xml": //freeswitch autoload modules define.
		var newmodules []byte
		filepath := defaultDirectory + "autoload_configs/modules.conf.xml"
		defaultfilepath := defaultDirectory + "autoload_configs/modules.conf.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if modules, e := os.ReadFile(filepath); e != nil {
				err = e
			} else {
				os.WriteFile(defaultfilepath, modules, 0660)
				//    <load module="mod_enum"/>
				newmodules = Update(modules, `    <load module="mod_enum"/>`, `    <!--<load module="mod_enum"/>-->`)
				//    <!-- <load module="mod_xml_curl"/> -->
				newmodules = Update(newmodules, `    <!-- <load module="mod_xml_curl"/> -->`, `    <load module="mod_xml_curl"/>`)
				//    <load module="mod_cdr_csv"/>
				newmodules = Update(newmodules, `    <load module="mod_cdr_csv"/>`, `    <load module="mod_odbc_cdr"/>`)
				//    <load module="mod_loopback"/>
				newmodules = Update(newmodules, `    <load module="mod_loopback"/>`, `    <!--<load module="mod_loopback"/>-->`)
				//    <load module="mod_rtc"/>
				newmodules = Update(newmodules, `    <load module="mod_rtc"/>`, `    <!--<load module="mod_rtc"/>-->`)
				//    <load module="mod_verto"/>
				newmodules = Update(newmodules, `    <load module="mod_verto"/>`, `    <!--<load module="mod_verto"/>-->`)
				//    <load module="mod_signalwire"/>
				newmodules = Update(newmodules, `    <load module="mod_signalwire"/>`, `    <!--<load module="mod_signalwire"/>-->`)
				//    <load module="mod_expr"/>
				newmodules = Update(newmodules, `    <load module="mod_expr"/>`, `    <!--<load module="mod_expr"/>-->`)
				//    <load module="mod_valet_parking"/>
				newmodules = Update(newmodules, `    <load module="mod_valet_parking"/>`, `    <!--<load module="mod_valet_parking"/>-->`)
				//    <load module="mod_httapi"/>
				newmodules = Update(newmodules, `    <load module="mod_httapi"/>`, `    <!--<load module="mod_httapi"/>-->`)
				//    <load module="mod_dialplan_asterisk"/>
				newmodules = Update(newmodules, `    <load module="mod_dialplan_asterisk"/>`, `    <!--<load module="mod_dialplan_asterisk"/>-->`)
				//    <load module="mod_spandsp"/>
				newmodules = Update(newmodules, `    <load module="mod_spandsp"/>`, `    <!--<load module="mod_spandsp"/>-->`)
				//    <load module="mod_b64"/>
				newmodules = Update(newmodules, `    <load module="mod_b64"/>`, `    <!--<load module="mod_b64"/>-->`)
				//    <load module="mod_lua"/>
				newmodules = Update(newmodules, `    <load module="mod_lua"/>`, `    <!--<load module="mod_lua"/>-->`)
				//    <load module="mod_say_en"/>
				newmodules = Update(newmodules, `    <load module="mod_say_en"/>`, `    <!--<load module="mod_say_en"/>-->`)
				os.WriteFile(filepath, newmodules, 0660)
			}
		}
	case "xml_curl.conf.xml": //freeswitch dynamically control interface.
		var newxmlcurl []byte
		filepath := defaultDirectory + "autoload_configs/xml_curl.conf.xml"
		defaultfilepath := defaultDirectory + "autoload_configs/xml_curl.conf.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if xmlcurl, e := os.ReadFile(filepath); e != nil {
				err = e
			} else {
				os.WriteFile(defaultfilepath, xmlcurl, 0660)
				//<!-- <param name="gateway-url" value="http://www.freeswitch.org/gateway.xml" bindings="dialplan"/> -->
				newxmlcurl = Update(xmlcurl,
					`<!-- <param name="gateway-url" value="http://www.freeswitch.org/gateway.xml" bindings="dialplan"/> -->`,
					`<param name="gateway-url" value="http://$${local_ip_v4}/fsapi" bindings="dialplan|configuration|directory|phrases"/>`)
				os.WriteFile(filepath, newxmlcurl, 0660)
			}
		}

	case "internal.xml": //sip_profiles/internal.xml
		var internal []byte
		filepath := defaultDirectory + "sip_profiles/internal.xml"
		defaultfilepath := defaultDirectory + "sip_profiles/internal.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if data, e := os.ReadFile(filepath); e != nil {
				err = e
			} else {
				os.WriteFile(defaultfilepath, data, 0660)
				//<!--<param name="odbc-dsn" value="pgsql://hostaddr=127.0.0.1 dbname=freeswitch user=freeswitch password='' options='-c client_min_messages=NOTICE' application_name='freeswitch'" />-->
				internal = Update(data,
					`<!--<param name="odbc-dsn" value="pgsql://hostaddr=127.0.0.1 dbname=freeswitch user=freeswitch password='' options='-c client_min_messages=NOTICE' application_name='freeswitch'" />-->`,
					`<param name="odbc-dsn" value="$${pg_handle}"/>`)
				internal = Update(internal,
					`<param name="force-register-domain" value="$${domain}"/>`,
					`<!--<param name="force-register-domain" value="$${domain}"/>-->`)
				internal = Update(internal,
					`<param name="force-subscription-domain" value="$${domain}"/>`,
					`<!--<param name="force-subscription-domain" value="$${domain}"/>-->`)
				internal = Update(internal,
					`<param name="force-register-db-domain" value="$${domain}"/>`,
					`<!--<param name="force-register-db-domain" value="$${domain}"/>-->`)
				os.WriteFile(filepath, internal, 0600)
			}
		}
	case "internal-ipv6.xml": //sip_profiles/internal-ipv6.xml
		var internalipv6 []byte
		filepath := defaultDirectory + "sip_profiles/internal-ipv6.xml"
		defaultfilepath := defaultDirectory + "sip_profiles/internal-ipv6.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if data, e := os.ReadFile(filepath); e != nil {
				err = e
			} else {
				os.WriteFile(defaultfilepath, data, 0660)
				//<!--<param name="odbc-dsn" value="dsn:user:pass"/>-->
				internalipv6 = Update(data,
					`<!--<param name="odbc-dsn" value="dsn:user:pass"/>-->`,
					`<param name="odbc-dsn" value="$${pg_handle}"/>`)
				internalipv6 = Update(internalipv6,
					`<param name="force-register-domain" value="$${domain}"/>`,
					`<!--<param name="force-register-domain" value="$${domain}"/>-->`)
				internalipv6 = Update(internalipv6,
					`<param name="force-subscription-domain" value="$${domain}"/>`,
					`<!--<param name="force-subscription-domain" value="$${domain}"/>-->`)
				internalipv6 = Update(internalipv6,
					`<param name="force-register-db-domain" value="$${domain}"/>`,
					`<!--<param name="force-register-db-domain" value="$${domain}"/>-->`)
				os.WriteFile(filepath, internalipv6, 0600)
			}
		}
	case "external.xml": //sip_profiles/external.xml
		var newdata []byte
		filepath := defaultDirectory + "sip_profiles/external.xml"
		defaultfilepath := defaultDirectory + "sip_profiles/external.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if data, e := os.ReadFile(filepath); e != nil {
				err = e
			} else {
				os.WriteFile(defaultfilepath, data, 0660)
				//<!-- ************************************************* -->
				newdata = Update(data,
					`<!-- ************************************************* -->`,
					`<param name="odbc-dsn" value="$${pg_handle}"/>`)
				os.WriteFile(filepath, newdata, 0600)
			}
		}
	case "external-ipv6.xml": //sip_profiles/external-ipv6.xml
		var newdata []byte
		filepath := defaultDirectory + "sip_profiles/external-ipv6.xml"
		defaultfilepath := defaultDirectory + "sip_profiles/external-ipv6.xml.default"
		if _, e := os.Stat(defaultfilepath); os.IsNotExist(e) {
			if data, e := os.ReadFile(filepath); e != nil {
				err = e
			} else {
				os.WriteFile(defaultfilepath, data, 0660)
				//<!-- ************************************************* -->
				newdata = Update(data,
					`<!-- ************************************************* -->`,
					`<param name="odbc-dsn" value="$${pg_handle}"/>`)
				os.WriteFile(filepath, newdata, 0600)
			}
		}

	default:
		err = errors.New(`unsupport xml name`)
	}
	return err
}

func Update(b []byte, p string, v string) (s []byte) {
	return bytes.ReplaceAll(b, []byte(p), []byte(v))
}
