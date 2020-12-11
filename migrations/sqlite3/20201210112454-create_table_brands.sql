
-- +migrate Up

create table if not exists brands(
    brand_id integer primary key autoincrement,
    brand_name varchar(40) not null
);

-- +migrate Down

drop table if exists brands;