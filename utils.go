package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"syscall"
)

func Singleton() {
	pid := fmt.Sprint(os.Getpid())
	wd, _ := os.Getwd()
	fp := wd + "/my.pid"
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		//file is not exist
		fmt.Println("Create pid file:", fp, pid)
		f, _ := os.Create(fp)
		f.WriteString(pid)
		f.Close()
		return
	}

	f, err := os.OpenFile(fp, os.O_RDWR, 0666)
	if err != nil {
		panic(err.Error())
	}
	defer f.Close()

	x, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err.Error())
	}
	iPid, _ := strconv.Atoi(fmt.Sprintf("%s", x))
	process, err := os.FindProcess(iPid)

	if err != nil {
		panic(err)
	}
	err = process.Signal(syscall.Signal(0))
	//fmt.Printf("process.Signal on pid %d returned: %v\n", iPid, err)
	if err != nil {
		f.Close()
		os.Remove(fp)
		f, _ := os.Create(fp)
		f.WriteString(pid)
		f.Close()
		fmt.Println("process:", pid)

	} else {
		fmt.Println("Stopped! One instance is working.PID:", iPid)
		os.Exit(0)
	}

}
