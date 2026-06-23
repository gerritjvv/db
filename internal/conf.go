package internal

import (
	"errors"
	"os"
	"path/filepath"
)

type DBConf struct {
	DSN string `yaml:"db_dsn"`
}

type Conf struct {

	// Calculated either from HOME or the DB_CONF environment variable
	DbConfDir string
	// value of the DB_DSN variable
	DBDSN string
}

func NewConf() (*Conf, error) {

	dbConfDir := os.Getenv("DB_CONF")
	dbDSN := os.Getenv("DB_DSN")

	if dbConfDir == "" {
		home := os.Getenv("HOME")
		if home == "" {
			return nil, errors.New("$HOME or DB_CONF must be set")
		}

		dbConfDir = filepath.Join(home, ".db", "conf")
		err1 := os.MkdirAll(dbConfDir, 0755)
		if err1 != nil {
			return nil, err1
		}
	}

	return &Conf{
		DbConfDir: dbConfDir,
		DBDSN:     dbDSN,
	}, nil
}
