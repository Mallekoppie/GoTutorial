-- CREATE TABLE `car` (
--   `id` 			varchar(200) 			NOT NULL,
--   `team_name` 	varchar(1000) 	DEFAULT NULL,
--   `name` 		varchar(1000) 	DEFAULT NULL,
--   `driver_name` varchar(1000) 	DEFAULT NULL,
--   `make` 		varchar(1000) 	DEFAULT NULL,
--   `model` 		varchar(1000) 	DEFAULT NULL,
--   `year` 		int			 	DEFAULT NULL,
--   `engine_size`	int 			DEFAULT NULL,
--   `value` 		int 			DEFAULT NULL,
--   `create_time` timestamp NULL 	DEFAULT NULL,
--   `update_time` timestamp NULL 	DEFAULT NULL,
--   PRIMARY KEY (`id`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

DELIMITER $$
CREATE PROCEDURE `insert_car` 
(
	IN id_param varchar(200), 
	IN team_name_param varchar(1000), 
	IN name_param varchar(1000),
	IN driver_name_param varchar(1000),
	IN make_param varchar(1000),
	IN model_param varchar(1000),
    IN year_param int,
    IN engine_size_param int,
    IN value_param int
)
BEGIN

	INSERT INTO `stofgevreet`.`car`
	(
		`id`,
		`team_name`,
		`name`,
		`driver_name`,
		`make`,
		`model`,
        `year`,
        `engine_size`,
        `value`,
		`create_time`,
		`update_time`
	)
	values
	(
		id_param,
		team_name_param,
		name_param,
		driver_name_param,
		make_param,
		model_param,
        year_param,
        engine_size_param,
        value_param,
		current_timestamp(),
		current_timestamp()
	);
    
END$$
DELIMITER ;

DELIMITER $$
CREATE PROCEDURE `update_car` 
(
	IN id_param varchar(200), 
	IN team_name_param varchar(1000), 
	IN name_param varchar(1000),
	IN driver_name_param varchar(1000),
	IN make_param varchar(1000),
	IN model_param varchar(1000),
	IN year_param int,
    IN engine_size_param int,
    IN value_param int
)
BEGIN

	update `stofgevreet`.`car`
	set 
		`team_name` = 	team_name_param,
		`name` = name_param,
		`driver_name` = driver_name_param,
		`make` = make_param,
		`model` = model_param,		
        `year` = year_param,
        `engine_size` = engine_size_param,
        `value` = value_param,
		`update_time` = current_timestamp()	
    where
		`id` = id_param;
    
END$$
DELIMITER ;
