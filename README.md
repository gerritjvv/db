# Overview

# Install

*Linux/Mac*

```
curl https://raw.githubusercontent.com/gerritjvv/db/refs/heads/main/install.sh | sh
```

*Windows*

```
irm https://raw.githubusercontent.com/gerritjvv/db/refs/heads/main/install.ps1 | iex 
```
# Requirements

Can take SQL commands for postgrsql or mysql.

### Command interface:

#### ENV VARS:

| VAR     | Description                                                                                                                                                    |
|---------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| DB_DSN  | Database connection url in URI format e.g `postgres://user:pass@localhost:5432/mydb?sslmode=disable` or `mysql://user:pass@localhost:3306/mydb?parseTime=true` |
| HOME    | The user's current home directory, most if not all OS set this variable                                                                                        | 
| DB_CONF | The directory wher configuration for this tool is stored, defaults to `DB_CONF=$HOME/.db/conf`                                                                 | 

#### Configuring a database connection

There are two ways:

* Environment variable `DB_DSN`
* DbConf configuration file.

The `DB_CONF` configuration file, is a single file which names the db, e.g. `mydb.yaml` yaml file. Yaml is friendly
enough to be human and LLM editable.

It contains a single yaml entry `DB_DSN`.

```yaml
db_dsn: <uri connection format string> 
```

#### Commands

| Command | Sub Command | Descripton                               | Options                                          |
|---------|-------------|------------------------------------------|--------------------------------------------------|
| Help    |             | Prints out help for this tool            |                                                  | 
| sql     |             | Query's the database                     | `-n <configured database name>`                  | 
| conf    | db-ls       | Lists the configured databases           |                                                  | 
| conf    | db-new      | Creates a configuration for a database   | `-n <name>`, `-dsn <uri database format string>` | 
| conf    | db-rm       | Removes the configuration for a database | `-n <name>`                                      |



## Developers

Only main branch runs binary builds and tags produce releases.

*Release*

```
git tag v<version>
git push --tags
```

