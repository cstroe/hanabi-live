// +build !windows

package main

import (
	"os"
	"os/signal"
	"syscall"
)

// We want to be able to control the server from the command-line
// without having to send chat messages to the lobby
// This function is meant to be called in a new goroutine
func signalInit() {
	signalChannel := make(chan os.Signal, 1)

	// Catch some user defined signals
	// (these are meant to be used for application-specific purposes)
	signal.Notify(signalChannel, syscall.SIGUSR1, syscall.SIGUSR2, syscall.SIGUSR3)

	for {
		signal := <-signalChannel
		if signal == syscall.SIGUSR1 {
			// Gracefully restart
			graceful(true)
		} else if signal == syscall.SIGUSR2 {
			// Shutdown
			graceful(false)
		} else if signal == syscall.SIGUSR3 {
			// Debug
			debug()
		}
	}
}
