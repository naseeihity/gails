-- TABLE FOR USERS -- 
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) DEFAULT '',
  `email` VARCHAR(50) DEFAULT '',
  `password` VARCHAR(100) DEFAULT '',
  `created_at` INT (10) unsigned DEFAULT '0',
  `updated_at` INT (10) unsigned DEFAULT '0',
  `deleted_at` INT (10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8;
INSERT INTO `users` (`id`, `name`, `email`, `password`)
VALUES ('1', 'admin', 'admin@gails.com', 'admin');
-- TABLE FOR ARTICLES --
CREATE TABLE `articles` (
  `id` INT(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(100) DEFAULT '',
  `desc` VARCHAR(255) DEFAULT '',
  `content` TEXT,
  `modified_by` VARCHAR(50) DEFAULT '',
  `created_at` INT (10) unsigned DEFAULT '0',
  `updated_at` INT (10) unsigned DEFAULT '0',
  `deleted_at` INT (10) unsigned DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;