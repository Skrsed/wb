BEGIN;

DROP TABLE IF EXISTS delivery;
CREATE TABLE delivery (
    id integer SERIAL PRIMARY KEY
    name varchar,
    phone varchar,
    zip varchar,
    city varchar,
    address varchar,
    region varchar,
    email varchar,
);

DROP TABLE IF EXISTS items;
CREATE TABLE items (
    id integer SERIAL PRIMARY KEY,
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
);

DROP TABLE IF EXISTS payments;
CREATE TABLE payments (
    id integer SERIAL PRIMARY KEY
    transaction varchar,
    request_id varchar,
    currency varchar,
    provider varchar,
    amount integer,
    payment_dt integer,
    bank varchar,
    delivery_cost integer,
    goods_total integer,
    custom_fee integer,
    order_uid varchar REFERENCES delivery,
);

DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
    order_uid varchar,
    track_number varchar,
    entry varchar,
    -- Why fk's should be here
    delivery_id integer REFERENCES delivery,
    payment_id integer REFERENCES payments,
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