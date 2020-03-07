package main

import (
	"fmt"
	"os"
	"testing"
)

func TestFileExists(t *testing.T) {
	const fileName = "file"

	if exist := fileExists(fileName); exist == true {
		t.Fatalf("it should return false because %s doesn't exist", fileName)
	}

	// setup for test
	file, err := os.Create(fileName)
	if err != nil {
		t.Fatalf("test setup failed")
	}
	defer os.Remove(fileName)
	defer file.Close()

	if exist := fileExists(fileName); exist == false {
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
	ch := make(chan string)

	go touch(fileName, ch)
	defer os.Remove(fileName)
	if <-ch != fmt.Sprintf("%s is created\n", fileName) {
		t.Fatalf("made bug")
	}

	go touch(fileName, ch)
	if <-ch != fmt.Sprintf("%s already exist\n", fileName) {
		t.Fatalf("made bug")
	}
}
