# football-api

## REST API for getting statistics of football matches
 ---
How this provides statistics such as:
* Date
* Tournament
* Team names
* Goals / goals in 1st half
* Link to extended statistics

--- 

Database contains over 80k matches. Matches of teams from the 5 leagues:
* England (Premier League)
* Spain (LaLiga)
* Germany (Bundesliga)
* Italy (Serie A)
* France (Ligue 1)

[List of all teams](./teams.csv)

---

## Usage
```golang
func main() {
    req, err := http.Get(
        "http://92.255.77.155/api/Chelsea/10")
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}

	var matches []Match
	json.Unmarshal(body, &matches)
}
```
[Match struct](./internal/model/match.go) for parsing json

Parsing date is 29.08.2021, data update later...

Examples of requests:
* /api/Chelsea/10 - get the last 10 matches of Chelsea; /api/{team_name}/{amount_of_last_matches}

* /api/versus/Chelsea/Liverpool - get all matches between Chelsea and Liverpool;   
/api/versus/{team_name_1}/{team_name_2}


---

## Adding hostname
I haven't come up with a hostname yet, but even this IP address looks pretty good.