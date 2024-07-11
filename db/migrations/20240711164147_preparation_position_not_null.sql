-- migrate:up
ALTER TABLE players
    DROP COLUMN preparation;
ALTER TABLE players
    DROP COLUMN position;

ALTER TABLE players
    ADD COLUMN preparation TEXT NOT NULL DEFAULT '';
ALTER TABLE players
    ADD COLUMN position TEXT NOT NULL DEFAULT '';

CREATE TEMPORARY TABLE temp AS
SELECT
    *
FROM
    players;

DROP TABLE players;

CREATE TABLE players (
    user_id INTEGER NOT NULL,
    birth_date DATE NOT NULL,
    is_male BOOL NOT NULL,
    phone TEXT NOT NULL,
    telegram TEXT NOT NULL,
    preparation TEXT NOT NULL,
    position TEXT NOT NULL,

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO players (
    user_id,
    birth_date,
    is_male,
    phone,
    telegram,
    preparation,
    position
)
SELECT * FROM temp;

DROP TABLE temp;
-- migrate:down

