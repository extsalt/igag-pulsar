create table personal_access_tokens
(
    id            bigint auto_increment primary key,
    access_token  varchar(200),
    identity_type varchar(80),
    identity_id   varchar(80),
    abilities     json,
    expires_in    bigint comment '0 means never expires',
    created_at    timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp null default null,
    deleted_at timestamp null default null
);