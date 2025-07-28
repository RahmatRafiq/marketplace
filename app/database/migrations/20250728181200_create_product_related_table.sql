-- +++ UP Migration
CREATE TABLE product_related (
    product_id BIGINT NOT NULL,
    related_product_id BIGINT NOT NULL,
    PRIMARY KEY (product_id, related_product_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE,
    FOREIGN KEY (related_product_id) REFERENCES products(id) ON DELETE CASCADE
);
-- --- DOWN Migration
DROP TABLE IF EXISTS product_related;
