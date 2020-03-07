package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	for _, name := range flag.Args() {
		result := make(chan string)
		go touch(name, result)
		fmt.Print(<-result)
	}
}

func touch(fileName string, ch chan<- string) {
	if fileExists(fileName) {
		ch <- fmt.Sprintf("%s already exist\n", fileName)
		return
	}

	if err := createFile(fileName); err == nil {
		ch <- fmt.Sprintf("%s is created\n", fileName)
		return
	}

	ch <- fmt.Sprintf("Failed to create %s", fileName)
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		return false
	}
	return true
}

func createFile(fileName string) error {
	if file, err := os.Create(fileName); err == nil {
		file.Close()
		return nil
	}
	return errors.New("")
}
