package second

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func TestUpdateMessage(t *testing.T){
	var wg sync.WaitGroup
	wg.Add(1)
	updateMessage("Hello World", &wg)
	wg.Wait()
	if msg != "Hello World"{
		t.Errorf("The string doesn't changed to Hello World: %s", msg)
	}
}

func TestPrintMessage(t *testing.T){
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup
	wg.Add(1)	
	updateMessage("Test Message", &wg)
	wg.Wait()
	printMessage()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "Test Message"){
		t.Errorf("The stdOut doesn't contains new message")
	}
}

func TestRunChallenge(t *testing.T){
	stdOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	
	RunChallenge()

	w.Close()
	result, _ := io.ReadAll(r)
	output := string(result)
	os.Stdout = stdOut
	if !strings.Contains(output, "Hello, world!"){
		t.Errorf("The stdOut doesn't contains one of the strings: Hello, world!")
	}
	if !strings.Contains(output, "Hello, cosmos!"){
		t.Errorf("The stdOut doesn't contains one of the strings: Hello, cosmos!")
	}
	if !strings.Contains(output, "Hello, universe!"){
		t.Errorf("The stdOut doesn't contains one of the strings: Hello, universe!")
	}
}