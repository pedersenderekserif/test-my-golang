package utils

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"

	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang-migrate/migrate/v4"
	postgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

var dbPool pgxpoolmock.PgxPool

func DbUrl() string {
	return fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v",
		os.Getenv("PG_USER"),
		url.QueryEscape(os.Getenv("PG_PASSWORD")),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DATABASE"),
	)
}

func DbMigrations() error {

	url := DbUrl() + "?sslmode=disable"

	db, err := sql.Open("postgres", url)
	if err != nil {
		return err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		os.Getenv("DB_MIGRATION_FOLDER"),
		os.Getenv("PG_DATABASE"),
		driver)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func NewDbPool() (*pgxpool.Pool, error) {

	config, err := pgxpool.ParseConfig(DbUrl())
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return pool, nil
}

func InitDbPool() error {
	db, err := NewDbPool()
	if err != nil {
		logrus.Fatal(err)
		return err
	}
	SetDbPool(db)
	return nil
}

func Ping() (result string, err error) {
	sql := `SELECT 'JOHNNY 5 ALIVE' AS ping `
	var rows pgx.Rows
	rows, err = GetDbPool().Query(context.Background(), sql)

	if err != nil {
		logrus.Debugf("Query failed: %v\n", err)
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&result,
		)
		if err != nil {
			logrus.Debugf("Scan failed: %v\n", err)
			return result, err
		}
	}
	return result, err
}

func SetDbPool(dbpool pgxpoolmock.PgxPool) {
	dbPool = dbpool
}

func GetDbPool() pgxpoolmock.PgxPool {
	return dbPool
}
