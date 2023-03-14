package main

import (
	"fmt"
	"github.com/Chever-John/BoardBinder/models"
	"github.com/Chever-John/BoardBinder/services"
	"github.com/xuri/excelize/v2"
	"strings"
)

const (
	//ExcelPath = "/Users/10027852/Work/loadbalancer/demos/getPodIPs/readExcel/ddd.xlsx"
	ExcelPath = "/Users/10027852/Happy/2023 年 12 周活动.xlsx"
)

func main() {
	AllPeopleName := services.GetAllPeopleNames()

	var dataFinal []models.Final
	dataFinal = make([]models.Final, len(AllPeopleName))

	for i, peopleName := range AllPeopleName {
		dataFinal[i].PeopleName = peopleName
		dataFinal[i].FavoriteGames = services.ReturnFavoriteGamesByPeopleName(peopleName)
		dataFinal[i].FreeDays = services.ReturnFreeDaysByPeopleName(peopleName)
	}

	// 输出结果
	//for _, v := range dataFinal {
	//	fmt.Printf("PeopleName: %s\n", v.PeopleName)
	//	fmt.Printf("FavoriteGames: %s\n", v.FavoriteGames)
	//	fmt.Printf("FreeDays: %s\n", v.FreeDays)
	//}

	// 将结果存储入 []DayAndPeopleAndTheirFavoriteGames 中
	var data []models.DayAndPeopleAndTheirFavoriteGames
	data = make([]models.DayAndPeopleAndTheirFavoriteGames, 5)

	// 将周一到周五依次初始化进 data.Day 中
	for i := 0; i < 5; i++ {
		switch today := i; today {
		case 0:
			data[i].Day = "周一"
		case 1:
			data[i].Day = "周二"
		case 2:
			data[i].Day = "周三"
		case 3:
			data[i].Day = "周四"
		case 4:
			data[i].Day = "周五"

		}
	}

	// 将每个人的名字和他喜欢的游戏依次初始化进 data.PeopleAndTheirFavoriteGames 中
	for _, v := range dataFinal {
		fmt.Printf("开始处理 %s 的数据。\n", v.PeopleName)

		for _, freeDay := range v.FreeDays {
			for i, dataOfToday := range data {
				if freeDay == dataOfToday.Day {
					fmt.Printf("%s 的空闲时间与 %s 相同，将这个人添加到数据中。\n", v.PeopleName, dataOfToday.Day)

					// 如果这个人的空闲时间和 data 中的某一天相同，那么就将这个人的名字添加到 data.Peoples 中
					fmt.Printf("在添加 %s 之前，%s这天 data.Peoples 的值为 %s\n", v.PeopleName, ReturnDayByID(i), dataOfToday.Peoples)

					// 将这个人喜欢的游戏添加到 data.GamesAndWeight 中
					if IsStringInStringArray(v.PeopleName, data[i].Peoples) {
						fmt.Printf("%s 已经在 %s 这天的 data.Peoples 中了，不需要再添加了。\n", v.PeopleName, ReturnDayByID(i))
					} else {
						fmt.Printf("%s 不在 %s 这天的 data.Peoples (%s)中，需要添加。\n", v.PeopleName, ReturnDayByID(i), data[i].Peoples)
						data[i].Peoples = append(data[i].Peoples, v.PeopleName)
						fmt.Printf("在添加 %s 之后，%s这天 data.Peoples 的值为 %s\n", v.PeopleName, ReturnDayByID(i), dataOfToday.Peoples)
					}

					for _, favoriteGame := range v.FavoriteGames {
						// 如果这个 dataOfToday.GamesAndWeight 为空，那么就将这个人喜欢的游戏添加到 data.GamesAndWeight 中
						if len(dataOfToday.GamesAndWeight) == 0 {
							fmt.Printf("%s 这天的 data.GamesAndWeight 为空，需要添加。\n", ReturnDayByID(i))
							data[i].GamesAndWeight = append(data[i].GamesAndWeight, models.GamesAndWeight{
								GameName: favoriteGame,
								Weight:   1,
							})
							continue
						} else {
							// 判断这个游戏是否已经在 data.GamesAndWeight 中，如果已经存在，那么将其权重加一，如果不存在，那么将其添加进去
							for i2, _ := range dataOfToday.GamesAndWeight {
								if favoriteGame == data[i].GamesAndWeight[i2].GameName {
									fmt.Printf("%s 已经在 %s 这天的 data.GamesAndWeight 中了，将其权重加一。\n", favoriteGame, ReturnDayByID(i))
									data[i].GamesAndWeight[i2].Weight++
									break
								} else {
									fmt.Printf("%s 不在 %s 这天的 data.GamesAndWeight 中，需要添加。\n", favoriteGame, ReturnDayByID(i))
									data[i].GamesAndWeight = append(data[i].GamesAndWeight, models.GamesAndWeight{
										GameName: favoriteGame,
										Weight:   1,
									})
									break
								}
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("=====================================")
	// 输出每一天的游玩人数和游戏
	for _, v := range data {
		fmt.Printf("Day: %s\n", v.Day)
		fmt.Printf("People: %s\n", v.Peoples)
		for _, v2 := range v.GamesAndWeight {
			if v2.Weight <= 2 {
				continue
			}
			fmt.Printf("Game: %s, Weight: %d\n", v2.GameName, v2.Weight)
		}
	}
}

func ReturnDayByID(id int) (day string) {
	switch today := id; today {
	case 0:
		day = "周一"
	case 1:
		day = "周二"
	case 2:
		day = "周三"
	case 3:
		day = "周四"
	case 4:
		day = "周五"
	}
	return
}

// IsStringInStringArray string 是否在 []string 中
func IsStringInStringArray(str string, strArray []string) (isIn bool) {
	isIn = false
	for _, v := range strArray {
		if str == v {
			isIn = true
			break
		}
	}
	return
}

// ReturnAllPeopleAndTheirFavoriteGames 首先创建一个函数一次读取 excel 表格 sheet2 中的所有数据，并将其存储在 AllPeopleAndTheirFavoriteGames 中
func ReturnAllPeopleAndTheirFavoriteGames(ExcelPath string) (data models.AllPeopleAndTheirFavoriteGames) {
	f, err := excelize.OpenFile(ExcelPath)
	if err != nil {
		println(err.Error())
		return
	}

	rows, err := f.GetRows("我和我喜欢的桌游")
	if err != nil {
		println(err.Error())
		return
	}

	for k, v := range rows {
		if k == 0 {
			continue
		}

		var peopleAndHisFavoriteGames models.PeopleAndHisFavoriteGames

		for i, value := range v {
			if i == 0 {
				strPrefix := strings.Split(value, "(")
				peopleAndHisFavoriteGames.PeopleName = strPrefix[0]
			}
			if i == 1 {
				vArray := strings.Split(value, ",")
				for _, v2 := range vArray {
					peopleAndHisFavoriteGames.HisFavoriteGames = append(peopleAndHisFavoriteGames.HisFavoriteGames, v2)
				}

			}
		}
		data.PeopleAndTheirFavoriteGames = append(data.PeopleAndTheirFavoriteGames, peopleAndHisFavoriteGames)
	}
	return
}

// ReturnAllPeopleAndTheirDays 首先创建一个函数一次读取 excel 表格 sheet1 中的所有数据，并将其存储在 AllPeopleAndTheirDays 中
func ReturnAllPeopleAndTheirDays(ExcelPath string) (data models.AllPeopleAndTheirDays) {
	f, err := excelize.OpenFile(ExcelPath)
	if err != nil {
		println(err.Error())
		return
	}

	rows, err := f.GetRows("我和我有空的时间")
	if err != nil {
		println(err.Error())
		return
	}

	for k, v := range rows {
		if k == 0 {
			continue
		}

		var peopleAndHisDays models.PeopleAndHisDays

		for i, value := range v {
			if i == 0 {
				// 将 Value 中 "(" 符号前的内容存入 PeopleAndHisDays.PeopleName 中
				strPrefix := strings.Split(value, "(")
				peopleAndHisDays.PeopleName = strPrefix[0]
			}

			if i == 1 {
				// 将 v 按照逗号拆分成一个数组，然后存入 PeopleAndHisDays.HisDays 中
				vArray := strings.Split(value, ",")
				for _, v2 := range vArray {
					peopleAndHisDays.HisDays = append(peopleAndHisDays.HisDays, v2)
				}
			}
		}

		data.PeopleAndTheirDays = append(data.PeopleAndTheirDays, peopleAndHisDays)
	}

	println("--------------------------------------------------")

	return
}
