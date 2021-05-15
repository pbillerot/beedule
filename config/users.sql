
-- Table users
CREATE TABLE IF NOT EXISTS "users" (
	"user_name"	varchar(20),
	"user_password"	varchar(100),
	"user_email"	varchar(100),
	"user_is_admin"	INTEGER,
	"user_groupes"	TEXT,
	PRIMARY KEY("user_name")
);