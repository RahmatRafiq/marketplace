-- +++ UP Migration
CREATE TABLE products (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    reference VARCHAR(255) NOT NULL,
    product_base_id BIGINT,
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) NOT NULL,
    brand VARCHAR(255),
    short_description TEXT,
    long_description TEXT,
    weight DECIMAL(10,2),
    dimension_1 DECIMAL(10,2),
    dimension_2 DECIMAL(10,2),
    dimension_3 DECIMAL(10,2),
    koli INT,
    sku VARCHAR(255),
    lowest_retail_price DECIMAL(15,2),
    branch_prices JSON,
    stock INT,
    images JSON,
    received_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    FOREIGN KEY (product_base_id) REFERENCES product_bases(id) ON DELETE SET NULL
);

-- --- DOWN Migration
DROP TABLE IF EXISTS products;
