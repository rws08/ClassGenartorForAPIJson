package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ObjInfo ... info 구조
type ObjInfo struct {
	Key       int    `json:"key"`        // id
	ServerKey string `json:"server_key"` // server_key
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

}

// UpdateInfo ...
func UpdateInfo(c *gin.Context) {

}

// DelInfo ...
func DelInfo(c *gin.Context) {

}
