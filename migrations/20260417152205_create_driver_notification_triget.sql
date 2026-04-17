-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION notify_drivers()
RETURNS TRIGGER AS $$
BEGIN
    PERFORM pg_notify('drivers', 'updates drivers!');
RETURN NULL;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER drivers_notify
    AFTER UPDATE OR INSERT ON drivers
    FOR EACH STATEMENT
EXECUTE FUNCTION notify_drivers();
-- +goose Down
DROP TRIGGER drivers_notify ON drivers;
DROP FUNCTION notify_drivers();
