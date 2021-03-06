package autoload

import (
	"os"

	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
)

var defaultConfname, defaultConffile, defaultData string

func init() {}

//MakeDefaultConfiguration.
func MakeDefaultConfiguration() {}

//ReadDefaultConfiguration from file.
func ReadDefaultConfiguration(n string) (s string, e error) {
	var err error

	defaultConfname = n
	defaultConffile = xmlbuilder.GetDefaultDirectory() + `autoload_configs/` + defaultConfname

	if _, e := os.Stat(defaultConffile); os.IsNotExist(e) {
		return defaultData, e
	}
	if data, e := os.ReadFile(defaultConffile); e != nil {
		err = e
	} else {
		defaultData = string(data)
	}
	return defaultData, err
}
