create table personal_access_tokens
(
    id            integer auto_increment primary key,
    access_token  varchar(200),
    identity_type varchar(80),
    identity_id   varchar(80),
    abilities     varchar(80),
    expires_in    integer comment '0 means never expires',
    created_at    datetime default CURRENT_TIMESTAMP
);