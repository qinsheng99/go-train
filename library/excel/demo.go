package excel

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func Create() error {
	f := excelize.NewFile()

	s := "sheet"
	index := f.NewSheet(s)

	sheetTileMap := make(map[string]string)
	sheetTileMap["A1"] = "A"
	sheetTileMap["B1"] = "B"
	sheetTileMap["C1"] = "C"

	for k, v := range sheetTileMap {
		_ = f.SetCellValue(s, k, v)
	}

	for k, v := range []int{1, 2, 3} {
		_ = f.SetCellInt(s, "A"+strconv.Itoa(k+2), v)
		_ = f.SetCellInt(s, "B"+strconv.Itoa(k+2), v+1)
		_ = f.SetCellInt(s, "C"+strconv.Itoa(k+2), v+2)
	}

	f.SetActiveSheet(index)
	if err := f.SaveAs("demo.xlsx"); err != nil {
		return err
	}

	return nil

}

func Read(f string) error {
	file, err := excelize.OpenFile(f)
	if err != nil {
		return err
	}
	if rows, rerr := file.GetRows("sheet"); rerr == nil {
		for _, row := range rows {
			fmt.Println(row)
		}
	}

	return nil
}
