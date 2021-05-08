package switch_main

import (
	"bytes"
	"os"

	"github.com/bob1118/fm/esl/eslclient"
	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
)

var defaultConfname, defaultConffile, defaultData string

func init() {
	defaultConfname = "switch.conf.xml"
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
			[]byte(`<param name="sessions-per-second" value="30"/>`),
			[]byte(`<param name="sessions-per-second" value="100"/>`))
		defaultData = string(data)
		data = bytes.ReplaceAll(data,
			[]byte(`<param name="loglevel" value="debug"/>`),
			[]byte(`<param name="loglevel" value="info"/>`))
		defaultData = string(data)
		/* function BuildPersonalConf() set core-db-dsn before switch boot.
		data = bytes.ReplaceAll(data,
			[]byte(`<!-- <param name="core-db-dsn" value="dsn:username:password" /> -->`),
			[]byte(`<param name="core-db-dsn" value="$${pg_handle}"/>`))
		*/
		defaultData = string(data)
	}
	go func() {
		eslclient.CHfsisrun <- true
	}()
	return defaultData, err
}
