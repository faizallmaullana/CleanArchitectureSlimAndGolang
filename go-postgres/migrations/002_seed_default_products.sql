DROP PROCEDURE IF EXISTS seed_default_products();

CREATE OR REPLACE PROCEDURE seed_default_products()
LANGUAGE plpgsql
AS $procedure$
BEGIN
    INSERT INTO products (id, name, price)
    VALUES
        (1, 'Product 1', 10.99),
        (2, 'Product 2', 19.99),
        (3, 'Product 3', 5.99)
    ON CONFLICT (id) DO UPDATE
    SET
        name = EXCLUDED.name,
        price = EXCLUDED.price,
        updated_at = NOW();

    PERFORM setval(
        pg_get_serial_sequence('products', 'id'),
        GREATEST((SELECT COALESCE(MAX(id), 1) FROM products), 1),
        true
    );
END;
$procedure$;

CALL seed_default_products();
