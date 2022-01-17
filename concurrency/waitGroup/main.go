package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	notify("one", "two", "three", "four")
}

func notify(services ...string) {
	var wg sync.WaitGroup
	wg.Add(len(services))
	for _, service := range services {
		//wg.Add(1)
		go func(s string) {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Finished notifying %s...\n", s)
			wg.Done()
		}(service)
	}

	wg.Wait()
	fmt.Println("All services notified!")
}
