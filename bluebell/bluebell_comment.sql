DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `comment_id` bigint(20) NOT NULL COMMENT '评论id',
    `post_id` bigint(20) NOT NULL COMMENT '帖子id',
    `author_id` bigint(20) NOT NULL COMMENT '评论者用户id',
    `content` varchar(2048) COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
    `parent_id` bigint(20) DEFAULT NULL COMMENT '父评论id，用于回复评论',
    `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '评论状态',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_comment_id` (`comment_id`),
    KEY `idx_post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
