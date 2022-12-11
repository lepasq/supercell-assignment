package main

import (
	"fmt"
	"strings"
	"sync"
	"testing"
)

func TestBroadcastSingleThreaded(t *testing.T) {
	testBroadcastFile("../tests/ex2/input1.txt", "../tests/ex2/output1.txt", 1, t)
}

func TestBroadcastFiveThreads(t *testing.T) {
	testBroadcastFile("../tests/ex2/input1.txt", "../tests/ex2/output1.txt", 5, t)
}

func TestBroadcastTenThreads(t *testing.T) {
	testBroadcastFile("../tests/ex2/input1.txt", "../tests/ex2/output1.txt", 10, t)
}

// testBroadcastFile runs the test cases one given input file
func testBroadcastFile(input, output string, routines int, t *testing.T) {
	fmt.Printf("Processing '%v' using %v routines...\n", input, routines)
	network := NewNetwork()
	c := make(chan string)
	go ReadFileLines(input, c)

	if routines == 1 {
		broadcastSingleThreaded(c, network)
	} else {
		broadcastMultiThreaded(c, network, 10)
	}

	message, _ := PrettyPrintBroadcast(network)
	assertBroadcastMatch(message, output, t)
}

// Runs commands from a channel on a network using one single go routine
func broadcastSingleThreaded(c chan string, network *Network) {
	for v := range c {
		ExecuteCommand(v, network)
	}
}

// Runs commands from a channel on a network using multiple go routines
func broadcastMultiThreaded(c chan string, network *Network, routines int) {
	wg := new(sync.WaitGroup)
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			for v := range c {
				go ExecuteCommand(v, network)
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

// assertBroadcastMatch checks that the actual output equals the expected output
func assertBroadcastMatch(actual string, output string, t *testing.T) {
	c := make(chan string)
	go ReadFileLines(output, c)

	lines := strings.Split(actual, "\n")

	i := 0
	for expected := range c {
		if lines[i] != expected {
			t.Errorf("\nexpected: \n%v \nreceived: \n%v \nat line %v", expected, lines[i], i+1)
		}
		i += 1
	}
}
