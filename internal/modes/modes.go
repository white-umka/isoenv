package mode

import (
	"fmt"
	"os"

	"isoenv/internal/flags"
)

func ChoosingMode() {
	mode := flags.ParseFlags()
	switch mode {
		case true:
			fmt.Println("true")
		case false:
			fmt.Println("false")
		default:
			fmt.Println("произошла ошибка при работе с флагами")
			os.Exit(1)
	}
}
