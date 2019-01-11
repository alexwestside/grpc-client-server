package main

import (
	"fmt"
	"io"
	"os"
	"encoding/csv"
)

func reader() {
	file, err := os.Open(*File)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	var record []string
	for {
		record, err = reader.Read()
		if err == io.EOF {
			close(ReadChann)
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		ReadChann <- record
	}
}
