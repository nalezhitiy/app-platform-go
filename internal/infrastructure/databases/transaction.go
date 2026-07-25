package databases

import (
	"context"

	"gorm.io/gorm"
)

type transactionKey struct{}

func withTransaction(
	ctx context.Context,
	tx *gorm.DB,
) context.Context {
	return context.WithValue(
		ctx,
		transactionKey{},
		tx,
	)
}

func transactionFromContext(
	ctx context.Context,
) (*gorm.DB, bool) {

	tx, ok := ctx.Value(transactionKey{}).(*gorm.DB)

	return tx, ok
}
