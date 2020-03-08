package main

import (
	"errors"
	"flag"
	"os"
	"sync"
)

func main() {
	flag.Parse()
	execute(flag.Args())
}

func execute(fileNames []string) {
	wg := sync.WaitGroup{}
	for _, name := range fileNames {
		wg.Add(1)
		go touch(name, &wg)
	}
	wg.Wait()
}

func touch(fileName string, wg *sync.WaitGroup) {
	defer wg.Done()

	if fileExists(fileName) {
		return
	}

	createFile(fileName)
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
