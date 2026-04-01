-- +goose Up
CREATE TYPE orderStatus AS ENUM(
    'created',
    'in_progress',
    'done'
);
CREATE TABLE orders
(
    id          BIGSERIAL PRIMARY KEY,
    user_id     BIGINT                        NOT NULL,
    driver_id   BIGINT,
    pickup_lat  DOUBLE PRECISION              NOT NULL,
    pickup_lon  DOUBLE PRECISION              NOT NULL,
    dropout_lat DOUBLE PRECISION              NOT NULL,
    dropout_lon DOUBLE PRECISION              NOT NULL,
    status      orderStatus DEFAULT 'created' NOT NULL,
    created_at  TIMESTAMP   DEFAULT NOW()     NOT NULL,
    updated_at  TIMESTAMP   DEFAULT NOW()     NOT NULL,

    CONSTRAINT user_id_fk FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT driver_id_fk FOREIGN KEY (driver_id) REFERENCES drivers (id)
);
-- +goose Down

DROP TABLE IF EXIST orders CASCADE;
DROP TYPE IF EXIST orderStatus;