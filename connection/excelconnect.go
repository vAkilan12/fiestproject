package connection

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

// Reading data from the Excel Sheet
func ExcelConnect(FileName, Sheet_no string, column_no int) ([]string, *excelize.File) {
	File, err := excelize.OpenFile(FileName)
	if err != nil {
		fmt.Println(err)
	}
	defer File.Close()
	readColumns, err := File.GetCols(Sheet_no)
	if err != nil {
		log.Fatal(err)
	}

	return readColumns[column_no], File
}

// printing the Excel Sheet
// func ExcelPrint(c string, b string) {
// 	_, S := ExcelConnect("Text Match.xlsx", "Sheet1", 0, 2)

// 	Sheet_no := "Sheet1"
// 	S.SetCellValue(Sheet_no, "A1", "DESCRIPTION")
// 	S.SetCellValue(Sheet_no, "B1", "            ")
// 	S.SetCellValue(Sheet_no, c, b)

// 	S.SetCellValue(Sheet_no, "C1", "Noun Found")
// 	S.SetCellValue(Sheet_no, "D1", "Type Found")
// 	S.SetCellValue(Sheet_no, "I1", "Material")
// 	S.SetCellValue(Sheet_no, "J1", "Pr Rating")
// 	S.SetCellValue(Sheet_no, "E1", "End Size1")
// 	S.SetCellValue(Sheet_no, "F1", "End Size2")
// 	S.SetCellValue(Sheet_no, "G1", "End type1")
// 	S.SetCellValue(Sheet_no, "H1", "End type2")
// 	//S.SetCellValue(Sheet_no, "J1", "Angle")

// 	if err := S.SaveAs("Text Match.xlsx"); err != nil {
// 		log.Println(err)
// 	}

// }
