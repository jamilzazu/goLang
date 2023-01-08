package main

import (
	"context"
	"fmt"
)

func bookHotel(newContext context.Context) {
	token := newContext.Value("token")
	fmt.Println(token)
}

func main() {
	newBooking := context.Background()
	newBooking = context.WithValue(newBooking, "token", "password")
	bookHotel(newBooking)
}
