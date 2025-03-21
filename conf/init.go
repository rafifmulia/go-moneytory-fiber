package conf

import (
	"flag"

	"github.com/gofiber/fiber/v2"
)

func init() {
	flag.StringVar(&bindHost, "bind", bindHost, "Set bind host")
	flag.BoolVar(&preforkFlag, "prefork", false, "Enable prefork mode")
	flag.BoolVar(&childFlag, "child", false, "Enable child mode")
	flag.BoolVar(&debugFlag, "debug", false, "Enable debug mode")
	flag.StringVar(&cpuPprof, "cpuprofile", "", "Enable CPU profiling mode")
	flag.StringVar(&memPprof, "memprofile", "", "Enable memory profiling mode")
	flag.BoolVar(&httpPprof, "httplivepprof", false, "Enable live profiling with http access")
}

// Why separated from init func?
// Because it will make error "flag provided but not defined: -test.paniconexit0"
// when go test.
// Thats why flag.Parse should not be runned when go test.
// Why return bool true?
// Because in normal run, flag.Parse should be called before any getFlag func happens.
// If not, getFlag will return its default value.
// See cmd/server/main.go for implementation.
func ParseFlag() bool {
	flag.Parse()
	childFlag = fiber.IsChild()
	return true
}
