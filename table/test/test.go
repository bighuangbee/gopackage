package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"table"
)

func main(){

	p,_ := os.Getwd()


	fmt.Println(strings.TrimRight(filepath.Base(p + "/房地产.xlsx"), filepath.Ext(p + "/房地产.xlsx")))


	filePath :=  p + "/房地产.csv"
	//filePath :=  p + "/房地产.xls"

	t, err := table.Engine(filePath)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	data, err := t.Read()
	fmt.Println(err, data)

	//t := table.Txlsx{}
	//data, err := t.Read(path + "/房地产.xlsx")
	//fmt.Println(err, data)
	//
	//txls := table.Txls{}
	//data1, err1 := txls.Read(path + "/房地产.xls")
	//fmt.Println(err1, data1)
}
