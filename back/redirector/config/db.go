package config

import (
	"context"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/v5/stdlib" // PGX driver
	"github.com/jmoiron/sqlx"
)

var (
	dbClient *DB
)

type DB struct {
	sqlxDB  *sqlx.DB
	builder squirrel.StatementBuilderType
}

func SetupDb() (*DB, error) {
	db_url := os.Getenv("DATABASE_URL")

	db, err := sqlx.Open("pgx", db_url)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	builder := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	dbClient = &DB{sqlxDB: db, builder: builder}

	return dbClient, nil
}

func (d *DB) Close() error {
	return d.sqlxDB.Close()
}

func (d *DB) QueryExecutor() squirrel.BaseRunner { return d.sqlxDB }

func (d *DB) SquirrelBuilder() squirrel.StatementBuilderType { return d.builder }

func GetDb() *DB {
	return dbClient
}
