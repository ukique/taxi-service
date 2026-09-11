-- +goose Up
CREATE INDEX idx_driver_locations_driver_id ON driver_locations (driver_id, id DESC);
CREATE INDEX idx_driver_locations_order_id_at ON driver_locations (order_id, id DESC);
CREATE INDEX idx_driver_locations_created_at ON driver_locations (created_at);
CREATE INDEX idx_driver_status ON drivers (status);
-- +goose Down
DROP INDEX idx_driver_locations_driver_id IF EXISTS;
DROP INDEX idx_driver_locations_order_id_at IF EXISTS;
DROP INDEX idx_driver_locations_created_at IF EXISTS;
DROP INDEX idx_driver_status IF EXISTS;
