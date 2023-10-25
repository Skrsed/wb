BEGIN;

CREATE TYPE ECURRENCY AS ENUM (
 'USD', 
 'RUB', 
 'EUR',
 'JPY'
);

CREATE TYPE ELOCALE AS ENUM (
    'en',
    'ru',
    'kz',
    'by'
);

CREATE TABLE orders (
    order_uid VARCHAR PRIMARY KEY,
    track_number VARCHAR,
    entry VARCHAR NOT NULL,
    locale ELOCALE NOT NULL,
    internal_signature VARCHAR,
    customer_id VARCHAR NOT NULL,
    delivery_service VARCHAR,
    shardkey VARCHAR NOT NULL,
    sm_id INTEGER NOT NULL CHECK (sm_id > 0),
    date_created TIMESTAMP NOT NULL,
    oof_shard VARCHAR NOT NULL
);

CREATE TABLE delivery (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    phone VARCHAR NOT NULL,
    zip SMALLINT NOT NULL,
    city VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    region VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    order_uid VARCHAR REFERENCES orders (order_uid)
);

CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    transaction VARCHAR NOT NULL,
    request_id VARCHAR,
    currency ECURRENCY NOT NULL,
    provider VARCHAR NOT NULL,
    amount NUMERIC NOT NULL CHECK (amount > 0),
    payment_dt TIMESTAMP NOT NULL,
    bank VARCHAR NOT NULL,
    delivery_cost NUMERIC NOT NULL CHECK (delivery_cost > 0),
    goods_total NUMERIC NOT NULL CHECK (goods_total > 0),
    custom_fee NUMERIC NOT NULL CHECK (custom_fee > 0),
    order_uid VARCHAR REFERENCES orders (order_uid)
);

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    chrt_id BIGINT NOT NULL CHECK (chrt_id > 0),
    price NUMERIC NOT NULL,
    rid VARCHAR NOT NULL,
    name VARCHAR NOT NULL,
    sale NUMERIC NOT NULL,
    size VARCHAR NOT NULL,
    total_price NUMERIC NOT NULL,
    nm_id BIGINT NOT NULL CHECK (nm_id > 0),
    brand VARCHAR NOT NULL,
    status SMALLINT NOT NULL,
    order_uid VARCHAR REFERENCES orders
);

COMMIT;