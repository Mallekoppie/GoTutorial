CREATE SCHEMA `test` ;

CREATE TABLE `test`.`user` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(200) NULL,
  `surname` VARCHAR(20) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `id_UNIQUE` (`id` ASC) VISIBLE);

USE `test`;
DROP procedure IF EXISTS `insert_user`;

DELIMITER $$
USE `test`$$
CREATE PROCEDURE `insert_user` (IN name_param varchar(200), IN surname_param varchar(200))
BEGIN
insert into user(name, surname)
values
(
name_param,
surname_param
);
END$$

DELIMITER ;

