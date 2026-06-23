package main

import (
	"db/internal"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"
	"gopkg.in/yaml.v3"
)

var cli struct {
	SQL struct {
		Name   string `short:"n" help:"Configured database name"`
		SQL    string `short:"s" help:"Database SQL query"`
		Format string `short:"f" help:"Output format csv (default), tsv"`
	} `cmd:"" help:"Query the database"`

	Conf struct {
		DBLs struct {
		} `cmd:"" help:"List configured databases"`

		DBNew struct {
			Name string `short:"n" required:"" help:"Configuration name"`
			DSN  string `help:"Database URI"`
		} `cmd:"" help:"Create a database configuration"`

		DBRm struct {
			Name string `short:"n" required:"" help:"Configuration name"`
		} `cmd:"" help:"Remove a database configuration"`
	} `cmd:"" help:"Configuration commands"`
}

func main() {
	conf, err := internal.NewConf()

	if err != nil {
		panic(err)
	}

	ctx := kong.Parse(&cli)

	switch ctx.Command() {
	case "sql":
		fmt.Println("Query db")
		dbConf, err := loadDbConf(conf, cli.SQL.Name)
		if err != nil {
			panic(err)
		}

		formatter := chooseFormatter(cli.SQL.Format)
		defer formatter.Close()
		err = internal.RunSql(dbConf.DSN, cli.SQL.SQL, formatter)
		if err != nil {
			panic(err)
		}
	case "conf db-ls":
		err := listDbConfs(conf)
		if err != nil {
			panic(err)
		}
	case "conf db-new":

		err := createNewDbConf(conf, cli.Conf.DBNew.Name, cli.Conf.DBNew.DSN)
		if err != nil {
			panic(err)
		}
		fmt.Println("Added a database configuration")
	case "conf db-rm":
		err := removeDbConf(conf, cli.Conf.DBRm.Name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Removed database configuration")
	default:
		panic(fmt.Errorf("unknown command: %s", ctx.Command()))
	}
}

func chooseFormatter(format string) internal.OutputFormat {
	switch format {
	case "byte1":
		return internal.NewCSVFormatter(0x01)
	case "tsv":
		return internal.NewCSVFormatter('\t')
	default:
		return internal.NewCSVFormatter(',')
	}
}

func listDbConfs(conf *internal.Conf) error {
	matches, err := filepath.Glob(filepath.Join(conf.DbConfDir, "db-*.yaml"))
	if err != nil {
		return err
	}
	for _, match := range matches {
		base := filepath.Base(match)
		name := strings.TrimSuffix(base, filepath.Ext(base))
		name = strings.TrimPrefix(name, "db-")
		fmt.Println(name)
	}
	return nil
}

func loadDbConf(conf *internal.Conf, name string) (*internal.DBConf, error) {
	f := filepath.Join(conf.DbConfDir, "db-"+name+".yaml")
	dbConf := internal.DBConf{}
	bts, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bts, &dbConf)
	if err != nil {
		return nil, err
	}
	return &dbConf, nil
}

func removeDbConf(conf *internal.Conf, name string) error {
	f := filepath.Join(conf.DbConfDir, "db-"+name+".yaml")
	return os.Remove(f)
}

func createNewDbConf(conf *internal.Conf, name string, dsn string) error {
	if name == "" {
		return fmt.Errorf("config name cannot be empty")
	}
	if dsn == "" {
		return fmt.Errorf("dsn is required")
	}

	dbConf := internal.DBConf{
		DSN: dsn,
	}
	data, err := yaml.Marshal(dbConf)
	if err != nil {
		return err
	}

	dbConfFileName := filepath.Join(conf.DbConfDir, "db-"+name+".yaml")
	return os.WriteFile(
		dbConfFileName,
		data,
		0755)
}
