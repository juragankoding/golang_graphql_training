
-- +migrate Up

create table if not exists categories(
    category_id INTEGER primary key autoincrement,
    category_name varchar(40) not null
);

-- +migrate Down

drop table if exists categories;