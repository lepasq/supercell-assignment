package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	inputFile := flag.String("i", "../tests/ex2/input1.txt", "file to run broadcasting with")
	routines := flag.Int("r", 10, "number of routines to run concurrently")
	flag.Parse()

	network := NewNetwork()
	c := make(chan string)
	go ReadFileLines(*inputFile, c)

	wg := new(sync.WaitGroup)
	wg.Add(*routines)
	for i := 0; i < *routines; i++ {
		go func() {
			for v := range c {
				go ExecuteCommand(v, network)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	pretty, err := PrettyPrintBroadcast(network)
	if err != nil {
		log.Print(err)
	}
	fmt.Println(pretty)
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
		log.Fatal(err)
	}
	close(c)
}
