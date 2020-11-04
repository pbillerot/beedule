--
-- table hugodoc 
CREATE TABLE "hugodoc" (
	"id"	INTEGER NOT NULL,
	"path"	varchar(100) NOT NULL,
	"base"	varchar(50) NOT NULL,
	"dir"	varchar(100) NOT NULL,
	"ext"	varchar(10),
	"isdir"	varchar(1),
	"level"	INTEGER,
	"title"	varchar(100),
	"date"	varchar(50),
	"draft"	varchar(10),
	"tags"	varchar(50),
	"categories"	varchar(50),
	"content"	TEXT,
	PRIMARY KEY("id" AUTOINCREMENT)
);
