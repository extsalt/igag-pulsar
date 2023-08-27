create table users
(
    id             bigint auto_increment primary key,
    username       varchar(80),
    email          varchar(255),
    password       varchar(255),
    avatar         varchar(800)   default null,
    active         tinyint        default 1,
    oauth_provider varchar(80)    default null,
    created_at     timestamp      default CURRENT_TIMESTAMP,
    updated_at     timestamp null default null,
    deleted_at     timestamp null default null,
    index (username),
    index (email)
);