package util

import (
	"database/sql"
)

/*
	传入sql语句 查询返回 json
*/
func Querys(sql string, db *sql.DB) interface{} {
	rows, _ := db.Query(sql)
	defer rows.Close()
	data, _ := Build(rows)
	return data
}

/*
	返回数据库查询结果集
*/
func Build(rows *sql.Rows) (interface{}, error) {
	columns, err := rows.Columns()
	if err != nil {
		return "", err
	}
	count := len(columns)
	tableData := make([]map[string]interface{}, 0)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	for rows.Next() {
		for i := 0; i < count; i++ {
			valuePtrs[i] = &values[i]
		}
		rows.Scan(valuePtrs...)
		entry := make(map[string]interface{})
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			entry[col] = v
		}
		tableData = append(tableData, entry)
	}
	if err != nil {
		return "", err
	}

	return tableData, nil
}

/*
	formatResult
	统一 JSON 返回格式
*/
func Fr(code int, message string, i interface{}) interface{} {
	f := make(map[string]interface{})
	f["code"] = code
	f["message"] = message
	f["data"] = i
	return f
}
