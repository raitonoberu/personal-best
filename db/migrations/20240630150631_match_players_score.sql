-- migrate:up
ALTER TABLE match_players ADD COLUMN win_score INTEGER;
ALTER TABLE match_players ADD COLUMN lose_score INTEGER;

-- migrate:down
ALTER TABLE match_players DROP COLUMN win_score;
ALTER TABLE match_players DROP COLUMN lose_score;
