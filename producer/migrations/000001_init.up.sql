DROP TABLE delivery IF EXISTS;
CREATE TABLE delivery (
    name varchar(),
    phone varchar(),
    zip varchar(),
    city varchar(),
    address varchar(),
    region varchar(),
    email varchar()
);

DROP TABLE IF EXISTS item;
CREATE TABLE item (
    chrt_id integer(),
    track_number varchar(),
    price varchar(),
    rid varchar(),
    name varchar(),
    sale integer(),
    size integer(),
    total_price integer(),
    nm_id integer(),
    brand varchar(),
    status integer()
)

