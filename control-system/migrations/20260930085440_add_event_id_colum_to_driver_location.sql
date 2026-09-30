-- +goose Up
ALTER TABLE driver_locations
    ADD event_id BIGINT NOT NULL DEFAULT 0;

-- +goose Down
ALTER TABLE driver_locations
DROP
COLUMN event_id;