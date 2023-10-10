package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Check if sourceLocation and outputLocation environment variables are set.
	sourceLocation := os.Getenv("SOURCE_LOCATION")
	outputLocation := os.Getenv("OUTPUT_LOCATION")

	if sourceLocation == "" || outputLocation == "" {
		fmt.Println("SOURCE_LOCATION and OUTPUT_LOCATION environment variables must be set")
		return
	}

	// Read the content from the source file.
	sourceContent, err := ioutil.ReadFile(sourceLocation)
	if err != nil {
		fmt.Printf("Error reading from %s: %v\n", sourceLocation, err)
		return
	}

	// Write the content to the output file.
	err = ioutil.WriteFile(outputLocation, sourceContent, 0644)
	if err != nil {
		fmt.Printf("Error writing to %s: %v\n", outputLocation, err)
		return
	}

	fmt.Printf("Successfully copied content from %s to %s\n", sourceLocation, outputLocation)
}
