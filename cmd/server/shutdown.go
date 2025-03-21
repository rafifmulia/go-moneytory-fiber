package main

import (
	"log"
	"os"
	"os/signal"
	"restfulapi/driver"
	"runtime/pprof"
	"syscall"
)

// Handle signal termination.
func handleSignals() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
	for sig := range c {
		log.Printf("Received signal %s", sig.String())
		signal.Stop(c)
		closeResources()
		os.Exit(15) // Exit as SIGTERM.
	}
}

// Global resources that must be closed
func closeResources() {
	if db := driver.ExportDbHandle(); db != nil {
		log.Println("Closing database connection")
		db.Close()
	}
	if cpuPprof != "" && cpuPprofFile != nil {
		log.Println("Closing cpu profiling")
		pprof.StopCPUProfile()
		cpuPprofFile.Close()
	}
	if memPprof != "" && memPprofFile != nil {
		log.Println("Closing memory profiling")
		memPprofFile.Close()
	}
}
