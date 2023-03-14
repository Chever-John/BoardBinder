package services

import (
	"github.com/xuri/excelize/v2"
	"strings"
)

func GetAllPeopleNames() (data []string) {
	f, err := excelize.OpenFile(ExcelPath)
	if err != nil {
		println(err.Error())
		return
	}

	rows, err := f.GetRows("2023 年 12 周活动_Free day of U")
	if err != nil {
		println(err.Error())
		return
	}

	for k, v := range rows {
		if k == 0 {
			continue
		}

		for i, value := range v {
			if i == 0 {
				namePrefix := strings.Split(value, "(")
				data = append(data, namePrefix[0])
			}
		}
	}
	return

}
