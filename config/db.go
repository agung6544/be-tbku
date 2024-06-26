package config

import (
	"os"
	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var DBUlbimongoinfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "TBWS",
}

var Ulbimongoconn = atdb.MongoConnect(DBUlbimongoinfo)