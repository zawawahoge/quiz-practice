-- +goose Up
-- SQL in this section is executed when the migration is applied.
DROP TABLE IF EXISTS users;

CREATE TABLE users(
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    email VARCHAR(256) NOT NULL
    );

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE users;