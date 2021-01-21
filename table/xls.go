package table

import (
	"github.com/extrame/xls"
)

type Txls struct {
	Filename string
}

func (t Txls) Read()([][]string, error){
	rows := [][]string{}
	f, err := xls.Open(t.Filename, "utf-8")
	if err != nil {
		return rows, err
	}

	sheet := f.GetSheet(0)
	for j := 0; j < int(sheet.MaxRow)+1; j++ {
		xlsRow := sheet.Row(j)
		rowColCount := xlsRow.LastCol()
		r := []string{}
		for i := 0; i < rowColCount; i++ {
			r = append(r, xlsRow.Col(i))
		}
		rows = append(rows, r)
	}
	return rows, nil
}
