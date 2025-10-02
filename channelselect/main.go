package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v7"
)

func main() {
	bpChannel := make(chan int)
	quitChannel := make(chan bool)
	go func() {
		t := time.NewTicker(1 * time.Second)

		defer t.Stop()

		for range t.C {
			bpChannel <- gofakeit.IntRange(60, 230)

		}
	}()

	//quit producer
	go func() {
		fmt.Println("Press ENTER to quit...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		quitChannel <- true
	}()
	//receve channel data

	for {
		select {
		case val := <-bpChannel:
			fmt.Printf("Blood Pressure: %d mmHg\n", val)

		case <-quitChannel:
			fmt.Println("Exiting...")
			return
		}
	}
}
