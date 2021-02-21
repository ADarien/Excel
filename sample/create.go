package main

import (
	"fmt"
	"math/rand"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		fmt.Println(err)
	}
	if err := streamWriter.SetRow("A1", []interface{}{
		excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
		fmt.Println(err)
	}
	for rowID := 2; rowID <= 10; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(440000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
