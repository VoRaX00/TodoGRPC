package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(AddColumnToTasksUserId, DropColumnToTasksUserId)
}

func AddColumnToTasksUserId(ctx context.Context, tx *sql.Tx) error {
	query := `ALTER TABLE tasks ADD COLUMN user_id INT NOT NULL DEFAULT 0;`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func DropColumnToTasksUserId(ctx context.Context, tx *sql.Tx) error {
	query := `ALTER TABLE tasks DROP COLUMN user_id;`
	_, err := tx.ExecContext(ctx, query)
	return err
}
