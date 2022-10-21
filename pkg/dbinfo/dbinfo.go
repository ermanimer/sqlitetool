package dbinfo

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

// Tables returns all table names.
func Tables(ctx context.Context, path string) ([]string, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err

	}
	defer db.Close()

	query := "SELECT name FROM sqlite_master WHERE type='table' ORDER BY name ASC;"

	statement, err := db.Prepare(query)
	if err != nil {
		return nil, err

	}

	res, err := statement.QueryContext(ctx)
	if err != nil {
		return nil, err

	}

	var tables []string
	for res.Next() {
		var table string
		if err := res.Scan(&table); err != nil {
			return nil, err

		}

		tables = append(tables, table)

	}

	return tables, nil
}
