package model

const templateModel = `// Code generated by {{.Generator}}. DO NOT EDIT.
package {{.Package}}

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"strings"{{.Imports}}
)

const (
	{{.Prefix}}ModelTableName = "{{.TableName}}"
)

func {{.Prefix}}ModelSetup(db *sql.DB) error {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS {{.TableName}} ({{.SQL.Setup}});")
	return err
}

func {{.Prefix}}ModelGet(db *sql.DB, key {{.PrimaryKey.GoType}}) (*{{.ModelIdent}}, int, error) {
	m := &{{.ModelIdent}}{}
	if err := db.QueryRow("SELECT {{.SQL.DBNames}} FROM {{.TableName}} WHERE {{.PrimaryKey.DBName}} = $1;", key).Scan({{.SQL.IdentRefs}}); err != nil {
		if err == sql.ErrNoRows {
			return nil, 2, err
		}
		if postgresErr, ok := err.(*pq.Error); ok {
			switch postgresErr.Code {
			case "42P01": // undefined_table
				return nil, 4, err
			default:
				return nil, 0, err
			}
		}
		return nil, 0, err
	}
	return m, 0, nil
}

func {{.Prefix}}ModelInsert(db *sql.DB, m *{{.ModelIdent}}) (int, error) {
	_, err := db.Exec("INSERT INTO {{.TableName}} ({{.SQL.DBNames}}) VALUES ({{.SQL.Placeholders}});", {{.SQL.Idents}})
	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok {
			switch postgresErr.Code {
			case "23505": // unique_violation
				return 3, err
			default:
				return 0, err
			}
		}
	}
	return 0, nil
}

func {{.Prefix}}ModelInsertBulk(db *sql.DB, models []{{.ModelIdent}}, allowConflict bool) (int, error) {
	conflictSQL := ""
	if allowConflict {
		conflictSQL = " ON CONFLICT DO NOTHING"
	}
	placeholders := make([]string, 0, len(models))
	args := make([]interface{}, 0, len(models)*{{.SQL.ColNum}})
	for c, m := range models {
		n := c * {{.SQL.ColNum}}
		placeholders = append(placeholders, fmt.Sprintf("({{.SQL.PlaceholderTpl}})", {{.SQL.PlaceholderCount}}))
		args = append(args, {{.SQL.Idents}})
	}
	_, err := db.Exec("INSERT INTO {{.TableName}} ({{.SQL.DBNames}}) VALUES "+strings.Join(placeholders, ", ")+conflictSQL+";", args...)
	if err != nil {
		if postgresErr, ok := err.(*pq.Error); ok {
			switch postgresErr.Code {
			case "23505": // unique_violation
				return 3, err
			default:
				return 0, err
			}
		}
	}
	return 0, nil
}

func {{.Prefix}}ModelUpdate(db *sql.DB, m *{{.ModelIdent}}) error {
	_, err := db.Exec("UPDATE {{.TableName}} SET ({{.SQL.DBNames}}) = ({{.SQL.Placeholders}}) WHERE {{.PrimaryKey.DBName}} = ${{.PrimaryKey.Num}};", {{.SQL.Idents}})
	return err
}

func {{.Prefix}}ModelDelete(db *sql.DB, m *{{.ModelIdent}}) error {
	_, err := db.Exec("DELETE FROM {{.TableName}} WHERE {{.PrimaryKey.DBName}} = $1;", m.{{.PrimaryKey.Ident}})
	return err
}
`
