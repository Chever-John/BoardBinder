package models

type GamesAndWeight struct {
	GameName string
	Weight   int
}

type DayAndPeopleAndTheirFavoriteGames struct {
	Day            string
	GamesAndWeight []GamesAndWeight
	Peoples        []string
}
