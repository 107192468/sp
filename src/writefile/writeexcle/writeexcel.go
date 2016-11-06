package writeexcel

import "github.com/tealeg/xlsx"

func CreateNewExcel(fileName string) (file *File, err error) {

	return xlsx.OpenFile(fileName)
}

func WriteNewExcelRow(fileName string, cells []string) {

}
