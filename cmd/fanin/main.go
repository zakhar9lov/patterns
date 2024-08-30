package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// channels for returning temperature from api
	ch1 := make(chan float64)
	ch2 := make(chan float64)
	ch3 := make(chan float64)

	// Fetch data from SUPER Weather API
	go func(ctx context.Context, ch chan<- float64) {
		fmt.Println("Calling the API SUPER Weather:")
		time.Sleep(1 * time.Second)
		select {
		case ch <- 30.0 + rand.Float64()*10: // random temperature
		case <-ctx.Done():
			fmt.Println("The response timed out. No response received from SUPER Weather")
			return
		}
	}(ctx, ch1)

	// Fetch data from CLOUD Weather API
	go func(ctx context.Context, ch chan<- float64) {
		fmt.Println("Calling the API CLOUD Weather:")
		time.Sleep(2 * time.Second)
		select {
		case ch <- 30.0 + rand.Float64()*10:
		case <-ctx.Done():
			fmt.Println("The response timed out. No response received from CLOUD Weather")
			return
		}
	}(ctx, ch2)

	// Fetch data from SLOW Weather API (Response time is too long)
	go func(ctx context.Context, ch chan<- float64) {
		fmt.Println("Calling the API SLOW Weather:")
		time.Sleep(10 * time.Second)
		select {
		case ch <- rand.Float64() * 10:
		case <-ctx.Done():
			fmt.Println("The response timed out. No response received from SLOW Weather")
			return
		}
	}(ctx, ch3)

	result := fanIn(context.Background(), ch1, ch2, ch3)

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()

	go func(ctx context.Context) {
		for {
			select {
			case temp := <-result:
				fmt.Println("TEMPERATURE: ", temp)

			case <-ctx.Done():
				return
			}
		}
	}(ctx2)

	time.Sleep(15 * time.Second)
} // main

// Receives multiple channels and returns one channel to which data from the others is transferred
func fanIn(ctx context.Context, channels ...<-chan float64) <-chan float64 {
	generalChannel := make(chan float64)

	var wg sync.WaitGroup
	wg.Add(len(channels))

	for _, ch := range channels {
		ch := ch // lvalue ch — is local var. rvalue ch — is for-var

		go func() {
			defer wg.Done()
			select {
			case content := <-ch:
				generalChannel <- content
			case <-ctx.Done():
				return
			}
		}()
	}

	go func() {
		wg.Wait()

		close(generalChannel)
	}()

	return generalChannel
}
