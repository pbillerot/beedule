
-- Table batch
CREATE TABLE "batch" (
	"id"	INTEGER,
	"label"	VARCHAR(200),
	"chaine"	VARCHAR(50),
	"planif"	VARCHAR(50),
	"sequence"	INTEGER,
	"sierreur"	VARCHAR(1),
	"actif"	VARCHAR(1),
	"etat"	VARCHAR(10),
	"type"	VARCHAR(50),
	"commandes"	TEXT,
	"options"	TEXT,
	"result"	TEXT,
	"nexec"	INTEGER,
	"heureresult"	DATETIME,
	"doc"	TEXT,
	"email"	TEXT,
	"heudedebut"	DATETIME,
	"heurefin"	DATETIME,
	"dureemn"	INTEGER,
	PRIMARY KEY("id" AUTOINCREMENT)
)