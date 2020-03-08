package main

import (
	"os"
	"sync"
	"testing"
)

func TestFileExists(t *testing.T) {
	const fileName = "file"

	if exist := fileExists(fileName); exist != false {
		t.Fatalf("it should return false because %s doesn't exist", fileName)
	}

	// setup for test
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("test setup failed")
	}
	defer os.Remove(fileName)
	defer file.Close()

	if exist := fileExists(fileName); exist != true {
		t.Fatalf("it should return true because %s exists", fileName)
	}
}

func TestCreateFile(t *testing.T) {
	const fileName = "file"
	err := createFile(fileName)
	defer os.Remove(fileName)

	if _, e := os.Stat(fileName); e != nil {
		t.Fatalf("it should create %s", fileName)
	}
	if err != nil {
		t.Fatalf("it should return nil becuase file is created now")
	}
}

func TestTouch(t *testing.T) {
	const fileName = "file"
	wg := sync.WaitGroup{}

	wg.Add(1)
	go touch(fileName, &wg)
	defer os.Remove(fileName)
	wg.Wait()
	if !fileExists(fileName) {
		t.Fatalf("it should create %s", fileName)
	}

	wg.Add(1)
	go touch(fileName, &wg)
	wg.Wait()
	if !fileExists(fileName) {
		t.Fatalf("it should remain %s", fileName)
	}
}

func TestExecute(t *testing.T) {
	fileNames := []string{"file1", "file2", "file3"}

	execute(fileNames)
	defer func() {
		for _, name := range fileNames {
			os.Remove(name)
		}
	}()
	for _, name := range fileNames {
		if exist := fileExists(name); exist != true {
			t.Fatalf("it should create %s", name)
		}
	}
}
