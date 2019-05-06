package main

import (
	"net/http"
	"os"
)

func LogHandler(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("events.log", os.O_RDONLY, 0)
	defer f.Close()
	if err != nil {
		errLog.Println("Open log file error:", err)
		http.Error(w, "Error", http.StatusExpectationFailed)
	}

	buf := make([]byte, 1024000)

	for {
		n, _ := f.Read(buf)
		if 0 == n {
			break
		}
		//os.Stdout.Write(buf[:n])
		w.Write(buf[:n])
	}

}
