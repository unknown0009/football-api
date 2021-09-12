package apiserver

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/fidesy/football-api/internal/model"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type APIServer struct {
	bind_addr string
	router    *mux.Router
	db        *sql.DB
}

func New() *APIServer {
	return &APIServer{
		bind_addr: ":80",
		router:    mux.NewRouter(),
	}
}

func (s *APIServer) Start(config *Config) error {
	s.configureRouter()

	if err := s.configureStore(config); err != nil {
		log.Fatal(err)
	}

	return http.ListenAndServe(s.bind_addr, s.router)
}

func (s *APIServer) configureStore(config *Config) error {
	db, err := sql.Open("postgres", config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	s.db = db

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/api/versus/{team_1}/{team_2}", s.getVersusMatches())
	s.router.HandleFunc("/api/{team}/{amount}", s.getMatchesByTeamName())
}

func (s *APIServer) getMatchesByTeamName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		team := mux.Vars(r)["team"]
		amount, _ := strconv.ParseInt(mux.Vars(r)["amount"], 10, 64)
		rows, err := s.db.Query(`select * from football where home_team=$1 or away_team=$1`,
			team)
		if err != nil {
			fmt.Println(err)
		}

		matches := collectMatchesFromRows(rows, amount)
		json.NewEncoder(w).Encode(matches)
	}
}

func (s *APIServer) getVersusMatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		team_1 := mux.Vars(r)["team_1"]
		team_2 := mux.Vars(r)["team_2"]

		rows, err := s.db.Query(`select * from football where home_team in ($1, $2) 
			and away_team in ($1, $2)`,
			team_1, team_2)
		if err != nil {
			fmt.Println(err)
		}

		matches := collectMatchesFromRows(rows, 1000)
		json.NewEncoder(w).Encode(matches)
	}
}

func collectMatchesFromRows(rows *sql.Rows, amount int64) []model.Match {
	matches := []model.Match{}
	var count int64 = 0
	for rows.Next() {
		if count == amount {
			break
		}
		count++
		match := model.Match{}
		err := rows.Scan(&match.Id, &match.Date, &match.Tournament, &match.HomeTeam,
			&match.HomeGoals, &match.AwayGoals, &match.AwayTeam, &match.HalfHomeGoals,
			&match.HalfAwayGoals, &match.Total, &match.Url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		matches = append(matches, match)
	}
	return matches
}
