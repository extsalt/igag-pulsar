create table blocks
(
    id            bigint auto_increment primary key,
    user_id       bigint,
    resource_type varchar(255),
    resource_id   bigint,
    created_at    timestamp      default CURRENT_TIMESTAMP,
    updated_at    timestamp null default null,
    deleted_at    timestamp null default null
);