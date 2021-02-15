package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ores := []string{"ore1", "ore2", "ore3"}
// 	oreChan := make(chan string, 3)

// 	go func(mines []string) {
// 		for _, mine := range mines {
// 			fmt.Printf("Send %v from Finder \n", mine)
// 			oreChan <- mine
// 		}
// 	}(ores)

// 	go func() {
// 		for i := 3; i < 3; i++ {
// 			ore := <-oreChan
// 			fmt.Printf("Received %v from Finder\n", ore)
// 		}
// 	}()
// 	<-time.After(time.Second * 5) // Again, ignore this for now
// }

func main() {
	bufferedChan := make(chan string, 3)
	go func() {
		bufferedChan <- "first"
		fmt.Println("Sent 1st")
		bufferedChan <- "second"
		fmt.Println("Sent 2nd")
		bufferedChan <- "third"
		fmt.Println("Sent 3rd")
	}()
	<-time.After(time.Second * 1)
	go func() {
		firstRead := <-bufferedChan
		fmt.Println("Receiving..")
		fmt.Println(firstRead)
		secondRead := <-bufferedChan
		fmt.Println(secondRead)
		thirdRead := <-bufferedChan
		fmt.Println(thirdRead)
	}()
	<-time.After(time.Second * 1)
}
