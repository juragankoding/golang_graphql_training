
-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
    id INTEGER primary key autoincrement ,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    display_name VARCHAR(30) NOT NULL
);

-- +migrate Down

DROP TABLE users;