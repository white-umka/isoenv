package chroot

import (
	"log"
	"os"
	"syscall"
)

func Chroot() {
	// получаем имя юзера
	username := os.Getenv("SUDO_USER")

	// chroot (home/username/isoenv/rootfs/alpine теперь является корневым каталогом "/")
	if err := syscall.Chroot("/home/" + username + "/isoenv/rootfs/alpine"); err != nil {
		log.Fatal(err)
	}

	// заходим в новый каталог (аналог cd)
	if err := syscall.Chdir("/"); err != nil {
		log.Fatal(err)
	}

	if err := os.Setenv("PATH", "/bin"); err != nil {
		log.Fatalf("не удалось уставноить PATH: %v", err)
	}
	// запускаем /bin/sh (Environ передаст переменные окружения в формате key:value)
	if err := syscall.Exec("/bin/sh", []string{"/bin/sh"}, os.Environ()); err != nil {
		log.Fatal(err)
	}

	/*
		// читаем новый корень (логируем)
		files, err := os.ReadDir("/")
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("содержимое корня /")
		for _, file := range files {
			fmt.Println(file.Name())
		} */
}
