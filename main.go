package main

import (
	"githug.com/bob118/fm/config/fmconfig"
)

func main() {
	cfg := fmconfig.NewFmconfig()
	cfg.Read()
	cfg.Write()
}
