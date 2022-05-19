package main

import (
	"os"
)

func CreateWriteFile(fileName, data string) string {
	file, _ := os.Create(fileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	file.WriteString(data)
	// WriteData(file, "data")

	return file.Name()
}
