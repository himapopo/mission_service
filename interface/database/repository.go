package database

import (
	"context"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type dbUtil interface {
	Error(error) error
	GetDao(context.Context) boil.ContextExecutor
}
