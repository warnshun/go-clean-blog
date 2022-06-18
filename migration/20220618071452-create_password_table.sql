-- +migrate Up

CREATE TABLE IF NOT EXISTS `password` (
    `user_id` INT(10) UNSIGNED NOT NULL,
    `password` VARCHAR(256) NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `deleted_at` DATETIME,
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

-- +migrate Down

DROP TABLE IF EXISTS `password`;