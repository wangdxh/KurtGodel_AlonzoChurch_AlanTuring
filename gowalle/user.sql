DROP DATABASE IF EXISTS systest;##
CREATE DATABASE IF NOT EXISTS systest default character set utf8 collate utf8_general_ci;##

use systest;##

create table sys_user(
    id bigint not null auto_increment,
    user_name varchar(50),
    user_password varchar(128),
    user_email varchar(50),
    user_info text ,
    head_img blob,
    create_time datetime,
    primary key(id)
);##

create table sys_role(
    id bigint not null auto_increment comment 'role id',
    role_name varchar(50),
    enabled int,
    create_by bigint,
    create_time datetime ,
    primary key(id)
);##

create table sys_privilege(
    id bigint not null auto_increment ,
    privilege_name varchar(50),
    privilege_url varchar(200),
    primary key(id)
);##

create table sys_user_role(
    user_id bigint,
    role_id bigint
);##

create table sys_role_privilege(
    role_id bigint,
    privilege_id bigint
);##

insert into sys_user values ('1', 'admin', '123456', 'admin@mybatis.tk', 'administrator', null, '2019-04-03 18:00:00');##
insert into sys_user values ('1001', 'test', '123456', 'test@mybatis.tk', 'testor', null, '2019-04-04 18:00:00');##
insert into sys_role values ('1', 'administrator', '1', '1', '2019-04-03 18:01:00');##
insert into sys_role values ('2', 'normalor', '1', '1', '2019-04-03 18:03:00');##
insert into sys_user_role values('1','1');##
insert into sys_user_role values('1','2');##
insert into sys_user_role values('1001','1');##

insert into sys_privilege value ('1', "user mangeer", '/users');##
insert into sys_privilege value ('2', "role  mangeer", '/roles');##
insert into sys_privilege value ('3', "logs mangeer", '/logs');##
insert into sys_privilege value ('4', "person mangeer", '/persons');##
insert into sys_privilege value ('5', "comapnies mangeer", '/companies');##

insert into sys_role_privilege values ('1', '1');##
insert into sys_role_privilege values ('1', '3');##
insert into sys_role_privilege values ('1', '2');##
insert into sys_role_privilege values ('2', '4');##
insert into sys_role_privilege values ('2', '5');##

show engines;