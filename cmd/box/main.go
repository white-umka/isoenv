package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	CPU_LIMIT = 50 // 50%
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

	pid := os.Getpid() // получаем id процесса 	
	os.WriteFile(cgroupPath+"/cgroup.procs", []byte(strconv.Itoa(pid)), 0644) // применяем лимиты

	fmt.Println("процесс успешно изолирован!")
}
