package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/eiannone/keyboard"
)

var serverProcess *exec.Cmd
var serverArgs []string // Store command and arguments to restart

// Function to start the server with given arguments
func startServer() {
	if len(serverArgs) == 0 {
		fmt.Println("No command provided to run.")
		return
	}

	serverProcess = exec.Command(serverArgs[0], serverArgs[1:]...)
	serverProcess.Stdout = os.Stdout // Pipe server output to main CLI stdout
	serverProcess.Stderr = os.Stderr

	err := serverProcess.Start()
	if err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
		return
	}

	fmt.Println("Server started with PID:", serverProcess.Process.Pid)

	// Wait for the server process to exit
	go func() {
		err := serverProcess.Wait()
		if err != nil {
			fmt.Printf("Server process exited with error: %s\n", err)
		} else {
			fmt.Println("Server process exited normally.")
		}
	}()
}

// Function to stop the server
func stopServer() {
	if serverProcess != nil && serverProcess.Process != nil {
		// Kill the server process
		err := serverProcess.Process.Kill()
		if err != nil {
			fmt.Printf("Failed to kill server: %s\n", err)
		} else {
			// Wait for the process to actually terminate
			err := serverProcess.Wait()
			if err != nil {
				fmt.Printf("Failed to wait for process termination: %s\n", err)
			} else {
				fmt.Println("Server stopped successfully.")
			}
		}
		serverProcess = nil
	}
}

// Function to restart the server
func restartServer() {
	fmt.Println("Attempting to stop and restart the server...")
	stopServer()
	time.Sleep(1 * time.Second) // Add small delay for the process to terminate properly
	startServer()
}

func main() {
	// Parse CLI arguments to get the command to run
	if len(os.Args) < 2 {
		fmt.Println("Usage: looprun <command> [args...]")
		os.Exit(1)
	}

	// Store the command and its arguments
	serverArgs = os.Args[1:]

	// Start server initially with provided command and args
	startServer()

	// Listen for OS interrupts (e.g., Ctrl+C) to clean up the server process
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		stopServer()
		os.Exit(0)
	}()

	// Initialize keyboard listener
	if err := keyboard.Open(); err != nil {
		fmt.Printf("Failed to open keyboard listener: %s\n", err)
		os.Exit(1)
	}
	defer keyboard.Close()

	fmt.Println("Press 'Ctrl+R' to restart the server, or 'Ctrl+C' to quit.")

	// Command loop to listen for input commands
	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			fmt.Printf("Failed to read keyboard input: %s\n", err)
			continue
		}

		if key == keyboard.KeyCtrlR {
			fmt.Println("Restarting server...")
			restartServer()
		} else if key == keyboard.KeyCtrlC {
			fmt.Println("Exiting...")
			stopServer()
			os.Exit(0)
		} else {
			fmt.Println("Invalid input, press 'Ctrl+R' to restart or 'Ctrl+C' to quit.")
		}
	}
}
