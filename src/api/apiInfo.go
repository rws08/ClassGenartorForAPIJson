package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ObjInfo ... info 구조
type ObjInfo struct {
	Key       int    `json:"key"`        // key
	ServerKey int    `json:"server_key"` // server_key
	Prefix    string `json:"prefix"`     // prefix
}

// CreateInfo ...
func CreateInfo(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	// key={n}&prefix=
	_serverKey, _ := c.GetQuery("key")
	_prefix, _ := c.GetQuery("prefix")

	// 정보 추가
	result, err := baseDB.Exec("INSERT INTO info(server_key, prefix) VALUES ($1, $2)", _serverKey, _prefix)
	if err != nil {
		ReturnError(c, err)
		return
	}

	infoKey, _ := result.LastInsertId()

	if c.Keys == nil {
		return
	}

	ret := map[string]interface{}{}
	ret["info_key"] = fmt.Sprint(infoKey)
	c.JSON(http.StatusOK, ret)
}

// ReadInfo ...
func ReadInfo(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	// server_key=3
	_serverKey, _ := c.GetQuery("server_key")

	row, err := baseDB.Query("SELECT * FROM info WHERE server_key = $1", _serverKey)
	defer row.Close()
	if err != nil {
		log.Fatal(err)
	}

	var obj ObjInfo
	for row.Next() {
		row.Scan(&obj.Key, &obj.ServerKey, &obj.Prefix)
		log.Println("info: ", obj.Prefix)
	}

	ret := map[string]interface{}{}
	ret["INFO"] = obj
	c.JSON(http.StatusOK, ret)
}

// UpdateInfo ...
func UpdateInfo(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	// server_key=3&prefix=test
	_serverKey, _ := c.GetQuery("server_key")
	_prefix, _ := c.GetQuery("prefix")

	if len(_serverKey) > 0 {
		row, _ := baseDB.Query("SELECT key FROM info WHERE server_key = " + _serverKey)
		if !row.Next() {
			ReturnError(c, errors.New("error : not key"))
			return
		}
		row.Close()
	} else {
		ReturnError(c, errors.New("error : required key"))
		return
	}

	strQuery := MakeUpdateQuery("info", []string{"prefix"}, []interface{}{_prefix}, "server_key = "+_serverKey)

	log.Println(strQuery)
	// 정보 변경
	_, err := baseDB.Exec(strQuery)
	if err != nil {
		ReturnError(c, err)
		return
	}

	ret := map[string]interface{}{}
	ret["SERVER_KEY"] = _serverKey
	c.JSON(http.StatusOK, ret)
}

// DelInfo ...
func DelInfo(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	// server_key=3
	_serverKey, _ := c.GetQuery("server_key")

	// 서버 삭제
	_, err := baseDB.Exec("DELETE FROM info WHERE server_key = $1", _serverKey)
	if err != nil {
		ReturnError(c, err)
		return
	}

	ret := map[string]interface{}{}
	c.JSON(http.StatusOK, ret)
}
