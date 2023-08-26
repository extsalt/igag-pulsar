create table comments
(
    id             bigint auto_increment primary key,
    user_id        bigint,
    resource_id    bigint,
    resource_type  int default 1 comment "1 = Post, 2 = Comment, 3 = Reply",
    body           text,
    like_count     bigint         default 0,
    dislike_count  bigint         default 0,
    reported       tinyint        default 0,
    original_image varchar(800)   default null,
    sm_image       varchar(800)   default null,
    md_image       varchar(800)   default null,
    lg_image       varchar(800)   default null,
    created_at     timestamp      default CURRENT_TIMESTAMP,
    updated_at     timestamp null default null,
    deleted_at     timestamp null default null
);