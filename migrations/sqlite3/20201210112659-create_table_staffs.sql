
-- +migrate Up

create table if not exists staffs(
    staff_id integer primary key autoincrement,
    first_name varchar(40) not null,
    last_name varchar(40),
    email varchar(20) not null,
    phone varchar(10),
    active integer,
    store_id integer references stores(store_id),
    manager_id integer
);

-- +migrate Down

drop table if exists staffs;