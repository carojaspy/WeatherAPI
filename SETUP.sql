drop table `weather_d_b`
    DROP TABLE IF EXISTS `weather_d_b`

create table `weather_d_b`
    -- --------------------------------------------------
    --  Table Structure for `WeatherAPI/models.WeatherDB`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `weather_d_b` (
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