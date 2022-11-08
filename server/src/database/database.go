package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"log"
	"server/src/lib"
)

func ConnectToDb() *sqlx.DB {
	var (
		host     = lib.GetEnvWithDefault("PGHOST", "localhost")
		port     = lib.GetEnvWithDefault("PGPORT", "5432")
		user     = lib.GetEnvWithDefault("PGUSER", "postgres")
		password = lib.GetEnvWithDefault("PGPASSWORD", "password")
		dbname   = lib.GetEnvWithDefault("PGDATABASE", "postgres")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func InsertWeight(db *sqlx.DB, userID string, name string, weight float32) error {
	result := db.MustExec("INSERT INTO weights(user_id, name, weight) values($1, $2, $3)", userID, name, weight)
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Could not insert weight")
	}
	if rows != 1 {
		return errors.New("No weights inserted")
	}
	return nil
}

func DeleteTodaysWeight(db *sqlx.DB, userID string) error {
	result := db.MustExec("DELETE FROM weights WHERE user_id = $1 AND created_at::date = CURRENT_DATE", userID)
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "Could not delete weight")
	}
	if rows == 0 {
		return errors.New("No weights deleted")
	}
	return nil
}

func GetTodaysWeight(db *sqlx.DB, userID string) (float32, error) {
	row := db.QueryRow("SELECT weight FROM weights WHERE user_id = $1 AND created_at::date = CURRENT_DATE", userID)
	var weight float32
	err := row.Scan(&weight)
	if err != nil {
		return 0, errors.Wrap(err, "Could not find any weight today for the corresponding user")
	}
	return weight, nil
}
