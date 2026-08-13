package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func run() error {
	if len(os.Args) < 1 {
		return fmt.Errorf("no time has been declared")
	}

	durationString := os.Args[1]
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	go func() {
		<-ch
		os.Exit(0)
	}()

	time.Sleep(duration)
	for {
		fmt.Println("BEEP BEEP")
		time.Sleep(1 * time.Second)
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Usage: jaft duration (e.g. \"300ms\", \"-1.5h\" or \"2h45m\")")
		os.Exit(1)
	}
}
