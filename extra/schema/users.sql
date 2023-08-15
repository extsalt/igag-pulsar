create table users
(
    id         bigint auto_increment primary key,
    username   varchar(80),
    email      varchar(255),
    password   varchar(255),
    avatar     varchar(800)   default null,
    active     tinyint        default 1,
    created_at timestamp      default CURRENT_TIMESTAMP,
    updated_at timestamp null default null,
    deleted_at timestamp null default null
);