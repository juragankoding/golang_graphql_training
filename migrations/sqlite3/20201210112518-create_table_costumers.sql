
-- +migrate Up

create table if not exists customers(
    customer_id integer primary key autoincrement,
    first_name varchar(20) not null,
    last_name varchar(20),
    phone varchar(20) not null,
    email varchar(20) not null,
    street varchar(20),
    city varchar (15),
    state varchar(15),
    zip_code integer
);

-- +migrate Down

drop table if exists customers;
