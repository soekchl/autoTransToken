package main

import (
	"autoTransToken/src/control"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	go control.CheckData()
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt) // 设置信号量 当接受 Ctrl+C 的时候出发

	fmt.Println("Server Closing...")
	<-stopChan // wait for SIGINT
	fmt.Println("Server Closed!")
}
