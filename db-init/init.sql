CREATE TABLE orders (
    order_uid VARCHAR(50) PRIMARY KEY,
    track_number VARCHAR(50) NOT NULL,
    entry VARCHAR(50),
    locale VARCHAR(10),
    internal_signature VARCHAR(255),
    customer_id VARCHAR(50),
    delivery_service VARCHAR(50),
    shardkey VARCHAR(10),
    sm_id INT,
    date_created TIMESTAMP,
    oof_shard VARCHAR(10)
);

CREATE TABLE deliveries (
    delivery_id SERIAL PRIMARY KEY,
    order_uid VARCHAR(50) REFERENCES orders(order_uid) ON DELETE CASCADE,
    name VARCHAR(100),
    phone VARCHAR(20),
    zip VARCHAR(20),
    city VARCHAR(50),
    address VARCHAR(255),
    region VARCHAR(50),
    email VARCHAR(100)
);

CREATE TABLE payments (
    payment_id SERIAL PRIMARY KEY,
    order_uid VARCHAR(50) REFERENCES orders(order_uid) ON DELETE CASCADE,
    transaction VARCHAR(50),
    request_id VARCHAR(50),
    currency VARCHAR(10),
    provider VARCHAR(50),
    amount DECIMAL(10, 2),
    payment_dt BIGINT,
    bank VARCHAR(50),
    delivery_cost DECIMAL(10, 2),
    goods_total DECIMAL(10, 2),
    custom_fee DECIMAL(10, 2)
);

CREATE TABLE items (
    item_id SERIAL PRIMARY KEY,
    order_uid VARCHAR(50) REFERENCES orders(order_uid) ON DELETE CASCADE,
    chrt_id INT,
    track_number VARCHAR(50),
    price DECIMAL(10, 2),
    rid VARCHAR(50),
    name VARCHAR(100),
    sale INT,
    item_size VARCHAR(10),
    total_price DECIMAL(10, 2),
    nm_id INT,
    brand VARCHAR(100),
    status INT
);
