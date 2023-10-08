-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users (
    "vk_id" integer not null primary key,
    "state" varchar(255) not null,
    "colors" BOOLEAN DEFAULT false,
    "style_game" BOOLEAN DEFAULT false,
    "history" BOOLEAN DEFAULT false,
    "champions" BOOLEAN DEFAULT false,
    "local_fans" BOOLEAN DEFAULT false,
    "region" BOOLEAN DEFAULT false,
    "star_player" BOOLEAN DEFAULT false,
    "strong_def" BOOLEAN DEFAULT false,
    "fast_game" BOOLEAN DEFAULT false,
    "young_talent" BOOLEAN DEFAULT false,
    "strong_attack" BOOLEAN DEFAULT false,
    "foreign_players" BOOLEAN DEFAULT false,
    "technical_game" BOOLEAN DEFAULT false,
    "experienced_players" BOOLEAN DEFAULT false,
    "strong_character" BOOLEAN DEFAULT false,
    "young_coaches" BOOLEAN DEFAULT false,
    "team_play" BOOLEAN DEFAULT false,
    "nationwide_fans" BOOLEAN DEFAULT false,
    "strong_leadership" BOOLEAN DEFAULT false,
    "young_trainers" BOOLEAN DEFAULT false
);


-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;
