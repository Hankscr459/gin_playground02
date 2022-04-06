package excel

import (
	export "ginValid/dto/export"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// type lkExcelExport struct {
// 	file      *excelize.File
// 	sheetName string //可定義預設 sheet 名稱
// }

func Employee(c *gin.Context) {
	employee := &export.Employee{
		Id:         "1",
		Name:       "Hank",
		Department: "Manifacture",
	}
	f := excelize.NewFile()
	index := f.NewSheet("Sheet1")
	list := []export.Employee{
		{Id: "2", Name: "Andrew", Department: "Service"},
	}
	list = append(list, *employee)
	index = 1
	for _, e := range list {
		values := map[string]interface{}{
			"A" + strconv.Itoa(index): e.Id, "B" + strconv.Itoa(index): e.Department, "C" + strconv.Itoa(index): e.Name}
		for k, v := range values {
			f.SetCellValue("Sheet1", k, v)
		}
		index += 1
	}

	f.SetActiveSheet(index)
	buf, _ := f.WriteToBuffer()
	c.Header("Content-Type", "application/vnd.ms-excel;charset=utf8")
	c.Header("Content-Disposition", "attachment; filename=test.xlsx")
	_, _ = c.Writer.Write(buf.Bytes())
}
