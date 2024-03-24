package second

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestPrintSomething(t *testing.T) {
	// Save the original standard output (stdout)
	stdOut := os.Stdout

	// Create a new pipe to capture the output
	r, w, _ := os.Pipe()

	// Redirect the standard output to the write end of the pipe
	os.Stdout = w

	// Create a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup

	// Add 1 to the WaitGroup counter
	wg.Add(1)

	// Call the function being tested, passing the captured standard output and the WaitGroup
	PrintSomething("Hello World", &wg)

	// Wait for all goroutines to finish
	wg.Wait()

	// Close the write end of the pipe
	w.Close()

	// Read all data from the read end of the pipe (captured output)
	result, _ := io.ReadAll(r)

	// Convert the captured output to a string
	output := string(result)

	// Restore the original standard output
	os.Stdout = stdOut

	// Check if the captured output contains the expected string
	if !strings.Contains(output, "Hello World") {
		// If the expected string is not found, report a test failure
		t.Errorf("No Hello World in the output")
	}
}