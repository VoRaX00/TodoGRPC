package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpTypeTask, DownTypeTask)
}

func UpTypeTask(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS type_tasks (
    	id SERIAL PRIMARY KEY,
    	type_task TEXT NOT NULL UNIQUE
	);`

	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownTypeTask(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS type_tasks;`
	_, err := tx.ExecContext(ctx, query)
	return err
}
