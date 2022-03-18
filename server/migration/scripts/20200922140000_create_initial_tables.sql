-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS users
(
    id           UUID                      NOT NULL,
    username     VARCHAR(255)              NOT NULL,
    password     VARCHAR(255)              NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT uix_users_username UNIQUE (username)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS users;