package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

const wait = 90 //  Seconds before timeout

func main() {
	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), wait*time.Second)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	// Create the command with our context
	cmdargs := os.Args[1:]
	cmd := exec.CommandContext(ctx, os.Args[0]+".orig", cmdargs...)

	// This time we can simply use Output() to get the result.
	out, err := cmd.Output()

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Command timed out:", os.Args)
		return
	}

	// If there's no context error, we know the command completed (or errored).
	fmt.Println(string(out))
	if err != nil {
		log.Println("Non-zero exit code:", err)
	}
}
