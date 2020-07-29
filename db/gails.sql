-- TABLE FOR USERS -- 
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) DEFAULT '',
  `email` VARCHAR(50) DEFAULT '',
  `password` VARCHAR(100) DEFAULT '',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8;
INSERT INTO `users` (`id`, `name`, `email`, `password`)
VALUES ('1', 'admin', 'admin@gails.com', 'admin');
-- TABLE FOR ARTICLES --
DROP TABLE IF EXISTS `articles`;
CREATE TABLE `articles` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) DEFAULT '',
  `abstract` VARCHAR(255) DEFAULT '',
  `content` TEXT,
  `user_id` INT(10) unsigned,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` TIMESTAMP DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;