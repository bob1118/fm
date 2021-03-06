package odbc_cdr

import (
	"bytes"
	"log"
	"os"
	"runtime"

	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
)

//freeswitch mod_odbc_cdr configuration.
//default configuration is file odbc_cdr.conf.xml

//1，request
// http://10.10.10.250/fsapi
//data: [hostname=bob-office&section=configuration&tag_name=configuration&key_name=name&key_value=odbc_cdr.conf]

//2，response
// <document type="freeswitch/xml">
//   <section name="configuration">
//     <configuration name="odbc_cdr.conf" description="ODBC CDR Configuration">
//       <settings>
//          <!--ADD your parameters here-->
//       </settings>
//     </configuration>
//   </section>
// </document>

var defaultConfname, defaultConffile, defaultData string

func init() {
	defaultConfname = "odbc_cdr.conf.xml"
	defaultConffile = xmlbuilder.GetDefaultDirectory() + `autoload_configs/` + defaultConfname
}

//MakeDefaultConfiguration.
func MakeDefaultConfiguration() {
	if e := os.WriteFile(defaultConffile, []byte(ODBC_CDR), 0660); e != nil {
		log.Println(e)
	}
}

//ReadConfiguration from file.
func ReadConfiguration() (s string, e error) {
	var err error

	if _, e := os.Stat(defaultConffile); os.IsNotExist(e) {
		MakeDefaultConfiguration()
	}
	if data, e := os.ReadFile(defaultConffile); e != nil {
		err = e
	} else {
		data = bytes.ReplaceAll(data,
			[]byte(`<param name="odbc-dsn" value="pgsql://hostaddr=192.168.0.100 dbname=freeswitch user=freeswitch password='freeswitch' options='-c client_min_messages=NOTICE'"/>`),
			[]byte(`<param name="odbc-dsn" value="$${pg_handle}"/>`))
		switch runtime.GOOS {
		case "windows":
			data = bytes.ReplaceAll(data,
				[]byte(`<param name="csv-path" value="/usr/local/freeswitch/log/odbc_cdr"/>`),
				[]byte(`<param name="csv-path" value="C:/Program Files/FreeSWITCH/log/odbc_cdr"/>`))
			data = bytes.ReplaceAll(data,
				[]byte(`<param name="csv-path-on-fail" value="/usr/local/freeswitch/log/odbc_cdr/failed"/>`),
				[]byte(`<param name="csv-path-on-fail" value="C:/Program Files/FreeSWITCH/log/odbc_cdr/failed"/>`))
		case "linux":
			data = bytes.ReplaceAll(data,
				[]byte(`<param name="csv-path" value="/usr/local/freeswitch/log/odbc_cdr"/>`),
				[]byte(`<param name="csv-path" value="/var/log/freeswitch/odbc_cdr"/>`))
			data = bytes.ReplaceAll(data,
				[]byte(`<param name="csv-path-on-fail" value="/usr/local/freeswitch/log/odbc_cdr/failed"/>`),
				[]byte(`<param name="csv-path-on-fail" value="/var/log/freeswitch/odbc_cdr/failed"/>`))
		}
		defaultData = string(data)
	}
	return defaultData, err
}
