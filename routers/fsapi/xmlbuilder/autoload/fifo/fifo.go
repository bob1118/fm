package fifo

import (
	"bytes"
	"os"

	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
)

var defaultConfname, defaultConffile, defaultData string

func init() {
	defaultConfname = "fifo.conf.xml"
	defaultConffile = xmlbuilder.GetDefaultDirectory() + `autoload_configs/` + defaultConfname
}

//MakeDefaultConfiguration.
func MakeDefaultConfiguration() {}

//ReadConfiguration from file.
func ReadConfiguration() (s string, e error) {
	var err error

	if _, e := os.Stat(defaultConffile); os.IsNotExist(e) {
		return defaultData, e
	}
	if data, e := os.ReadFile(defaultConffile); e != nil {
		err = e
	} else {
		data = bytes.ReplaceAll(data,
			[]byte(`<param name="delete-all-outbound-member-on-startup" value="false"/>`),
			[]byte(`<param name="odbc-dsn" value="$${pg_handle}"/>
    <param name="delete-all-outbound-member-on-startup" value="false"/>`))
		//
		defaultData = string(data)
	}
	return defaultData, err
}
