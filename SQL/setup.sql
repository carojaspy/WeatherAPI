
use weatherapi;

-- drop table `task`
--     DROP TABLE IF EXISTS `task`
-- --------------------------------------------------
--  Table Structure for `github.com/carojaspy/WeatherAPI/models.Task`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `task` (
	`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
	`is_active` bool NOT NULL DEFAULT FALSE ,
	`city` varchar(255) NOT NULL DEFAULT '' ,
	`country` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB;


-- drop table `weather_d_b`
--     DROP TABLE IF EXISTS `weather`
-- --------------------------------------------------
--  Table Structure for `WeatherAPI/models.WeatherDB`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `weather` (
	`id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
	`location` varchar(255) NOT NULL DEFAULT '' ,
	`temperature` varchar(255) NOT NULL DEFAULT '' ,
	`wind` varchar(255) NOT NULL DEFAULT '' ,
	`cloudines` varchar(255) NOT NULL DEFAULT '' ,
	`presure` varchar(255) NOT NULL DEFAULT '' ,
	`humidity` varchar(255) NOT NULL DEFAULT '' ,
	`sunrise` varchar(255) NOT NULL DEFAULT '' ,
	`sunset` varchar(255) NOT NULL DEFAULT '' ,
	`geo_coordinates` varchar(255) NOT NULL DEFAULT '' ,
	`requested_time` datetime NOT NULL
) ENGINE=InnoDB;


-- show tables;
select * from weather;
select * from task;
-- delete from task where id<10;

