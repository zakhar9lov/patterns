package main

import (
	"context"
	"fmt"
	"time"

	. "patterns/internal/takefirstn"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Create a slice of 100 nums
	nums := GetRandomSliceInt(100)

	chValues := make(chan interface{})

	// Sending slice elements to the channel
	go SendToChannel(nums, chValues)

	// Get the first N elements
	chFirstN := TakeFirstN(ctx, chValues, 15)

	fmt.Println("Chan contains:")
	for v := range chFirstN {
		fmt.Println(v.(int))
	}
}
