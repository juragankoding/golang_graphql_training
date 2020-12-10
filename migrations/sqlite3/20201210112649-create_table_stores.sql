
-- +migrate Up

create table IF NOT EXISTS stores(
    store_id integer primary key autoincrement,
    store_name varchar(40) not null,
    phone varchar(10) not null,
    email varchar(20) not null,
    street varchar(10),
    city varchar(20),
    state varchar(20),
    zip_code varchar(10)
);

-- +migrate Down

drop table if exists stores;
