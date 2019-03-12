package excel

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
)

func ReadExcel(src string) {
	xlsx, err := excelize.OpenFile(src)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//获取指定单元格数据
	cell := xlsx.GetCellValue("Sheet1", "B2")
	println(cell)

	// 获取所有行数据
	rows := xlsx.GetRows("Sheet1")
	var data [][]string
	for _, row := range rows {
		var rowData []string
		for _, col := range row {
			rowData = append(rowData, col)
		}
		data = append(data, rowData)
	}
	println(data[1][0])
}

func WriteExcel() {
	xlsx := excelize.NewFile()
	// 创建工作表
	sheet := xlsx.NewSheet("sheet1")
	// 设置指定cell的值
	xlsx.SetCellValue("sheet1", "A2", "test")
	xlsx.SetActiveSheet(sheet)
	// 保存文件
	err := xlsx.SaveAs("test.xlsx")
	if err != nil {
		fmt.Println(err)
	}
}
