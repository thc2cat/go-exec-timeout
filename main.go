package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	// We'll use ping as it will provide output and we can control how long it runs.
	cmd := exec.Command(os.Args[0]+".orig", os.Args...)

	// Use a bytes.Buffer to get the output
	var buf bytes.Buffer
	cmd.Stdout = &buf

	cmd.Start()

	// Use a channel to signal completion so we can use a select statement
	done := make(chan error)
	go func() { done <- cmd.Wait() }()

	// Start a timer
	timeout := time.After(30 * time.Second)

	// The select statement allows us to execute based on which channel
	// we get a message from first.
	select {
	case <-timeout:
		// Timeout happened first, kill the process and print a message.
		cmd.Process.Kill()
		fmt.Println("Command timed out : ", os.Args[1:])
	case err := <-done:
		// Command completed before timeout. Print output and error if it exists.
		fmt.Println(buf.String())
		if err != nil {
			fmt.Println(err)
		}
	}
}
