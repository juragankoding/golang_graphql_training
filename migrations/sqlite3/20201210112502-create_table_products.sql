
-- +migrate Up

create table if not exists products (
    product_id integer primary key autoincrement,
    product_name varchar(40) not null,
    brand_id integer references brands(brand_id),
    category_id integer references categories(category_id),
    model_year varchar(10),
    list_price integer not null
);

-- +migrate Down

drop table if exists products;
