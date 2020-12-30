drop database if exists iddb;
create database iddb;
use iddb;
drop table if exists publisher_3;
create table publisher_3(pid int primary key, pname char(100), nation char(3));
drop table if exists orders_3;
create table orders_3(ocid int, obid int, quantity int);

alter table publisher_3 add primary key(pid);
alter table orders_3 add primary key(ocid, obid, quantity);