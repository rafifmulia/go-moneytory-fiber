package conf

// These flag is intended for enable profiling.
// Either using "runtime/pprof" package, or with http access such as
// import _ "net/http/pprof",

var (
	cpuPprof  string = ""    // Write CPU profile to this file.
	memPprof  string = ""    // Write memory profile to this file.
	httpPprof bool   = false // Live profiling with http access.
)

func GetMemPprof() string {
	return memPprof
}

func GetCpuPprof() string {
	return cpuPprof
}

func GetHttpPprof() bool {
	return httpPprof
}
