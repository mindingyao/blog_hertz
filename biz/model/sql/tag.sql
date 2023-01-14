CREATE TABLE `tag`
(
    `id`          bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'PK',
    `name`        varchar(128) NOT NULL DEFAULT '' COMMENT 'User name',
    `state`       tinyint(3) NOT NULL DEFAULT 1 COMMENT 'User gender',
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'User information create time',
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'User information update time',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'User information delete time',
    PRIMARY KEY (`id`),
    KEY           `idx_name` (`name`,`deleted_at`) COMMENT 'User name index'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='Tag information table'