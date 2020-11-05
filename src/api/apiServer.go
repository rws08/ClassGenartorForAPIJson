package api

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ObjServer ... server 구조
type ObjServer struct {
	Key  int    `json:"key"`         // id
	Name string `json:"name"`        // name
	URL  string `json:"url"`         // url
	Desc string `json:"description"` // desc
}

// CreateServer ...
func CreateServer(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	// name=main&url=https%3A%2F%2Fapi.com&description=설명
	_name, _ := c.GetQuery("name")
	_url, _ := c.GetQuery("url")
	_description, _ := c.GetQuery("description")

	// 서버 추가
	result, err := baseDB.Exec("INSERT INTO server(name, url, description) VALUES ($1, $2, $3)", _name, _url, _description)
	if err != nil {
		ReturnError(c, err)
		return
	}
	serverKey, _ := result.LastInsertId()
	// 서버 정보 생성
	result, err = baseDB.Exec("INSERT INTO info(server_key, prefix) VALUES ($1, $2)", serverKey, "")
	if err != nil {
		ReturnError(c, err)
		return
	}

	ret := map[string]interface{}{}
	ret["server_key"] = fmt.Sprint(serverKey)
	c.JSON(http.StatusOK, ret)
}

// ReadServer ...
func ReadServer(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	row, err := baseDB.Query("SELECT * FROM server ORDER BY name")
	defer row.Close()
	if err != nil {
		log.Fatal(err)
	}

	arrObj := []ObjServer{}
	for row.Next() {
		var obj ObjServer
		row.Scan(&obj.Key, &obj.Name, &obj.URL, &obj.Desc)
		arrObj = append(arrObj, obj)
		log.Println("server: ", obj.Name, " ", obj.URL, " ", obj.Desc)
	}

	ret := map[string]interface{}{}
	ret["LIST"] = arrObj
	c.JSON(http.StatusOK, ret)
}

// UpdateServer ...
func UpdateServer(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	// key=3&name=main&url=https%3A%2F%2Fapi.com&description=설명
	_key, _ := c.GetQuery("key")
	_name, _ := c.GetQuery("name")
	_url, _ := c.GetQuery("url")
	_description, _ := c.GetQuery("description")

	if len(_key) > 0 {
		row, _ := baseDB.Query("SELECT key FROM server WHERE key = " + _key)
		if !row.Next() {
			ReturnError(c, errors.New("error : not key"))
			return
		}
		row.Close()
	} else {
		ReturnError(c, errors.New("error : required key"))
		return
	}

	strQuery := MakeUpdateQuery("server", []string{"name", "url", "description"}, []interface{}{_name, _url, _description}, "key = "+_key)

	log.Println(strQuery)
	// 서버 정보 변경
	_, err := baseDB.Exec(strQuery)
	if err != nil {
		ReturnError(c, err)
		return
	}

	ret := map[string]interface{}{}
	ret["server_key"] = fmt.Sprint(_key)
	c.JSON(http.StatusOK, ret)
}

// DelServer ...
func DelServer(c *gin.Context) {
	baseDB := OpenDB()
	defer baseDB.Close()

	// key=3
	_key, _ := c.GetQuery("key")

	// 서버 정보 변경
	_, err := baseDB.Exec("DELETE FROM server WHERE key = $1", _key)
	if err != nil {
		ReturnError(c, err)
		return
	}

	ret := map[string]interface{}{}
	c.JSON(http.StatusOK, ret)
}
