-- +++ UP Migration
CREATE TABLE product_tags (
    product_id BIGINT NOT NULL,
    tag_id BIGINT NOT NULL,
    PRIMARY KEY (product_id, tag_id),
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
    -- FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);
-- --- DOWN Migration
DROP TABLE IF EXISTS product_tags;
