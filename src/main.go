package main

import (
	"os"
	"os/signal"
	"time"
)

func main() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt) // 设置信号量 当接受 Ctrl+C 的时候出发
	<-stopChan                            // wait for SIGINT

	time.Sleep(time.Second * 3) // test
}
