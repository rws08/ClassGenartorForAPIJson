package api

import "database/sql"

// ServerName ...
const ServerName = `localhost`

// ServerPort ...
const ServerPort = `8000`

// DriverName ...
const DriverName = `sqlite3`

// DbName ...
const DbName = `sqlite-apiclass.db`

// QueryCreateName ...
const QueryCreateName = `JsonToClass.sql`

// OpenDB ...
func OpenDB() *sql.DB {
	baseDB, _ := sql.Open(DriverName, DbName)
	return baseDB
}
