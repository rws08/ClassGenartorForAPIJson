package api

import "strings"

// MakeUpdateQuery ... make update query
func MakeUpdateQuery(table string, keys []string, vals []interface{}, where string) string {
	strQuery := "UPDATE " + table + " SET "

	for idx, val := range vals {
		str := val.(string)
		if len(str) > 0 {
			strQuery += keys[idx] + " = \"" + str + "\", "
		}
	}

	strQuery = strings.TrimRight(strQuery, ", ")
	strQuery += " WHERE " + where

	return strQuery
}
