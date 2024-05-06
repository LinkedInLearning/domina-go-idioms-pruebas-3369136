package main

import (
	"context"
	"fmt"
)

func contextBackground() {
	err := httpCall(context.Background())
	if err != nil {
		fmt.Println("Error making request:", err)
	}
}
