package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	DONE_MESSAGE    = "Coming soon !\n"
	RUNNING_MESSAGE = "Back in %s !\n"
)

func createFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}

func write(path, s string) error {
	fmt.Printf("[DEBUG] %s", s)

	fd, err := os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer fd.Close()
	_, err = fd.WriteString(s)
	if err != nil {
		return err
	}
	err = fd.Sync()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var p, ds string
	flag.StringVar(&ds, "duration", "00h15m00s", "Countdown duration")
	flag.StringVar(&p, "out-file", "/tmp/countdown.raw", "Path to OBS watchfile")
	flag.Parse()

	err := createFile(p)
	if err != nil {
		log.Fatal(err)
	}

	d, err := time.ParseDuration(ds)
	if err != nil {
		log.Fatal(err)
	}
	end := time.Now().Add(d)

	done := make(chan bool)
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	go func() {
		time.Sleep(d)
		done <- true
	}()

	for {
		select {
		case <-done:
			err = write(p, DONE_MESSAGE)
			if err != nil {
				log.Print(err)
			}
			return
		case <-ticker.C:
			res := fmt.Sprintf(RUNNING_MESSAGE, time.Until(end).Round(time.Second).String())
			err = write(p, res)
			if err != nil {
				log.Print(err)
				return
			}
		}
	}
}
