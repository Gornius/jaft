package main

import (
	"embed"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/gopxl/beep"
	"github.com/gopxl/beep/mp3"
	"github.com/gopxl/beep/speaker"
)

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("not enough arguments")
	}

	durationString := os.Args[1]
	duration, err := time.ParseDuration(durationString)
	if err != nil {
		return err
	}

	time.Sleep(duration)

	if err := beepIndefinitely(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)
	<-ch

	return nil
}

//go:embed embed
var embedFs embed.FS

func beepIndefinitely() error {
	f, err := embedFs.Open("embed/timer_end.mp3")
	if err != nil {
		return err
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		return err
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	loop := beep.Loop(-1, streamer)
	done := make(chan bool)
	speaker.Play(beep.Seq(loop, beep.Callback(func() {
		done <- true
	})))

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
		fmt.Println("Usage: jaft DURATION (e.g. \"300ms\", \"-1.5h\" or \"2h45m\")")
		os.Exit(1)
	}
}
