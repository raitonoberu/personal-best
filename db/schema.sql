CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(255) primary key);
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    role_id INTEGER NOT NULL,
    email TEXT NOT NULL UNIQUE,
    password TEXT NOT NULL,

    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    middle_name TEXT NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (role_id) REFERENCES roles(id)
);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE roles (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,

    can_view BOOL NOT NULL,
    can_participate BOOL NOT NULL,
    can_create BOOL NOT NULL,
    is_free BOOL NOT NULL,
    is_admin BOOL NOT NULL
);
CREATE TABLE players (
    user_id INTEGER NOT NULL,
    birth_date DATE NOT NULL,
    is_male BOOL NOT NULL,
    phone TEXT NOT NULL,
    telegram TEXT NOT NULL, preparation TEXT NOT NULL DEFAULT '', position TEXT NOT NULL DEFAULT '',

    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE documents (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    player_id INTEGER NOT NULL,

    name TEXT NOT NULL,
    url TEXT NOT NULL,

    expires_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (player_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE competitions (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    trainer_id INTEGER NOT NULL,

    name TEXT NOT NULL,
    description TEXT NOT NULL,
    tours INTEGER NOT NULL,
    age INTEGER NOT NULL,
    size INTEGER NOT NULL,
    closes_at DATETIME NOT NULL,

    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (trainer_id) REFERENCES users(id) ON DELETE CASCADE
);
CREATE TABLE competition_days (
    competition_id INTEGER NOT NULL,
    date DATE NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,

    PRIMARY KEY (competition_id, date),
    FOREIGN KEY (competition_id) REFERENCES competitions(id) ON DELETE CASCADE
);
CREATE TABLE matches (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    competition_id INTEGER NOT NULL,
    start_time datetime NOT NULL,

    left_score INTEGER,
    right_score INTEGER,

    FOREIGN KEY (competition_id) REFERENCES competitions(id) ON DELETE CASCADE
);
CREATE VIEW user_players AS
    SELECT players.* FROM users LEFT JOIN players ON users.id = players.user_id
/* user_players(user_id,birth_date,is_male,phone,telegram,preparation,position) */;
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
CREATE TABLE match_players (
    match_id INTEGER NOT NULL,
    player_id INTEGER NOT NULL,
    position TEXT NOT NULL,
    team BOOL NOT NULL, win_score INTEGER, lose_score INTEGER,

    PRIMARY KEY (match_id, player_id),
    FOREIGN KEY (match_id) REFERENCES matches(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES users(id) ON DELETE CASCADE
);
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20240329185535'),
  ('20240503110458'),
  ('20240620194214'),
  ('20240630150631'),
  ('20240711164147');
