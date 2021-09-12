package model

type Match struct {
	Id            int
	Date          string
	Tournament    string
	HomeTeam      string
	HomeGoals     int
	AwayGoals     int
	AwayTeam      string
	HalfHomeGoals int
	HalfAwayGoals int
	Total         int
	Url           string
}