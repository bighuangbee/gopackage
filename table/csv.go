package table

import (
	"encoding/csv"
	"io"
	"os"
)

type Tcsv struct {
	Filename string
}

func (t Tcsv)Read()([][]string, error){
	rows := [][]string{}

	file, err := os.Open(t.Filename)
	if err != nil {
		return rows, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			break
		}
		if err == io.EOF {
			break
		}
		rows = append(rows, row)
	}
	return rows, nil
}