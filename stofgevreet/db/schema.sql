CREATE SCHEMA `stofgevreet` ;

CREATE TABLE `stofgevreet`.`scan` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `car` VARCHAR(1000) NULL,
  `scantime` VARCHAR(1000) NULL,
  `method` VARCHAR(1000) NULL,
  `user` VARCHAR(1000) NULL,
  `points` VARCHAR(1000) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `stofgevreet`.`point` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `car` VARCHAR(1000) NULL,
  `scantime` VARCHAR(1000) NULL,
  `method` VARCHAR(1000) NULL,
  `user` VARCHAR(1000) NULL,
  `points` VARCHAR(1000) NULL,
  PRIMARY KEY (`id`));

CREATE TABLE `stofgevreet`.`stopwatch` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `car` VARCHAR(1000) NULL,
  `scantime` VARCHAR(1000) NULL,
  `method` VARCHAR(1000) NULL,
  `user` VARCHAR(1000) NULL,
  `lap` VARCHAR(1000) NULL,
  PRIMARY KEY (`id`));

USE `stofgevreet`;
DROP procedure IF EXISTS `save_scan`;

USE `stofgevreet`;
DROP procedure IF EXISTS `stofgevreet`.`save_scan`;
;

DELIMITER $$
USE `stofgevreet`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `save_scan`(IN car_param varchar(1000), IN scantime_param datetime, IN method_param varchar(1000), IN user_param varchar(1000))
BEGIN
	INSERT INTO `stofgevreet`.`scan`
	(
		`car`,
		`scantime`,
		`method`,
		`user`,
		`create_time`
    )
    values
    (
		car_param,
        scantime_param,
        method_param,
        user_param,
        current_timestamp()
    );
    
END$$

DELIMITER ;
;




USE `stofgevreet`;
DROP procedure IF EXISTS `save_point`;

USE `stofgevreet`;
DROP procedure IF EXISTS `stofgevreet`.`save_point`;
;

DELIMITER $$
USE `stofgevreet`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `save_point`(IN car_param varchar(1000), IN scantime_param varchar(1000), IN method_param varchar(1000), IN user_param varchar(1000), IN points_param int)
BEGIN
	INSERT INTO `stofgevreet`.`point`
	(
		`car`,
		`scantime`,
		`method`,
		`user`,
		`points`,
        `create_time`
    )
    values
    (
		car_param,
        scantime_param,
        method_param,
        user_param,
        points_param,
        current_timestamp()
    );
    
END$$

DELIMITER ;
;




USE `stofgevreet`;
DROP procedure IF EXISTS `save_stopwatch`;

USE `stofgevreet`;
DROP procedure IF EXISTS `stofgevreet`.`save_stopwatch`;
;

DELIMITER $$
USE `stofgevreet`$$
CREATE DEFINER=`root`@`localhost` PROCEDURE `save_stopwatch`(IN car_param varchar(1000), IN scantime_param datetime, IN method_param varchar(1000), IN user_param varchar(1000), IN lap_param varchar(1000))
BEGIN
	INSERT INTO `stofgevreet`.`scan`
	(
		`car`,
		`scantime`,
		`method`,
		`user`,
		`lap`,
        `create_time`
    )
    values
    (
		car_param,
        scantime_param,
        method_param,
        user_param,
        lap_param,
        current_timestamp()
    );
    
END$$

DELIMITER ;
;


  