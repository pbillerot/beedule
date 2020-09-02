-- Table tasks
CREATE TABLE IF NOT EXISTS "tasks" (
	"task_id"	INTEGER,
	"task_name"	varchar(100) NOT NULL,
	"task_user"	varchar(20) NOT NULL,
	"task_status"	varchar(20) NOT NULL,
	"task_note"	TEXT,
	PRIMARY KEY("task_id" AUTOINCREMENT)
);