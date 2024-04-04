-- migrate:up
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,
    is_trainer BOOLEAN NOT NULL,
    birth_date DATETIME
);

CREATE TABLE competitions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    start_date DATETIME NOT NULL,
    trainer_id INTEGER NOT NULL,
    FOREIGN KEY (trainer_id) REFERENCES users(id) ON DELETE CASCADE
);

-- migrate:down
DROP TABLE users;

DROP TABLE competitions;