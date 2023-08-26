create table likes_dislikes
(
    id            bigint auto_increment primary key,
    user_id       bigint,
    resource_id   bigint,
    resource_type int     default 1 comment "1 = Post, 2 = Comment, 3 = Reply",
    action        tinyint default 1 comment "1 = Like, 0 = Dislike"
);