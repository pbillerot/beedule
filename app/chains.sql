
-- Table chains
CREATE TABLE "chains" (
	"chain_id"	INTEGER,
	"label"	VARCHAR(50),
	"planif"	VARCHAR(50),
	"active"	VARCHAR(1),
	"etat"	VARCHAR(10),
	"result"	TEXT,
	"doc"	TEXT,
	"email"	TEXT,
	"heuredebut"	DATETIME,
	"heurefin"	DATETIME,
	"dureemn"	INTEGER,
	PRIMARY KEY("chain_id" AUTOINCREMENT)
)