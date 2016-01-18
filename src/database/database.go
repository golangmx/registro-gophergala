package database

import (
	"database/sql"
	"log"

	// importar driver de Postgres
	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	// declarar err para evitar hacer shadowing a db
	var err error
	db, err = sql.Open("postgres",
		"host=db user=ggala password=gopherz dbname=gophergala sslmode=disable")
	if err != nil {
		log.Fatalf("no se puede abrir conexión a db: %s", err.Error())
	}
}

// DB proporciona el handler a la base de datos
func DB() *sql.DB {
	return db
}

// Dispose cierra la conexión a la base de datos
func Dispose() {
	db.Close()
}
