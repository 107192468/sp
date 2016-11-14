package readexcel

import (
	"github.com/tealeg/xlsx"
)

func ReadExcelRow(filename string, index int) []*xlsx.Row {

	xlFile, err := xlsx.OpenFile(filename)

	if err != nil {
		panic(err)
	}
	sheet := xlFile.Sheets[index]
	return sheet.Rows

}

func ReadExcelCell(cells []*xlsx.Cell, n int) string {
	v := cells[n].Value
	return v
}

func ReadExcelCells(cells []*xlsx.Cell, n int) []string {
	var vs []string
	for i, v := range cells {
		vs[i] = v.Value
	}
	return vs
}
