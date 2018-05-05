package main

import (
	"github.com/gen2brain/beeep"
	"log"
	"time"
)

func main() {
	ticker := time.NewTicker(20 * time.Minute).C
	doneChan := make(chan bool)
	go func() {
		log.Print("Starting the Countdown")
		time.Sleep(10 * time.Hour)
		doneChan <- true
	}()
	for {
		select {
		case <- ticker :
			err := beeep.Notify("Check Your Posture", "Look Away for 20 seconds", "posture.png")
			if err != nil {
				log.Panic(err)
			}
		case <- doneChan:
			log.Print("Timer Stopped")
			return
		}
	}



}