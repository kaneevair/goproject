package main

import (
	"fmt"
	"time"
)

const dateFormat = "02-01-2006 15:04"

func main() {
	fmt.Printf("Hello from Rebrain! %s", time.Now().Format(dateFormat))
}
