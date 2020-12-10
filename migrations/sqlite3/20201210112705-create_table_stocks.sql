
-- +migrate Up

create table if not exists stocks(
    store_id integer references stores(store_id),
    product_id integer references products(product_id),
    quantity integer not null,
    primary key (store_id, product_id)
);

-- +migrate Down

drop table if exists stocks;
