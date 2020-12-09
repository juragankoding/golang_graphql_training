
-- +migrate Up
create table barang (
    id integer primary key AUTOINCREMENT,
    nama varchar(50) not null,
    desc varchar(100),
    jumlah integer,
    jenis_barang integer
);

-- +migrate Down
drop table if exists barang;