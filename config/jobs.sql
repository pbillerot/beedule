
-- Table batch
CREATE TABLE "jobs" (
	"job_id"	INTEGER,
	"label"	VARCHAR(50),
	"chain_id"	VARCHAR(50),
	"sequence"	INTEGER,
	"sierreur"	VARCHAR(1),
	"active"	VARCHAR(1),
	"type"	VARCHAR(50),
	"commandes"	TEXT,
	"options"	TEXT,
	"etat"	VARCHAR(10),
	"result"	TEXT,
	"heuredebut"	DATETIME,
	"heurefin"	DATETIME,
	"dureemn"	INTEGER,
	"doc"	TEXT,
	PRIMARY KEY("job_id" AUTOINCREMENT)
);