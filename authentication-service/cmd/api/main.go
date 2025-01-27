package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"auth.svc/data"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const webPort = "80"

var Counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Println("starting authentication service")

	// TODO connect DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("can't connect to postgress")
	}

	// setup config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn, exists := os.LookupEnv("DSN")
	if !exists || dsn == "" {
		log.Panic("DSN environment variable is not set")
	}

	fmt.Println("dsn: ", dsn)

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Println("postgres not yet ready..: ", err.Error())
			Counts++
		} else {
			log.Println("Connected to postgress")
			return connection
		}

		if Counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off two seconds")
		time.Sleep(2 * time.Second)
	}
}
