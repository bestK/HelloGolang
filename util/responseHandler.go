package util

import (
	"database/sql"
	"github.com/labstack/echo"
	"net/http"
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
func Fr(code int, message string, data interface{}) interface{} {
	response := make(map[string]interface{})
	response["code"] = code
	response["message"] = message
	response["data"] = data
	response["success"] = code == http.StatusOK
	return response
}

func Ok(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, Fr(http.StatusOK, message, data))
}

func Bad(c echo.Context, code int, message string) error {
	return c.JSON(code, Fr(code, message, nil))
}
