CREATE DATABASE IF NOT EXISTS gails;
USE gails;
-- TABLE FOR USERS -- 
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) DEFAULT '',
  `email` VARCHAR(50) DEFAULT '',
  `password` VARCHAR(100) DEFAULT '',
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8;
-- TABLE FOR ARTICLES --
CREATE TABLE IF NOT EXISTS `articles` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) DEFAULT '',
  `abstract` VARCHAR(255) DEFAULT '',
  `content` TEXT,
  `user_id` INT(10) unsigned,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;