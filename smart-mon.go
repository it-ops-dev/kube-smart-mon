package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	// Get env variables for port and path
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "9898" // Set a default port number if the environment variable is not set
	}
	path := os.Getenv("WORK_DIR")
	if path == "" {
		path = "/opt" // Set a default path if the environment variable is not set
	}

	// Start a Goroutine to execute the Bash script of smart output every 10 minutes
	go executeScript()

	//Check directory
	dirname, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("Could not get absolute path fo directory: %s: %s", dirname, err.Error())
	}
	// Listen on Port
	log.Printf("Serving %s on port %s", dirname, port)

	err = Serve(dirname, port)
	if err != nil {
		log.Fatalf("Could not serve directory: %s: %s", dirname, err.Error())
	}
}

func Serve(dirname string, port string) error {
	fs := http.FileServer(http.Dir(dirname))
	http.Handle("/", fs)

	return http.ListenAndServe(":"+port, nil)
}

func executeScript() {
	// Get env variables for sleep timer (in minutes)
	envVal := os.Getenv("SLEEP_TIMER")

	// Convert the value to an integer
	intValue, err := strconv.Atoi(envVal)
	if err != nil {
		intValue = 10
	}
	// Calculate the duration in minutes
	duration := time.Duration(intValue) * time.Minute
	// Define the command to execute the Bash script
	cmd := exec.Command("bash", "/opt/smart-mon-script.sh")

	// Set the environment for the command (optional)
	cmd.Env = os.Environ()

	// Start an infinite loop to execute the script every 10 minutes
	for {
		// Execute the command and check for any errors
		err := cmd.Run()
		if err != nil {
			log.Println("Failed to execute the script smart mon bash script:", err)
		}

		// Wait for duration minutes before executing the script again
		time.Sleep(duration)
	}
}
