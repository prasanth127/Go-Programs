package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestCreateWriteFile(t *testing.T) {
	fn := "sample.txt"
	data := "prasanth"
	fileName := CreateWriteFile(fn, data)
	_, e := os.Stat(fileName)
	if e != nil {
		// Checking if the given file exists or not
		// Using IsNotExist() function
		if os.IsNotExist(e) {
			t.Errorf("File not Found !!")
		}
	}

	databyte, _ := ioutil.ReadFile(fileName)

	if string(databyte) != data {
		if os.IsNotExist(e) {
			t.Errorf("Data in the file not matched!!")
		}
	}
}
