-- +goose Up
ALTER TABLE driver_locations
ADD created_at TIMESTAMP NOT NULL;
-- +goose Down
ALTER TABLE driver_locations
DROP COLUMN created_at;