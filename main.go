package main

import (
	"fmt"
	"time"
)

func main() {
	const dateFormat = "02-01-2006 15:04"

	fmt.Printf("Hello from Rebrain! %s", time.Now().Format(dateFormat))
}
