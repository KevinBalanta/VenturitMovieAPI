SET FOREIGN_KEY_CHECKS=0;

DROP TABLE IF EXISTS `movies`;
CREATE TABLE `movies`(
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(255) NOT NULL,
    `rating` decimal(2,1) NOT NULL,
    `released_year` smallint NOT NULL,
    `genres` varchar(655) NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;