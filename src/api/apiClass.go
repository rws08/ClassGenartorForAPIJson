package api

import (
	"github.com/gin-gonic/gin"
)

// ObjClass ... class 구조
type ObjClass struct {
	Key         int    `json:"key"`         // key
	APIKey      int    `json:"api_key"`     // api_key
	Name        string `json:"name"`        // name
	Description string `json:"description"` // description
}

// ObjVar ... var 구조
type ObjVar struct {
	Key         int    `json:"key"`          // key
	Name        string `json:"name"`         // name
	TypeKey     int    `json:"type_key"`     // type_key
	SubTypeKey  int    `json:"sub_type_key"` // sub_type_key
	Description string `json:"description"`  // description
}

// CreateClass ...
func CreateClass(c *gin.Context) {
}

// ReadClass ...
func ReadClass(c *gin.Context) {

}

// UpdateClass ...
func UpdateClass(c *gin.Context) {

}

// DelClass ...
func DelClass(c *gin.Context) {

}
