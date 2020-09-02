package main

import (
	_ "dota.city/m/boot"
	_ "dota.city/m/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
