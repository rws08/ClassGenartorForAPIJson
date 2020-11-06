package api

import (
	"strings"

	"github.com/gin-gonic/gin"
)

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

// AddQuery ... AddQuery(c, []string{"key2"}, []string{"val2"}) -> key1=val1&key2=val2
func AddQuery(c *gin.Context, keys []string, vals []string) {
	for idx, val := range vals {
		if len(val) > 0 {
			c.Request.URL.RawQuery += "&" + keys[idx] + "=" + val
		}
	}
}
