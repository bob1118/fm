package fsapi

//notfound
const NOT_FOUND string = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<document type="freeswitch/xml">
	<section name="result">
		<result status="not found"/>
	</section>
</document>`

//USERAGENT
const USERAGENT string = `<document type="freeswitch/xml" encoding="UTF-8">
<section name="directory">
 <domain name="%s">
  <groups>
  <group name="%s">
   <users>
   <user id="%s"  cacheable="%s">
	<params>
	 <param name="password" value="%s"/>
	 <param name="dial-string" value="{presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(${dialed_user}@${dialed_domain})}"/>
	</params>
	<variables>
	 <variable name="user_context" value="default"/>
	 <variable name="record_stereo" value="true"/>
	</variables>
   </user>
   </users>
  </group>
  </groups>
 </domain>
</section>
</document>`

//USERAGENT_A1HASH
const USERAGENT_A1HASH string = `<document type="freeswitch/xml" encoding="UTF-8">
<section name="directory">
 <domain name="%s">
  <groups>
  <group name="%s">
   <users>
   <user id="%s"  cacheable="%s">
	<params>
	 <param name="a1-hash" value="%s"/>
	 <param name="dial-string" value="{presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(${dialed_user}@${dialed_domain})}"/>
	</params>
	<variables>
	 <variable name="user_context" value="default"/>
	 <variable name="record_stereo" value="true"/>
	</variables>
   </user>
   </users>
  </group>
  </groups>
 </domain>
</section>
</document>`

//USERAGENT_REVERSE
const USERAGENT_REVERSE string = `<document type="freeswitch/xml" encoding="UTF-8">
<section name="directory">
 <domain name="%s">
  <groups>
  <group name="%s">
   <users>
   <user id="%s"  cacheable="%s">
	<params>
	 <param name="reverse-auth-user" value="%s"/>
	 <param name="reverse-auth-pass" value="%s"/>
	</params>
   </user>
   </users>
  </group>
  </groups>
 </domain>
</section>
</document>`

//domains
const DOMAINS string = `<document type="freeswitch/xml">
 <section name="directory">	
  %s
 </section>
</document>`

//domain
const DOMAIN string = `    <domain name="%s">
      <params>
        <param name="dial-string" value="{presence_id=${dialed_user}@${dialed_domain}}${sofia_contact(${dialed_user}@${dialed_domain})}"/>
      </params>
      <variables>
        <variable name="example_var" value="example_value_1"/>
      </variables>
      <user id="default" />
    </domain>
`

// CONFIGURATION
const CONFIGURATION string = `<document type="freeswitch/xml"  encoding="UTF-8">
<section name="configuration">
%s
</section>
</document>`

//dialplan inbound.
const DialplanInbound string = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<document type="freeswitch/xml">
  <section name="dialplan" description="dialplan inbound for FreeSwitch">
    <context name="default">
      <extension name="inbound">
        <condition>
          <action application="set" data="continue_on_fail=true"/>
          <action application="park"/>
        </condition>
      </extension>
    </context>
  </section>
</document>
`

//dialplan outbound.
const DialplanOutbound string = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<document type="freeswitch/xml">
  <section name="dialplan" description="dialplan outbound FreeSwitch">
    <context name="default">
      <extension name="outbound">
        <condition>
          <action application="socket" data="%s async full"/>
        </condition>
      </extension>
    </context>
  </section>
</document>
`
