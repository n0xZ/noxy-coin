package database

import (
	"database/sql"
	"log"
	"noxy/internal/config"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Service interface {
	Close() error
}

type dbService struct {
	client *sql.DB
}

var (
	dbInstance *dbService
)

func New() *dbService {
	if dbInstance != nil {
		return dbInstance
	}
	dbUrl := config.Get("TURSO_DATABASE_URL")
	dbAuthToken := config.Get("TURSO_DATABASE_TOKEN")
	dbUrlAuthenticated := dbUrl + "?authToken=" + dbAuthToken

	db, err := sql.Open("libsql", dbUrlAuthenticated)

	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &dbService{
		client: db,
	}
	return dbInstance
}
func (db *dbService) Close() error {

	return db.client.Close()
}
