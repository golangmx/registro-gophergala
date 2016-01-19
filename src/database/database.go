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

	t := "CREATE TABLE IF NOT EXISTS teams " +
		"(id serial primary key, nombre varchar(30), " +
		"proyecto varchar(30));"
	_, err = db.Exec(t)
	if err != nil {
		log.Fatalf("imposible crear base de datos 'teams': %s",
			err.Error())
	}

	u := "CREATE TABLE IF NOT EXISTS " +
		"users (id serial primary key, " +
		"team int references teams(id), " +
		"nombre varchar(30), tipo_id int, num_id varchar(50));"
	_, err = db.Exec(u)
	if err != nil {
		log.Fatalf("imposible crear base de datos 'users': %s",
			err.Error())
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
