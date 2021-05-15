
-- Table groups
CREATE TABLE "orders" (
	"orders_id"	INTEGER NOT NULL,
	"orders_ptf_id"	TEXT DEFAULT '',
	"orders_order"	TEXT NOT NULL '',
	"orders_time"	TEXT DEFAULT '',
	"orders_quote"	REAL DEFAULT 0,
	"orders_quantity"	INTEGER DEFAULT 0,
	"orders_buy"	REAL DEFAULT 0,
	"orders_sell"	REAL DEFAULT 0,
	"orders_cost_price"	REAL DEFAULT 0,
	"orders_cost"	REAL DEFAULT 0,
	"orders_debit"	REAL DEFAULT 0,
	"orders_credit"	REAL DEFAULT 0,
	"orders_gain"	REAL DEFAULT 0,
	"orders_gainp"	REAL DEFAULT 0,
	"orders_sell_time"	TEXT DEFAULT '',
	"orders_sell_cost"	REAL DEFAULT 0,
	"orders_sell_gain"	REAL DEFAULT 0,
	"orders_sell_gainp"	REAL DEFAULT 0,
	"orders_rem"	TEXT DEFAULT '',
	PRIMARY KEY("orders_id" AUTOINCREMENT)
);

-- REPRISE
UPDATE "orders" set orders_ptf_id = '' where orders_ptf_id is null;
UPDATE "orders" set orders_order = '' where orders_order is null;
UPDATE "orders" set orders_time = '' where orders_time is null;
UPDATE "orders" set orders_quote = 0 where orders_quote is null;
UPDATE "orders" set orders_quantity = 0 where orders_quantity is null;
UPDATE "orders" set orders_buy = 0 where orders_buy is null;
UPDATE "orders" set orders_sell = 0 where orders_sell is null;
UPDATE "orders" set orders_cost_price = 0 where orders_cost_price is null;
UPDATE "orders" set orders_cost = 0 where orders_cost is null;
UPDATE "orders" set orders_debit = 0 where orders_debit is null;
UPDATE "orders" set orders_credit = 0 where orders_credit is null;
UPDATE "orders" set orders_gain = 0 where orders_gain is null;
UPDATE "orders" set orders_gainp = 0 where orders_gainp is null;
UPDATE "orders" set orders_sell_time = '' where orders_sell_time is null;
UPDATE "orders" set orders_sell_cost = 0 where orders_sell_cost is null;
UPDATE "orders" set orders_sell_gain = 0 where orders_sell_gain is null;
UPDATE "orders" set orders_sell_gainp = 0 where orders_sell_gainp is null;
UPDATE "orders" set orders_rem = 0 where orders_rem is null;

