package odbc_cdr

import "github.com/gin-gonic/gin"

//freeswitch mod_odbc_cdr configuration.
//default configuration is file odbc_cdr.conf.xml

//1，request
// http://10.10.10.250/fsapi
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=odbc_cdr.conf]

//2，response
// <document type="freeswitch/xml">
//   <section name="configuration">
////     <configuration name="odbc_cdr.conf" description="ODBC CDR Configuration">
//       <settings>
//          <!--ADD your parameters here-->
//       </settings>
////     </configuration>
//   </section>
// </document>

//read configuration from file, and then write into db.
func MakeDefaultConfiguration() {}

//read configuration from db.
func ReadConfiguration(c *gin.Context) (b string) { return "" }

//write configuration into db.
func WriteConfiguration(s string) {}
