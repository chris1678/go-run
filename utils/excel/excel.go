package excel

/*
* 同时支持xls和xlxs格式
* 2022-5-23
 */
import (
	"bytes"
	"errors"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/huangtao-sh/xls"
	"io"
	"io/ioutil"
	"path/filepath"
)

type ex interface {
	read(filename string) error
	reader(reader io.Reader) error
	GetSheetsName() map[int]string
	GetRow(sheetName string) ([][]string, error)
	GetRowIndex(i int) ([][]string, error)
}

type xlsT struct {
	file *xls.WorkBook
}

func (e *xlsT) read(filename string) error {

	file, err2 := xls.Open(filename, "utf-8")

	if err2 != nil {
		return err2
	}
	e.file = file
	return nil
}
func (e *xlsT) reader(reader io.Reader) error {
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}
	read := bytes.NewReader(data)
	openReader, err2 := xls.OpenReader(read, "utf-8")
	if err2 != nil {
		return err2
	}
	e.file = openReader
	return nil
}
func (e *xlsT) GetSheetsName() map[int]string {
	m := make(map[int]string)
	for i := 0; i < e.file.NumSheets(); i++ {
		m[i] = e.file.GetSheet(i).Name
	}
	return m
}

func (e *xlsT) GetRow(sheetName string) ([][]string, error) {
	rows, err := e.file.GetRows(sheetName)
	if err != nil {
		return rows, err
	}
	return rows, nil
}
func (e *xlsT) GetRowIndex(i int) ([][]string, error) {
	rows, err := e.file.GetRowsIndex(i)
	if err != nil {
		return rows, err
	}
	return rows, nil
}

type xlsTx struct {
	file *excelize.File
}

func (e *xlsTx) read(filename string) error {
	var err error
	e.file, err = excelize.OpenFile(filename)
	return err
}
func (e *xlsTx) reader(reader io.Reader) error {
	openReader, err := excelize.OpenReader(reader)
	if err != nil {
		return err
	}
	e.file = openReader
	return nil
}
func (e *xlsTx) GetSheetsName() map[int]string {
	return e.file.GetSheetMap()
}

func (e *xlsTx) GetRow(sheetName string) ([][]string, error) {
	rows := e.file.GetRows(sheetName)
	return rows, nil
}

// GetRowIndex 索引从1开始的
func (e *xlsTx) GetRowIndex(i int) ([][]string, error) {
	i++
	a := e.GetSheetsName()
	value, ok := a[i]
	if ok == false {
		value = a[1]
	}
	return e.GetRow(value)

}

// Import 文件格式倒入
func Import(filename string) (ex, error) {
	suffix := filepath.Ext(filename)
	if suffix == "" {
		return nil, errors.New("文件类型错误")
	}
	if suffix == ".xls" {
		file := &xlsT{}
		err := file.read(filename)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		file := &xlsTx{}
		err := file.read(filename)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
}

// ImportByReader io.reader方式导入
func ImportByReader(reader io.Reader, fileType string) (ex, error) {

	if fileType == "xls" {
		file := &xlsT{}
		err := file.reader(reader)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		file := &xlsTx{}
		err := file.reader(reader)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
}
