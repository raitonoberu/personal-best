-- migrate:up
CREATE TEMPORARY TABLE temp AS
SELECT
    *
FROM
    registrations;

DROP TABLE registrations;

CREATE TABLE registrations (
    competition_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    is_approved BOOL NOT NULL,
    is_dropped BOOL NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (competition_id, player_id),
    FOREIGN KEY (competition_id) REFERENCES competitions(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO registrations (
    competition_id,
    player_id,
    is_approved,
    is_dropped,
    created_at
)
SELECT
    *
FROM
    temp;

DROP TABLE temp;


CREATE TEMPORARY TABLE temp AS
SELECT
    *
FROM
    match_players;

DROP TABLE match_players;

CREATE TABLE match_players (
    match_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    position TEXT NOT NULL,
    team BOOL NOT NULL,

    PRIMARY KEY (match_id, player_id),
    FOREIGN KEY (match_id) REFERENCES matches(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO match_players (
    match_id,
    player_id,
    position,
    team
)
SELECT
    *
FROM
    temp;

DROP TABLE temp;
-- migrate:down

