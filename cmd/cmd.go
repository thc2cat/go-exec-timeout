package main

import (
	"fmt"
	"os"
	"time"
)

// AS a cmd exemple, sleep more than timeout and prints argsv
func main() {
	time.Sleep(95 * time.Second)
	fmt.Print(os.Args[1:])
}
