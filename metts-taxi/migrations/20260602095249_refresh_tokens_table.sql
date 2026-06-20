-- +goose Up
CREATE TABLE refresh_tokens
(
    id            BIGSERIAL PRIMARY KEY,
    username      VARCHAR(16)  NOT NULL,
    refresh_token VARCHAR(512) NOT NULL,
    created_at    TIMESTAMP    NOT NULL,
    expires_at    TIMESTAMP    NOT NULL
);

-- +goose Down
DROP TABLE refresh_tokens;