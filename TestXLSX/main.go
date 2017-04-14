package main

import (
	"fmt"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

func main() {
	// var file *xlsx.File
	// var sheet *xlsx.Sheet
	// var row *xlsx.Row
	// var cell *xlsx.Cell
	// var err error

	// file = xlsx.NewFile()
	// sheet, err = file.AddSheet("Sheet1")
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }
	// row = sheet.AddRow()
	// cell = row.AddCell()
	// cell.Value = "I am a cell!"
	// err = file.Save("MyXLSXFile.xls")
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// }

	// read XLSX
	// Read()

	// read XLS
	ReadXLS()
}

func ReadXLS() {
	//Output: read the content of first two cols in each row
	if xlFile, err := xls.Open("工作.xls", "utf-8"); err == nil {
		if sheet1 := xlFile.GetSheet(0); sheet1 != nil {
			fmt.Print("Total Lines ", sheet1.MaxRow, sheet1.Name)
			col1 := sheet1.Row(0).Col(0)
			col2 := sheet1.Row(0).Col(0)
			for i := 0; i <= (int(sheet1.MaxRow)); i++ {
				row1 := sheet1.Row(i)
				col1 = row1.Col(0)
				col2 = row1.Col(1)
				fmt.Print("\n", col1, ",", col2)
			}
		}
	}
}

// cannot read .xls suffix
func Read() {
	excelFileName := "工作.xls"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("read err: ", err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text, _ := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
}
