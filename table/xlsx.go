package table

import (
	"github.com/tealeg/xlsx"
)

type Txlsx struct {
	Filename string
}


func (t Txlsx)Read()([][]string, error){
	rows := [][]string{}

	file, err := xlsx.OpenFile(t.Filename)
	if err != nil{
		return rows, err
	}

	sheet1 := file.Sheets[0]
	for k := 0; k < sheet1.MaxRow; k++{
		rowXlsx := sheet1.Row(k)
		r := []string{}
		for i := 0; i < len(rowXlsx.Cells); i++{
			r = append(r, rowXlsx.Cells[i].Value)
		}
		rows = append(rows, r)
	}

	return rows, nil
}