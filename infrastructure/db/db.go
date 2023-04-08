package db

import (
	"context"
	"database/sql"
	"errors"
	"os"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

type db struct {
	db *sql.DB
}

func Init() {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	if err := DB.Ping(); err != nil {
		panic(err)
	}

	boil.DebugMode = true
	boil.SetDB(DB)
}

func NewDB() db {
	return db{
		db: DB,
	}
}

func (d db) Error(err error) error {
	if err == nil || errors.Is(sql.ErrNoRows, err) {
		return nil
	}
	return err
}

func InTx(ctx context.Context, f func(c context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	result, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	return result, nil
}
