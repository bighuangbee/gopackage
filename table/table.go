package table

import (
	"errors"
	"path"
)

type ITable interface {
	Read()([][]string, error)
	//Write(path string)error
}


func Engine(filename string)(ITable, error){
	ext := path.Ext(filename)
	var t ITable
	if ext == ".xlsx"{
		t = Txlsx{Filename: filename}
	}else if ext == ".xls"{
		t = Txls{Filename: filename}
	}else if ext == ".csv"{
		t = Tcsv{Filename: filename}
	}else {
		return nil, errors.New("ext error")
	}
	return t, nil
}