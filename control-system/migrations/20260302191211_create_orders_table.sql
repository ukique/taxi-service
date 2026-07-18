-- +goose Up
CREATE TYPE orderStatus AS ENUM(
    'created',
    'in_progress',
    'done'
);
CREATE TABLE orders
(
    id          BIGSERIAL PRIMARY KEY,
    driver_id   BIGINT,
    status      orderStatus DEFAULT 'created' NOT NULL,
    created_at  TIMESTAMP                     NOT NULL,
    finished_at TIMESTAMP,

    CONSTRAINT driver_id_fk FOREIGN KEY (driver_id) REFERENCES drivers (id)
);
-- +goose Down

DROP TABLE IF EXIST orders CASCADE;
DROP TYPE IF EXIST orderStatus;