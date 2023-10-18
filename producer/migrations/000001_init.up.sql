BEGIN;

DROP TABLE IF EXISTS delivery;
CREATE TABLE delivery (
    name varchar,
    phone varchar,
    zip varchar,
    city varchar,
    address varchar,
    region varchar,
    email varchar,
    order_uid varchar
);

DROP TABLE IF EXISTS items;
CREATE TABLE items (
    chrt_id integer,
    track_number varchar,
    price varchar,
    rid varchar,
    name varchar,
    sale integer,
    size integer,
    total_price integer,
    nm_id integer,
    brand varchar,
    status integer,
    order_uid varchar
);

DROP TABLE IF EXISTS payments;
CREATE TABLE payments (
    name varchar,
    phone varchar,
    zip varchar,
    city varchar,
    address varchar,
    region varchar,
    email varchar,
    order_uid varchar
);

DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
    order_uid varchar,
    track_number varchar,
    entry varchar,
    delivery_id integer,
    locale varchar,
    internal_signature varchar,
    customer_id varchar,
    delivery_service varchar,
    shardkey integer,
    sm_id integer,
    date_created timestamp,
    oof_shard integer
);

COMMIT;