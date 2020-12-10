
-- +migrate Up

create table IF NOT EXISTS order_items(
    item_id integer primary key autoincrement,
    order_id integer references orders(order_id),
    product_id integer references procucts(product_id),
    quantity integer not null,
    list_price integer not null,
    discount integer not null
);

-- +migrate Down

drop table if exists order_items;
