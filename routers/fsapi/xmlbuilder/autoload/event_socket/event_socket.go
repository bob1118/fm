package event_socket

import (
	"bytes"
	"os"

	"github.com/bob1118/fm/routers/fsapi/xmlbuilder"
)

var defaultConfname, defaultConffile, defaultData string

func init() {
	defaultConfname = "event_socket.conf.xml"
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
			[]byte(`<param name="listen-ip" value="::"/>`),
			[]byte(`<param name="listen-ip" value="::"/>`))
		data = bytes.ReplaceAll(data,
			[]byte(`<param name="listen-port" value="8021"/>`),
			[]byte(`<param name="listen-port" value="8021"/>`))
		data = bytes.ReplaceAll(data,
			[]byte(`<param name="password" value="ClueCon"/>`),
			[]byte(`<param name="password" value="ClueCon"/>`))
		data = bytes.ReplaceAll(data,
			[]byte(`<!--<param name="apply-inbound-acl" value="loopback.auto"/>-->`),
			[]byte(`<param name="apply-inbound-acl" value="loopback.auto"/>`))

		defaultData = string(data)
	}
	//
	// go func() {
	// 	eslclient.CHfsisrun <- true
	// }()
	return defaultData, err
}
