package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func TestBroadcastBasic(t *testing.T) {
	testBroadcastFile("../tests/ex1/input1.txt", "../tests/ex1/output1.txt", t)
}

func TestBroadcastEmpty(t *testing.T) {
	testBroadcastFile("../tests/ex1/input2.txt", "../tests/ex1/output2.txt", t)
}

func TestBroadcastUnmarkFriends(t *testing.T) {
	testBroadcastFile("../tests/ex1/input3.txt", "../tests/ex1/output3.txt", t)
}

func TestBroadcastNewEntries(t *testing.T) {
	testBroadcastFile("../tests/ex1/input4.txt", "../tests/ex1/output4.txt", t)
}

func TestBroadcastNewEntriesFriends(t *testing.T) {
	testBroadcastFile("../tests/ex1/input5.txt", "../tests/ex1/output5.txt", t)
}

// testBroadcastFile runs the test cases one given input file
// Since the tested method prints the udpate directly to stdout without returning the value,
// it is needed to compare the expected output directly against stdout.
func testBroadcastFile(input, output string, t *testing.T) {
	fmt.Printf("Processing '%v'... \n", input)
	// Backup stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	runCommands(input)

	stdout := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		stdout <- buf.String()
	}()

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout

	out := <-stdout
	assertBroadcastMatch(out, output, t)
}

// runCommands executes all the commands from an input file inside a network
func runCommands(input string) {
	network := NewNetwork()
	c := make(chan string)
	go ReadFileLines(input, c)

	for i := range c {
		ExecuteCommand(i, network)
	}
}

// assertBroadcastMatch checks that the actual broacast output equals the expected output
func assertBroadcastMatch(out string, output string, t *testing.T) {
	c := make(chan string)
	go ReadFileLines(output, c)

	lines := strings.Split(out, "\n")
	i := 0
	for expected := range c {
		if lines[i] != expected {
			t.Errorf("\nexpected: \n%v \nreceived: \n%v \nat line %v", expected, lines[i], i+1)
		}
		i += 1
	}
}
