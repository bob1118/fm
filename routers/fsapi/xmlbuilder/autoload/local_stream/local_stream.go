package local_stream

import (
	"os"

	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
)

var defaultConfname, defaultConffile, defaultData string

func init() {
	defaultConfname = "local_stream.conf.xml"
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
	//	if data, e := os.ReadFile(defaultConffile); e != nil {
	//		err = e
	//	} else {
	//		defaultData = string(data)
	//	}
	defaultData = LOCAL_STREAM
	return defaultData, err
}
