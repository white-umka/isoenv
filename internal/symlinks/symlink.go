package symlinks

import (
	"log"
	"fmt"
	"os"
)

// симлинк для clear
func Clear() {
	// export PATH=/bin (чтобы рабоали ls, mkdir и т.д.)
	if err := os.Setenv("PATH", "/bin"); err != nil {
		log.Fatal(err)
	}

	_, err := os.Stat("/bin/clear") // проверяем на наличие симлинка clear (если он уже есть не выполняем код ниже)
	if err != nil {
		fmt.Println("симлинк на clear уже существует")
	} else {
		// создаем симлинк (псевдоним) для clear
		if err = os.Symlink("/bin/busybox", "/bin/clear"); err != nil {
			fmt.Println("не удалось создать симлинк для clear:\n", err)
		} 
	}
}