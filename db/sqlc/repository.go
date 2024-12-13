package db

import (
	"database/sql"
)

// contract to be implemented by both mock and real db
type DatabaseContract interface {
	Querier
}

type Repository struct {
	*Queries
	db *sql.DB
}

func NewRepository(db *sql.DB) DatabaseContract {
	return &Repository{db: db, Queries: New(db)}
}

// func (r *Repository) execTx(ctx context.Context, fn func(queries *Queries) error) error {
// 	sqlTx, err := r.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	q := New(sqlTx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := sqlTx.Rollback(); rbErr != nil {
// 			return rbErr
// 		}
// 		return err
// 	}

// 	return sqlTx.Commit()
// }
