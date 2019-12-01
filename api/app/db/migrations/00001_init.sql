-- +goose Up
-- SQL in this section is executed when the migration is applied.

SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS user_passwords;
SET FOREIGN_KEY_CHECKS=1;

CREATE TABLE users(
    id VARCHAR(128) NOT NULL PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP);

CREATE TABLE user_passwords(
    user_id VARCHAR(128) NOT NULL PRIMARY KEY,
    password_hash VARCHAR(255),

    CONSTRAINT fk_user_passwords_user_id FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
