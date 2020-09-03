-- Table tasks

-- sqlite
CREATE TABLE IF NOT EXISTS "tasks" (
	"task_id"	INTEGER,
	"task_name"	varchar(100) NOT NULL,
	"task_user"	varchar(20) NOT NULL,
	"task_status"	varchar(20) NOT NULL,
	"task_note"	TEXT,
	PRIMARY KEY("task_id" AUTOINCREMENT)
);

-- Mysql
CREATE TABLE tasks (
	task_id	INT NOT NULL AUTO_INCREMENT,
	task_name	varchar(100) NOT NULL,
	task_user	varchar(20) NOT NULL,
	task_status	varchar(20) NOT NULL,
	task_note	TEXT,
	PRIMARY KEY(task_id)
);

CREATE USER 'beedule'@'localhost' IDENTIFIED BY 'beedule';
GRANT SELECT, INSERT, UPDATE, DELETE ON beedule TO 'beedule'@'localhost';