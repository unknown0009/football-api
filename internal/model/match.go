package model

type Match struct {
	Id         int
	Date       string
	Tournament string
	HomeTeam   string
	HomeGoals  int
	AwayGoals  int
	AwayTeam   string
	HalfHome   int
	HalfAway   int
	Total      int
	Url        string
}
