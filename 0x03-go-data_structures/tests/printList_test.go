package assignments

import (
	"bytes"
	golangDsa "dsa/assignments"

	"fmt"
	"io"
	"os"
	"testing"
)

func TestPrintListInteger(t *testing.T) {
	// Prepare the list of integers to test
	testList := []int{1, 2, 3, 4, 5}

	// Create a pipe to capture standard output
	r, w, _ := os.Pipe()
	// Save the original os.Stdout
	oldStdout := os.Stdout
	// Set os.Stdout to the write end of the pipe
	os.Stdout = w

	// Run the function to generate output
	golangDsa.PrintListInteger(testList)

	// Close the writer and restore os.Stdout
	w.Close()
	os.Stdout = oldStdout

	// Read the output from the read end of the pipe
	var buf bytes.Buffer
	io.Copy(&buf, r)

	// Define the expected output
	expectedOutput := "1\n2\n3\n4\n5\n"

	// Verify the output
	if buf.String() != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, buf.String())
	}
	fmt.Println("All tests passed")
}
