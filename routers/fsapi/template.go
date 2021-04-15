package fsapi

const Notfound string = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<document type="freeswitch/xml">
	<section name="result">
		<result status="not found"/>
	</section>
</document>`

//useragent
const Useragent string = `<document type="freeswitch/xml" encoding="UTF-8">
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
const UseragentA1hash string = `<document type="freeswitch/xml" encoding="UTF-8">
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
const UseragentReverse string = `<document type="freeswitch/xml" encoding="UTF-8">
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

//gateway
