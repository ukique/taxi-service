-- +goose Up
CREATE TABLE driver_locations
(
    id         BIGSERIAL PRIMARY KEY,
    order_id   BIGINT,
    driver_id  BIGINT                  NOT NULL,
    lat        DOUBLE PRECISION        NOT NULL,
    lon        DOUBLE PRECISION        NOT NULL,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL,

    CONSTRAINT order_id_fk FOREIGN KEY (order_id) REFERENCES orders (id),
    CONSTRAINT driver_id_fk FOREIGN KEY (driver_id) REFERENCES drivers (id)
);
-- +goose Down
DROP TABLE driver_locations;