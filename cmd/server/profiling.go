package main

import (
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

var (
	cpuPprof     string
	memPprof     string
	cpuPprofFile *os.File
	memPprofFile *os.File
)

// Profiling in this package is intended for
// enabled runtime profiling without http access such as
// import _ "net/http/pprof".
func profiling() {
	var err error
	if cpuPprof != "" {
		cpuPprofFile, err = os.Create(cpuPprof)
		if err != nil {
			log.Fatal(err)
		}
		runtime.SetCPUProfileRate(500)
		err = pprof.StartCPUProfile(cpuPprofFile)
		if err != nil {
			log.Fatal(err)
		}
	}
	if memPprof != "" {
		memPprofFile, err = os.Create(memPprof)
		if err != nil {
			log.Fatal(err)
		}
		runtime.GC()
		err = pprof.WriteHeapProfile(memPprofFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}
