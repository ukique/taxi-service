-- +goose Up
ALTER TABLE driver_locations
    ADD COLUMN generated_time TIMESTAMP NOT NULL;
-- +goose Down
ALTER TABLE driver_locations
    DROP COLUMN generated_time;