package main

import (
	"githug.com/bob118/fm/config/fmconfig"
)

func main() {
	fmconfig.NewFmconfig().Read(fmconfig.CFGFILE)
}
