-- +migrate Up

CREATE TABLE IF NOT EXISTS `post_photos` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT(10) UNSIGNED NOT NULL,
    `url` TEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- +migrate Down

DROP TABLE IF EXISTS `post_photos`;