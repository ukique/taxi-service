-- +goose Up

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION notify_orders()
RETURNS TRIGGER AS $$
BEGIN
    PERFORM pg_notify('orders', 'updates order!');
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
   
CREATE TRIGGER orders_notify
    AFTER UPDATE OR INSERT ON orders
    FOR EACH STATEMENT
    EXECUTE FUNCTION notify_orders();
-- +goose Down
DROP TRIGGER orders_notify ON orders;
DROP FUNCTION notify_orders();
