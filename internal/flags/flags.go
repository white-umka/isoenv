package flags

import (
	"flag"
)

func ParseFlags() bool {
	mode := flag.Bool("child", false, "child mode: default - true")
	// вызываем без флага - parent
	// вызываем с флагом - child
	
	flag.Parse()
	
	return *mode
}