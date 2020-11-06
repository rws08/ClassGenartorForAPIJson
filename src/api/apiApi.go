package api

import (
	"github.com/gin-gonic/gin"
)

// ObjAPI ... api 구조
type ObjAPI struct {
	Key         int    `json:"key"`         // key
	ServerKey   int    `json:"server_key"`  // server_key
	Name        string `json:"name"`        // name
	URL         string `json:"url"`         // url
	Description string `json:"description"` // description
}

// CreateAPI ...
func CreateAPI(c *gin.Context) {
}

// ReadAPI ...
func ReadAPI(c *gin.Context) {

}

// UpdateAPI ...
func UpdateAPI(c *gin.Context) {

}

// DelAPI ...
func DelAPI(c *gin.Context) {

}
