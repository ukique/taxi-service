-- +goose Up
CREATE TYPE driverStatus AS ENUM (
    'driving',
    'searching',
    'offline'
);
CREATE TABLE drivers
(
    id         BIGSERIAL PRIMARY KEY,
    username   VARCHAR(24)                    NOT NULL UNIQUE,
    status     driverStatus DEFAULT 'offline' NOT NULL,
    created_at TIMESTAMP    DEFAULT NOW()     NOT NULL

);
-- +goose Down
DROP TABLE drivers CASCADE;
DROP TYPE driverStatus;