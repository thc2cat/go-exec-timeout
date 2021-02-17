package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	time.Sleep(10 * time.Second)
	fmt.Println(os.Args[1:])
}
