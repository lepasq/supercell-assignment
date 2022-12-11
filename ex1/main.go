package main

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func main() {
	inputFile := flag.String("i", "../tests/ex1/input1.txt", "file to run broadcasting with")
	flag.Parse()

	network := NewNetwork()
	c := make(chan string)
	go ReadFileLines(*inputFile, c)

	for i := range c {
		ExecuteCommand(i, network)
	}
}

// ReadFileLines reads an inputFile and sends the lines to a given channel
func ReadFileLines(path string, c chan string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	close(c)
}
