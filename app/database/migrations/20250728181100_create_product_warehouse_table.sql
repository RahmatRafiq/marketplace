-- +++ UP Migration
CREATE TABLE product_warehouses (
    product_id BIGINT NOT NULL,
    warehouse_id BIGINT NOT NULL,
    PRIMARY KEY (product_id, warehouse_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
    -- FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE
);
-- --- DOWN Migration
DROP TABLE IF EXISTS product_warehouses;
