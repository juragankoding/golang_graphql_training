
-- +migrate Up
create table jenis_barang (
    id integer primary key AUTOINCREMENT,
    jenis varchar(25) not null
);

-- +migrate Down
drop table if exists jenis_barang;
