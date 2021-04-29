package sofia

//PROFILES
var PROFILE string = `  <configuration name="sofia.conf" description="sofia Endpoint">
    <global_settings>
    </global_settings>
    <profiles>
    %s
    </profiles>
  </configuration>`

//GATEWAY, profiles->profile->gateways->gateway.
var GATEWAY string = `  <gateway name="%s">
   <param name="username" value="%s"/>
   <param name="realm" value="%s"/>
   <param name="from-user" value="%s"/>
   <param name="from-domain" value="%s"/>
   <param name="password" value="%s"/>
   <param name="extension" value="%s"/>
   <param name="proxy" value="%s"/>
   <param name="register-proxy" value="%s"/>
   <param name="expire-seconds" value="%s"/>
   <param name="register" value="%s"/>
   <!--<param name="register-transport" value="udp"/>-->
   <!--<param name="retry-seconds" value="30"/>-->
   <param name="caller-id-in-from" value="%s"/>
   <!--<param name="contact-params" value=""/>-->
   <param name="extension-in-contact" value="%s"/>
   <param name="ping" value="%s"/>
   <!--<param name="cid-type" value="rpid"/>-->
   <!--<param name="rfc-5626" value="true"/>-->
   <!--<param name="reg-id" value="1"/>-->
  </gateway>
  
  `
