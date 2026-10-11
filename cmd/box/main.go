package main

import (
	"fmt"

	"isoenv/internal/cgroup"
	"isoenv/internal/chroot"
	"isoenv/internal/symlinks"
)

func main() {
	cgroup.Cgroup()  // инициализируем cgroup и задаем лимиты
	chroot.Chroot()  // меняем корневой каталог через chroot
	symlinks.Clear() // подключаем symlink для clear
	
	fmt.Println("процесс успешно изолирован!")
}
