package main

import (
	"time"
)

func main() {
	client := NewClient(5 * time.Second)
	config := &Config{
		apiClient: client,
	}

	replStart(config)
}