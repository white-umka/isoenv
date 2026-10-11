package main

import (
	"fmt"
	"log"
	"os"

	//"os/user"
	"strconv"
	"syscall"
)

const (
	CPU_LIMIT    = 50  // 50%
	MEMORY_LIMIT = 256 // 256mb
)

func main() {
	parentPath := "/sys/fs/cgroup/isoenv"
	os.MkdirAll(parentPath, 0755)

	err := os.WriteFile(parentPath+"/cgroup.subtree_control", []byte("+cpu +memory"), 0644)
	if err != nil {
		fmt.Printf("ошибка активации подсистем: %v, \nзапустите под sudo \n", err)
		os.Exit(1)
	}

	// создаем cgroup
	cgroupPath := parentPath + "/test"
	os.MkdirAll(cgroupPath, 0755)

	// задаем лимиты ресурсов для этой cgroup (test)
	os.WriteFile(cgroupPath+"/cpu.max", []byte("50000 100000"), 0644)
	os.WriteFile(cgroupPath+"/memory.max", []byte("268435456"), 0644)

	pid := os.Getpid()                                                        // получаем id процесса
	os.WriteFile(cgroupPath+"/cgroup.procs", []byte(strconv.Itoa(pid)), 0644) // применяем лимиты

	// получаем имя юзера
	username := os.Getenv("SUDO_USER")

	// chroot (home/username/isoenv/rootfs/alpine теперь является корневым каталогом "/")
	if err = syscall.Chroot("/home/" + username + "/isoenv/rootfs/alpine"); err != nil {
		log.Fatal(err)
	}

	// заходим в новый каталог (аналог cd)
	if err = syscall.Chdir("/"); err != nil {
		log.Fatal(err)
	}

	// читаем новый корень (логируем)
	files, err := os.ReadDir("/")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("содержимое корня /")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	// export PATH=/bin (чтобы рабоали ls, mkdir и т.д.)
	if err = os.Setenv("PATH", "/bin"); err != nil {
		log.Fatal(err)
	}

	// создаем симлинк (псевдоним) для clear
	if err = os.Symlink("/bin/busybox", "/bin/clear"); err != nil {
		fmt.Println("не удалось создать линк для clear:\n", err)
	} else {fmt.Println("симлинк на clear создан успешно!")}

	// запускаем /bin/sh (Environ передаст переменные окружения в формате key:value)
	if err := syscall.Exec("/bin/sh", []string{"/bin/sh"}, os.Environ()); err != nil {
		log.Fatal(err)
	}

	fmt.Println(username)
	fmt.Println(pid)
	fmt.Println("процесс успешно изолирован!")
}
