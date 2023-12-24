package main

import (
	app "github.com/ayberktandogan/melody/app/melody"
	"github.com/ayberktandogan/melody/config"
)

var version string

func main() {
	config.InitConfig(version)
	app.Main()
}
