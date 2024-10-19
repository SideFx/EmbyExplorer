// ---------------------------------------------------------------------------------------------------------------------
// (w) 2024 by Jan Buchholz
// XLSX export, using Excelite by xuri
// https://github.com/qax-os/excelize
// ---------------------------------------------------------------------------------------------------------------------

package export

import (
	"github.com/xuri/excelize/v2"
)

type Payload struct {
	XLSCell string
	Data    string
}

type HeaderData struct {
	XLSCell string
	Name    string
	Column  string
	Width   float64
}

func XlsxExport(data []Payload, header []HeaderData, path string, sheet string) error {
	var err error
	f := excelize.NewFile()
	var headerFont = excelize.Font{
		Bold:      true,
		Italic:    false,
		Underline: "single",
		Size:      14,
		Color:     "000000", //black
	}
	var standardFont = excelize.Font{
		Bold:   false,
		Italic: false,
		Size:   12,
		Color:  "5A5A5A", //dark grey
	}
	var headerColor = []string{"D7DBDD", "D7DBDD"}
	var headerFill = excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   headerColor,
		Shading: 0,
	}
	var standardColor = []string{"F8F9F9", "F8F9F9"}
	var standardFill = excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   standardColor,
		Shading: 0,
	}
	var headerStyleId, _ = f.NewStyle(&excelize.Style{Font: &headerFont, Fill: headerFill})
	var standardStyleId, _ = f.NewStyle(&excelize.Style{Font: &standardFont, Fill: standardFill})
	defer func() {
		if e := f.Close(); e != nil {
			err = e
		}
	}()
	index, err := f.NewSheet(sheet)
	if err != nil {
		return err
	}
	// Set header
	for _, h := range header {
		err = f.SetCellStr(sheet, h.XLSCell, h.Name)
		if err != nil {
			return err
		}
		err = f.SetCellStyle(sheet, h.XLSCell, h.XLSCell, headerStyleId)
		if err != nil {
			return err
		}
		err = f.SetColWidth(sheet, h.Column, h.Column, h.Width)
		if err != nil {
			return err
		}
	}
	// Set data
	for _, d := range data {
		err = f.SetCellStr(sheet, d.XLSCell, d.Data)
		if err != nil {
			return err
		}
		err = f.SetCellStyle(sheet, d.XLSCell, d.XLSCell, standardStyleId)
		if err != nil {
			return err
		}
	}
	f.SetActiveSheet(index)
	err = f.SaveAs(path)
	return err
}
