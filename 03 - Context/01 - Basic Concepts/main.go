package main

import (
	"context"
	"fmt"
	"time"
)

func bookHotel(newContext context.Context) {
	select {
	case <-newContext.Done():
		fmt.Println("Hotel booking cancelled. Timeout reached")
		return
	case <-time.After(5 * time.Second):
		fmt.Println("Hotel booked.")
		return
	}
}

func main() {
	newBooking := context.Background()
	newBooking, cancel := context.WithTimeout(newBooking, 3*time.Second)
	defer cancel()

	bookHotel(newBooking)
}
