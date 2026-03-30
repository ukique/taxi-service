-- +goose Up
DROP TABLE drivers CASCADE;
-- +goose Down
CREATE TABLE drivers
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(24)             NOT NULL UNIQUE,
    password   TEXT                    NOT NULL,
    email      VARCHAR(255)            NOT NULL UNIQUE,
    status     BOOLEAN                 NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL
);