package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jung-kurt/gofpdf"
)

func main() {

	arguments := os.Args

	if len(arguments) == 2 {
		fmt.Println("Usage  : Title  file1.txt   file2.txt  .. [filen.txt ... ]")
	}

	title := arguments[1]

	fileNames := arguments[2:]

	var resultStr string

	for _, textfile := range fileNames {

		text, err := ioutil.ReadFile(textfile)

		if err != nil {

			fmt.Println("An error occurred while reading the file " + textfile)

			continue
		}

		resultStr += string(text)

	}

	// first argument is page orientation we can use P for portrait and L for landscape
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 14)
	width, _ := pdf.GetPageSize()
	pdf.MoveTo(width/2, 10)
	pdf.Cell(1, 1, title)

	pdf.MoveTo(0, 20)
	pdf.SetFont("Arial", "", 14)

	pdf.MultiCell(width, 10, resultStr, "", "", false)

	err := pdf.OutputFileAndClose("../../Desktop/" + title + ".pdf")

	if err == nil {
		fmt.Println("PDF generated successfully")
	}

}
