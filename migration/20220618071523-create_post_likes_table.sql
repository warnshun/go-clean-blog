-- +migrate Up

CREATE TABLE IF NOT EXISTS `post_likes` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT(10) UNSIGNED NOT NULL,
    `user_id` INT(10) UNSIGNED NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- +migrate Down

DROP TABLE IF EXISTS `post_likes`;