-- +goose Up
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(24)  NOT NULL UNIQUE,
    password   TEXT         NOT NULL,
    email      VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP
);

-- +goose Down
DROP TABLE users;