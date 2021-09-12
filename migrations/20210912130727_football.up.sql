CREATE TABLE football (
        id serial PRIMARY KEY,
        date date,
        tournament varchar(100) NOT NULL,
        homeTeam varchar(50) NOT NULL,
        homeGoals INTEGER,
        awayGoals INTEGER,
        awayTeam varchar(50) NOT NULL,
        halfHome INTEGER,
        halfAway INTEGER,
        total INTEGER,
        url varchar(100) NOT NULL
);