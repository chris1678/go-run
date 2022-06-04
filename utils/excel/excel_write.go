package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
)

func Write(filename string) *excelize.File {
	f := excelize.NewFile()
	// 创建一个工作表
	//index := f.NewSheet("Sheet2")
	//// 设置单元格的值
	//f.SetCellValue("Sheet2", "A2", "Hello world.")
	//f.SetCellValue("Sheet1", "B2", 100)
	//// 设置工作簿的默认工作表
	//f.SetActiveSheet(index)
	// 根据指定路径保存文件
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}

	return f
}
