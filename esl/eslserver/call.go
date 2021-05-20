package eslserver

import (
	"fmt"

	"github.com/bob1118/fm/esl/run_time"
	"github.com/bob1118/fm/models"
)

type CALL struct {
	//CALL BASIC INFO
	coreuuid          string //Core-Uuid
	fsipv4            string //Freeswitch-Ipv4
	eventname         string //Event-Name
	uuid              string //Variable_uuid
	callid            string //Variable_call_uuid
	direction         string //Variable_direction
	profile           string //Variable_sofia_profile_name
	domain            string //Variable_domain_name
	gateway           string //Variable_sip_gateway
	ani               string //Caller-Ani
	distinationnumber string //Caller-Destination-Number
}

func init() {}

func (c *CALL) CallerIsUa() bool {
	mkey := c.ani + "@" + c.domain
	return run_time.IsUa(mkey)
}
func (c *CALL) CalleeIsUa() bool {
	mkey := c.distinationnumber + "@" + c.domain
	return run_time.IsUa(mkey)
}

func (c *CALL) CallFilterPassed() bool {
	passed := false
	if c.e164CalleeExist() {
		if !c.blacklistCallerExist() {
			passed = true
		}
	}
	return passed
}

func (c *CALL) e164CalleeExist() bool {
	is := false
	if len(c.distinationnumber) > 0 && len(c.gateway) > 0 {
		query := fmt.Sprintf(`gateway_name='%s' and e164_number='%s'`, c.gateway, c.distinationnumber)
		if e164s, err := models.GetE164s(query); err != nil {
			is = false
		} else {
			if e164s[0].Eenable {
				if !e164s[0].Elockin {
					is = true
				}
			}
		}
	}
	return is
}

func (c *CALL) blacklistCallerExist() bool {
	is := false
	if len(c.ani) > 0 {
		if _, exist := models.IsExistBlacklistCaller(c.ani); exist {
			is = true
		}
	}
	return is
}
