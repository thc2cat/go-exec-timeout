package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	time.Sleep(95 * time.Second)
	fmt.Print(os.Args[1:])
}
