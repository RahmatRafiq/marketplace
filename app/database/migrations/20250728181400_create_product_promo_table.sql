-- +++ UP Migration
CREATE TABLE product_promos (
    product_id BIGINT NOT NULL,
    promo_id BIGINT NOT NULL,
    PRIMARY KEY (product_id, promo_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
    -- FOREIGN KEY (promo_id) REFERENCES promos(id) ON DELETE CASCADE
);
-- --- DOWN Migration
DROP TABLE IF EXISTS product_promos;
