package main

// Usage :  build , and shadow the cmd you want to timeout.
// rename original cmd as cmd.orig,
// rename this tool as cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	waits := getenv("TIMEOUT", "90")
	wait, err := strconv.Atoi(waits)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(wait)*time.Second)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	// Create the command with our context
	cmdargs := os.Args[1:]
	cmd := exec.CommandContext(ctx, os.Args[0]+".orig", cmdargs...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

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
	fmt.Print(string(out))
	if err != nil {
		log.Println("Non-zero exit code:", err)
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
