CREATE TABLE "order"(
    id	CHAR(36) PRIMARY KEY,
    product_id CHAR(36),
    quantity INTEGER NOT NULL,
    user_name VARCHAR(30) NOT NULL,
    user_address VARCHAR(30) NOT NULL,
    user_phone VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
