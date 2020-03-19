package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("program to listen to OS signals started")

	done := make(chan bool, 1)

	terminationSignals := make(chan os.Signal, 1)
	signal.Notify(terminationSignals, syscall.SIGINT, syscall.SIGTERM)
	go terminationSignalsHandler(terminationSignals, done)

	otherSignals := make(chan os.Signal, 1)
	signal.Notify(otherSignals, syscall.SIGWINCH)
	go otherSignalsHandler(otherSignals)

	<-done
}

func terminationSignalsHandler(terminationSignals chan os.Signal, done chan bool) {
	<-terminationSignals
	fmt.Println("termination signal received, the program is going to exit soon")
	done <- true
}

func otherSignalsHandler(otherSignals chan os.Signal) {
	for {
		<-otherSignals
		fmt.Println("terminal resize signal received")
	}
}
