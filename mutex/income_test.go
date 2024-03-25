package mutex

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestRunIncome(t *testing.T){
	// Save the original standard output (stdout)
	stdOut := os.Stdout

	// Create a new pipe to capture the output
	r, w, _ := os.Pipe()

	// Redirect the standard output to the write end of the pipe
	os.Stdout = w

	RunIncome()
	w.Close()
	// Read all data from the read end of the pipe (captured output)
	result, _ := io.ReadAll(r)

	// Convert the captured output to a string
	output := string(result)

	// Restore the original standard output
	os.Stdout = stdOut

	// Check if the captured output contains the expected string
	if !strings.Contains(output, "82160") {
		// If the expected string is not found, report a test failure
		t.Errorf("Incorrect bank balance in the output")
	}
}