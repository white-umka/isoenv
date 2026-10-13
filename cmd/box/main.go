package main

import (
	"isoenv/internal/cgroup"
	"isoenv/internal/chroot"
	"isoenv/internal/symlinks"
	"isoenv/internal/modes"
)

func main() {
	cgroup.Cgroup()  // инициализируем cgroup и задаем лимиты
	mode.ChoosingMode()
	chroot.Chroot()  // меняем корневой каталог через chroot
	symlinks.Clear() // подключаем symlink для clear
}
