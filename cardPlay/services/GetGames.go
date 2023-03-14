package services

import (
	"github.com/xuri/excelize/v2"
	"strings"
)

const (
	//ExcelPath = "/Users/10027852/Work/loadbalancer/demos/getPodIPs/readExcel/ddd.xlsx"
	ExcelPath = "/Users/10027852/Happy/2023 年 12 周活动.xlsx"
)

func ReturnFavoriteGamesByPeopleName(peopleName string) (data []string) {
	f, err := excelize.OpenFile(ExcelPath)
	if err != nil {
		println(err.Error())
		return
	}

	rows, err := f.GetRows("2023 年 12 周活动_Favorite Games of")
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
				if namePrefix[0] == peopleName {
					vArray := strings.Split(v[1], ",")
					for _, v2 := range vArray {
						data = append(data, v2)
					}
				}
			}
		}
	}
	return
}
