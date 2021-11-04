package apiserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/fidesy/football-api/internal/model"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type APIServer struct {
	bind_addr string
	router    *mux.Router
	db        *sqlx.DB
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
		return err
	}
	defer s.db.Close()

	return http.ListenAndServe(s.bind_addr, s.router)
}

func (s *APIServer) configureStore(config *Config) error {
	db, err := sqlx.Open("postgres", config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
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
		w.Header().Set("Content-Type", "application/json")

		team := mux.Vars(r)["team"]
		amount, _ := strconv.ParseInt(mux.Vars(r)["amount"], 10, 64)

		matches := []model.Match{}
		err := s.db.Select(&matches, `select * from football where hometeam=$1 or  awayteam=$1`,
			team)
		if err != nil {
			w.WriteHeader(404)
			return
		}
		if len(matches) < int(amount) {
			json.NewEncoder(w).Encode(matches)
		} else {
			json.NewEncoder(w).Encode(matches[:amount])
		}

	}
}

func (s *APIServer) getVersusMatches() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		team_1 := mux.Vars(r)["team_1"]
		team_2 := mux.Vars(r)["team_2"]

		matches := []model.Match{}
		err := s.db.Select(&matches, `select * from football where hometeam in ($1, $2) 
			and awayteam in ($1, $2)`,
			team_1, team_2)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		json.NewEncoder(w).Encode(matches)
	}
}
