package repository

import (
	"context"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"GroupUp/internal/model"
	"GroupUp/pkg/log"
)

type TxKey string

const ctxTxKey TxKey = "TxKey"

type Repository struct {
	db     *gorm.DB
	logger *log.Logger
}

func NewRepository(logger *log.Logger, db *gorm.DB) *Repository {
	return &Repository{
		db:     db,
		logger: logger,
	}
}

func NewDb(conf *viper.Viper, l *log.Logger) *gorm.DB {
	l.Info("opening database connection")
	db, err := gorm.Open(sqlite.Open(conf.GetString("data.sqlite.filename")))
	if err != nil {
		l.Error("could not open a connection to the database properly:", zap.Error(err))
		return nil
	}
	err = db.AutoMigrate(
		&model.User{},
		&model.Group{},
		&model.Course{},
		&model.UserTag{},
		&model.Rating{},
		&model.RatingTag{},
	)
	if err != nil {
		l.Error("could not automigrate the dabase properly:", zap.Error(err))
		return nil
	}
	return db
}

func (r *Repository) DB(ctx context.Context) *gorm.DB {
	v := ctx.Value(ctxTxKey)
	if v != nil {
		if tx, ok := v.(*gorm.DB); ok {
			return tx
		}
	}
	return r.db.WithContext(ctx)
}

func (r *Repository) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, ctxTxKey, tx)
		return fn(ctx)
	})
}
