drop table if exists `profile`;
CREATE TABLE `profile` (
    `user_id` INTEGER PRIMARY KEY,
	`image` TEXT,
    `score` INTEGER,
    `create_date` DATE,
    `edit_date` DATE
);

drop table if EXISTS `party`;
CREATE TABLE `party` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT,
    `meet_time` DATE,
    `latitude` TEXT,
    `longitude` TEXT,
    `total_people` INTEGER,
    `current_people` INTEGER DEFAULT 1,
    `create_date` DATE,
    `edit_date` DATE
);

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `name` TEXT,
    `email` TEXT UNIQUE,
    `auth`  TEXT,
    `create_date` DATE,
    `edit_date` DATE
);

DROP TABLE IF EXISTS `party_member`
CREATE TABLE `party_member` (
    `user_id` INTEGER PRIMARY KEY,
    `party_id` INTEGER
);