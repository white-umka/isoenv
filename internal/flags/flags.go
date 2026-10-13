package flags

import (
	"flag"
)

func ParseFlags() bool {
	mode := flag.Bool("parent", false, "parent mode: default - false")
	// вызываем без флага - child
	// вызываем с флагом - parent
	
	flag.Parse()
	
	return *mode
}