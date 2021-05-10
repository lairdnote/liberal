package main

import (
	"github.com/lairdnote/liberal/conf"
	"github.com/lairdnote/liberal/server"
)

func main() {
	conf.init()
	r := server.NewRoute()
	r.run(":3000")
}
