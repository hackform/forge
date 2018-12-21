package main

const templateQueryGroupEq = `
func {{.Prefix}}ModelGet{{.ModelIdent}}Eq{{.PrimaryField.Cond.Ident}}Ord{{.PrimaryField.Ident}}(db *sql.DB, key {{.PrimaryField.Cond.GoType}}, orderasc bool, limit, offset int) ([]{{.ModelIdent}}, error) {
	order := "DESC"
	if orderasc {
		order = "ASC"
	}
	res := make([]{{.ModelIdent}}, 0, limit)
	rows, err := db.Query("SELECT {{.SQL.DBNames}} FROM {{.TableName}} WHERE {{.PrimaryField.Cond.DBName}} = $1 ORDER BY {{.PrimaryField.DBName}} "+order+" LIMIT $2 OFFSET $3;", key, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
		}
	}()
	for rows.Next() {
		m := {{.ModelIdent}}{}
		if err := rows.Scan({{.SQL.IdentRefs}}); err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return res, nil
}
`