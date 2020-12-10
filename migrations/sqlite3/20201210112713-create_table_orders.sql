
-- +migrate Up

create table if not exists orders(
    order_id integer primary key autoincrement,
    costumer_id integer references customers(customer_id),
    order_status integer not null,
    order_date varchar(14),
    required_date varchar(14),
    shipped_date varchar(14),
    store_id integer references stores(store_id),
    staff_id integer references staffs(staff_id)
);

-- +migrate Down

drop table if exists orders;
