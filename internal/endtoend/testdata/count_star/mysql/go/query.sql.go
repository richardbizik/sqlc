// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package querytest

import (
	"context"
)

const countStarLower = `-- name: CountStarLower :one
SELECT count(*) FROM bar
`

func (q *Queries) CountStarLower(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countStarLower)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countStarUpper = `-- name: CountStarUpper :one
SELECT COUNT(*) FROM bar
`

func (q *Queries) CountStarUpper(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, countStarUpper)
	var count int64
	err := row.Scan(&count)
	return count, err
}
