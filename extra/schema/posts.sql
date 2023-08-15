create table posts
(
    id             bigint auto_increment primary key,
    user_id        bigint,
    title          text,
    body           text,
    slug           text,
    original_image varchar(800)   default null,
    sm_image       varchar(800)   default null,
    md_image       varchar(800)   default null,
    lg_image       varchar(800)   default null,
    reported       tinyint        default 0,
    created_at     timestamp      default CURRENT_TIMESTAMP,
    updated_at     timestamp null default null,
    deleted_at     timestamp null default null
);