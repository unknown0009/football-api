CREATE TABLE football (
        id serial PRIMARY KEY,
        date date,
        tournament varchar(100) NOT NULL,
        home_team varchar(50) NOT NULL,
        home_goals INTEGER,
        away_goals INTEGER,
        away_team varchar(50) NOT NULL,
        half_home INTEGER,
        half_away INTEGER,
        total INTEGER,
        url varchar(100) NOT NULL
);
COPY football FROM './football.csv' DELIMITER ',' CSV HEADER;