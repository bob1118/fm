package xmlbuilder

import (
	"runtime"

	"github.com/bob1118/fm/config/fmconfig"
)

var defaultDirectory string
var defaultFreeswitchXml string

func init() {
	SetDefaultDirectory(fmconfig.CFG.Runtime.ConfDirectory)
	SetDefaultFreeswitchXml("freeswitch.xml")
}

func SetDefaultDirectory(dir string) {
	if len(dir) > 0 {
		defaultDirectory = dir
	} else {
		sysType := runtime.GOOS
		switch sysType {
		case "linux":
			defaultDirectory = "/etc/freeswitch/"
		case "windows":
			defaultDirectory = "C:/Program Files/FreeSWITCH/conf/"
		}
	}
}

func GetDefaultDirectory() (dir string) {
	return defaultDirectory
}

func SetDefaultFreeswitchXml(s string) {
	if len(s) > 0 {
		defaultFreeswitchXml = s
	} else {
		defaultFreeswitchXml = "freeswitch.xml"
	}
}
