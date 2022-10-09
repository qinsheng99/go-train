package test

import (
	"testing"

	"github.com/qinsheng99/goWeb/library/excel"
)

func TestExcel(t *testing.T) {
	err := excel.Create()
	if err != nil {
		t.Fatal(err)
	}
}

func TestExcelRead(t *testing.T) {
	err := excel.Read("demo.xlsx")
	if err != nil {
		t.Fatal(err)
	}
}
