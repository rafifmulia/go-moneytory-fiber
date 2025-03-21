package main

import (
	"restfulapi/conf"
	"restfulapi/router"
)

var (
	_        bool = conf.ParseFlag()
	bindHost string
)

func setFlags() {
	bindHost = conf.GetBindHost()
	cpuPprof = conf.GetCpuPprof()
	memPprof = conf.GetMemPprof()
}

func initResources() {
	setFlags()
	profiling()
	conf.InitDbConnX()
}

func main() {
	go handleSignals()
	initResources()
	app := router.InitRouter()
	app.Listen(bindHost)
}
