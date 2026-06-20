-- +goose Up
CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(16)  NOT NULL,
    password   TEXT         NOT NULL,
    email      VARCHAR(255) NOT NULL,
    created_at TIMESTAMP    NOT NULL
);

-- +goose Down
DROP TABLE users;
