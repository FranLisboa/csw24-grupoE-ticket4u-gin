package database

import (
	"database/sql"
	//"fmt"
	"log"
	//"os"

	_ "github.com/lib/pq"
)

func StartDB() *sql.DB {
 	//password :=  os.Getenv("DB_PASSWORD")

	const (
		host     = "postgres.cy3myhw5bsdp.us-east-1.rds.amazonaws.com"
		port     = 5432
		user     = "postgres"
		dbname   = "postgres"
	)

	//var psqlconn = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", "postgresql://postgres:xyV4YBeY8Qz2FuZ@postgres.cy3myhw5bsdp.us-east-1.rds.amazonaws.com:5432/postgres")
	if err != nil {
		log.Fatalf("Erro ao abrir a conexão com o banco de dados: %v", err)
	}


	err = db.Ping()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso")
	return db
}
