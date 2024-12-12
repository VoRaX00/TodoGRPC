package migrations

import (
	"context"
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(UpTasks, DownTasks)
}

func UpTasks(ctx context.Context, tx *sql.Tx) error {
	query := `CREATE TABLE IF NOT EXISTS tasks (
    	id SERIAL PRIMARY KEY,
    	name_task VARCHAR(255) NOT NULL,
    	descriptions TEXT NOT NULL,
    	type_task_id INT REFERENCES type_tasks (id) NOT NULL,
    	deadline TIMESTAMP NOT NULL
	);`
	_, err := tx.ExecContext(ctx, query)
	return err
}

func DownTasks(ctx context.Context, tx *sql.Tx) error {
	query := `DROP TABLE IF EXISTS tasks;`

	_, err := tx.ExecContext(ctx, query)
	return err
}
