package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"time"
)

func main() {
	notify("")
}

func notify(services ...string) {
	var g errgroup.Group

	for _, service := range services {
		// Copy the value from the service variable into a local variable to
		// avoid the bug described in the email
		s := service
		g.Go(func() error {
			fmt.Printf("Starting to notifing %s...\n", s)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
			fmt.Printf("Finished notifying %s...\n", s)
			return nil // or a real error if we had one
		})
	}

	err := g.Wait()
	if err != nil {
		fmt.Printf("Error notifying services: %v\n", err)
		return
	}
	fmt.Println("All services notified successfully!")
}
