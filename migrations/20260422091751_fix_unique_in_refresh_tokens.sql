-- +goose Up
ALTER TABLE refresh_tokens DROP CONSTRAINT refresh_tokens_username_key;

-- +goose Down
ALTER TABLE refresh_tokens ADD CONSTRAINT refresh_tokens_username_key UNIQUE (username);