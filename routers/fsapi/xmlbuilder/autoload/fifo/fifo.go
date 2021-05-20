package fifo

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/bob1118/fm/models"
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
		var strnew string
		if strings.EqualFold(runtime.GOOS, "windows") {
			strnew = `<param name="delete-all-outbound-member-on-startup" value="false"/>` + "\r\n" + `    <param name="odbc-dsn" value="$${pg_handle}"/>`
		} else {
			strnew = `<param name="delete-all-outbound-member-on-startup" value="false"/>` + "\n" + `    <param name="odbc-dsn" value="$${pg_handle}"/>`
		}
		data = bytes.ReplaceAll(data, []byte(`<param name="delete-all-outbound-member-on-startup" value="false"/>`), []byte(strnew))
		//
		if allfifos, e := buildXmlFifos(); err != nil {
			err = e
		} else {
			data = bytes.ReplaceAll(data, []byte(`<fifos>`), []byte(`<fifos>`+allfifos))
		}
		defaultData = string(data)
	}
	return defaultData, err
}

//buildXmlFifos only fifo no members.
func buildXmlFifos() (s string, e error) {
	var allfifo string
	var myerr error

	if fifos, err := models.GetFifos("1=1"); err != nil {
		myerr = err
	} else {
		for _, fifo := range fifos {
			if strings.EqualFold(runtime.GOOS, "windows") {
				allfifo += "\r\n" + fmt.Sprintf(FIFO, fifo.Fname, fifo.Fimportance, ``)
			} else {
				allfifo += "\n" + fmt.Sprintf(FIFO, fifo.Fname, fifo.Fimportance, ``)
			}
		}
	}
	return allfifo, myerr
}

//buildXmlFifosEx
// func buildXmlFifosEx() (s string, e error) {
// 	var allfifo string
// 	var myerr error

// 	if fifos, err := models.GetFifos("1=1"); err != nil {
// 		return ``, err
// 	} else {
// 		for _, fifo := range fifos {
// 			allfifomembers := ``
// 			fifoname := fifo.Fname
// 			condistion := fmt.Sprintf("fifo_name='%s'", fifoname)
// 			if fifomembers, err := models.GetFifomembers(condistion); err != nil {
// 				return ``, err
// 			} else {
// 				for _, fifomember := range fifomembers {
// 					allfifomembers += fmt.Sprintf(FIFOMEMBER, fifomember.Mtimeout, fifomember.Msimo, fifomember.Mlag, fifomember.Mstring)
// 				}
// 			}
// 			allfifo += fmt.Sprintf(FIFO, fifo.Fname, fifo.Fimportance, allfifomembers)
// 		}
// 	}
// 	return allfifo, myerr
// }
