package internal

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/microsoft/go-mssqldb"
)
import "net/url"

// Connects to any supported database, and returns a open connection or error
func SqlConnect(dsn string) (*sql.DB, error) {
	u, err := url.Parse(dsn)
	if err != nil {
		return nil, fmt.Errorf("dsn for database must be in url format, %w", err)
	}

	var driver = ""
	switch u.Scheme {
	case "postgres", "postgresql":
		driver = "pgx"
	case "mysql":
		driver = "mysql"
	case "sqlserver", "mssql":
		driver = "sqlserver"
	default:
		return nil, fmt.Errorf("%s is not a supported database driver", u.Scheme)
	}

	return sql.Open(driver, dsn)
}

func RunSql(dsn string, query string, format OutputFormat) error {
	db, err := SqlConnect(dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()
	cols, err := rows.Columns()
	if err != nil {
		return err
	}

	err = format.Header(cols)
	if err != nil {
		return err
	}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		scanArgs := make([]interface{}, len(values))
		for i := range values {
			scanArgs[i] = &values[i]
		}
		if err := rows.Scan(scanArgs...); err != nil {
			return err
		}
		err = format.Write(values)
		if err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}
