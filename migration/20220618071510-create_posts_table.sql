-- +migrate Up

CREATE TABLE IF NOT EXISTS `posts` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `content` VARCHAR(256) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- +migrate Down

DROP TABLE IF EXISTS `posts`;