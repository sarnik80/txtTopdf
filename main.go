package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Usage  :  file1.txt   file2.txt  .. [filen.txt ... ]")
	}

	fileNames := arguments[1:]

	var resultStr string

	for _, textfile := range fileNames {

		text, err := ioutil.ReadFile(textfile)

		if err != nil {

			fmt.Println("An error occurred while reading the file " + textfile)

			continue
		}

		resultStr += string(text)

	}

}
