package main

import (
	"time"

	"github.com/xuri/excelize/v2"
)

func main() {
	const sheet = "Sheet1"

	xlsx := excelize.NewFile()
	xlsx.SetCellValue(sheet, "A1", "Server Name")
	xlsx.SetCellValue(sheet, "B1", "Generation")
	xlsx.SetCellValue(sheet, "C1", "Acquisition Date")
	xlsx.SetCellValue(sheet, "D1", "CPU Vendor")

	xlsx.SetCellValue(sheet, "A2", "svlaa01")
	xlsx.SetCellValue(sheet, "B2", "12")
	xlsx.SetCellValue(sheet, "C2", mustParse("10/27/2021"))
	xlsx.SetCellValue(sheet, "D2", "Intel")

	xlsx.SetCellValue(sheet, "A3", "svlac14")
	xlsx.SetCellValue(sheet, "B3", "13")
	xlsx.SetCellValue(sheet, "C3", mustParse("12/13/2021"))
	xlsx.SetCellValue(sheet, "D3", "AMD")

	if err := xlsx.SaveAs("./Book1.xlsx"); err != nil {
		panic(err)
	}
}

func mustParse(s string) time.Time {
	const layout = "01/02/2006"

	time, err := time.Parse(layout, s)
	if err != nil {
		panic(err)
	}
	return time
}
