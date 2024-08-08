package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgxutil"
)

type (
	QueryPgx interface {
		Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
		Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
		QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	}

	Entity interface {
		TableName() string
		Columns() []string
		Data() map[string]any
	}

	PostgresRepo[T Entity] struct {
		db               QueryPgx
		QueryBasicSelect string
		QueryBasicCreate string
		QueryBasicUpdate string
	}
)

func NewPostgresRepo[T Entity](dbPool QueryPgx) *PostgresRepo[T] {
	var table T

	namedColumns := []string{}
	namedColumnsForUpdate := []string{}
	for _, column := range table.Columns() {
		namedColumns = append(namedColumns, fmt.Sprintf("@%s", column))
		namedColumnsForUpdate = append(namedColumnsForUpdate, fmt.Sprintf("%s=@%s", column, column))
	}

	columns := []string{}
	columns = append(columns, table.Columns()...)
	columns = append(columns, "created_at", "updated_at")

	tableName := table.TableName()

	return &PostgresRepo[T]{
		db: dbPool,
		QueryBasicSelect: fmt.Sprintf(`SELECT %s FROM %s `,
			strings.Join(columns, ","),
			tableName,
		),
		QueryBasicCreate: fmt.Sprintf("INSERT INTO %s(%s) VALUES (%s)",
			tableName,
			strings.Join(table.Columns(), ","),
			strings.Join(namedColumns, ","),
		),
		QueryBasicUpdate: fmt.Sprintf(`UPDATE %s SET %s WHERE id=@id`,
			tableName,
			strings.Join(namedColumnsForUpdate, ","),
		),
	}
}

func (pr *PostgresRepo[T]) Get(ctx context.Context, by string, val any) (*T, error) {
	q := fmt.Sprintf(pr.QueryBasicSelect+` WHERE %s = $1 LIMIT 1`, by)
	v, err := pgxutil.SelectRow(ctx, pr.db, q, []any{val}, pgx.RowToStructByNameLax[T])
	if err != nil {
		return nil, err
	}
	return &v, nil
}

func (pr *PostgresRepo[T]) GetAll(ctx context.Context) ([]T, error) {
	v, err := pgxutil.Select(ctx, pr.db, pr.QueryBasicSelect, nil, pgx.RowToStructByNameLax[T])
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (pr *PostgresRepo[T]) Create(ctx context.Context, payloads ...T) error {
	for _, payload := range payloads {

		var args pgx.NamedArgs = payload.Data()

		_, err := pr.db.Exec(ctx, pr.QueryBasicCreate, args)
		if err != nil {
			return err
		}
	}

	return nil
}

func (pr *PostgresRepo[T]) Edit(ctx context.Context, payloads ...T) error {
	for _, payload := range payloads {

		var args pgx.NamedArgs = payload.Data()

		_, err := pr.db.Exec(ctx, pr.QueryBasicUpdate, args)
		if err != nil {
			return err
		}
	}

	return nil
}
