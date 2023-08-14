create table tags
(
    id         bigint auto_increment primary key,
    tag        varchar(80),
    created_at timestamp      default CURRENT_TIMESTAMP,
    updated_at timestamp null default null,
    deleted_at timestamp null default null
);