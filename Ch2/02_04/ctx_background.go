package main

import (
	"context"
	"fmt"
)

func contextBackground() {
	err := httpCall(context.Background())
	if err != nil {
		fmt.Println("Context Background: Error making request:", err)
		return
	}

	fmt.Println("Context with Timeout: Success!")
}
