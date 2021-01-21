package main

import (
	"fmt"
	"os"
	"table"
	"testing"
)

func TestReadcsv(t *testing.T) {
	p,_ := os.Getwd()
	filePath :=  p + "/房地产.csv"

	e,err := table.Engine(filePath)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	data, err := e.Read()
	if err == nil {
		for k, val := range data{
			fmt.Println(k, val)
		}
	}
}

func TestReadxls(t *testing.T) {
	p,_ := os.Getwd()
	filePath :=  p + "/房地产.xls"

	e,err := table.Engine(filePath)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	data, err := e.Read()
	if err == nil {
		for k, val := range data{
			fmt.Println(k, val)
		}
	}
}

func TestReadxlsx(t *testing.T) {
	p,_ := os.Getwd()
	filePath :=  p + "/房地产.xlsx"

	e,err := table.Engine(filePath)
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	data, err := e.Read()
	if err == nil {
		for k, val := range data{
			fmt.Println(k, val)
		}
	}
}
