package tableinfo

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

// Columns return all column names and types.
func Columns(ctx context.Context, path string, name string) ([][]string, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err

	}
	defer db.Close()

	query := "SELECT name, type FROM pragma_table_info(?) ORDER BY name ASC;"

	statement, err := db.Prepare(query)
	if err != nil {
		return nil, err

	}

	res, err := statement.QueryContext(ctx, name)
	if err != nil {
		return nil, err

	}

	var columns [][]string
	for res.Next() {
		var name string
		var dataType string
		if err := res.Scan(&name, &dataType); err != nil {
			return nil, err

		}

		column := []string{name, dataType}

		columns = append(columns, column)

	}

	return columns, nil
}
