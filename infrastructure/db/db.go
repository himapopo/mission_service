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

var TxKey struct{}

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

func (d db) GetDao(ctx context.Context) boil.ContextExecutor {
	tx, ok := ctx.Value(TxKey).(*sql.Tx)
	if ok {
		return tx
	}

	return d.db
}

func (d db) Error(err error) error {
	if err == nil || errors.Is(sql.ErrNoRows, err) {
		return nil
	}
	return err
}

func InTx(ctx context.Context, f func(context.Context) error) error {
	tx, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, TxKey, tx)

	if err := f(ctx); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
