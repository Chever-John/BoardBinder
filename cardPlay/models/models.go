package models

// PeopleAndHisDays data
type PeopleAndHisDays struct {
	PeopleName string
	HisDays    []string
}

type AllPeopleAndTheirDays struct {
	PeopleAndTheirDays []PeopleAndHisDays
}

// PeopleAndHisFavoriteGames data2
type PeopleAndHisFavoriteGames struct {
	PeopleName       string
	HisFavoriteGames []string
}

type AllPeopleAndTheirFavoriteGames struct {
	PeopleAndTheirFavoriteGames []PeopleAndHisFavoriteGames
}

// ThatDaysGamesAndWeight data3
type ThatDaysGamesAndWeight struct {
	GameName string
	Weight   int
}

type DaysAndThatDaysGamesAndPeople struct {
	Day                    string
	ThatDaysGamesAndWeight []ThatDaysGamesAndWeight
	ThatDaysPeople         []string
}

type AllDaysAndThatDaysGamesAndPeople struct {
	DaysAndThatDaysGamesAndPeople []DaysAndThatDaysGamesAndPeople
}

// Final 最终出现的数据结构
type Final struct {
	PeopleName    string
	FavoriteGames []string
	FreeDays      []string
}
