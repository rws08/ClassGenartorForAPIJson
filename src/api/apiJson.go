package api

import (
	"github.com/gin-gonic/gin"
)

// ObjJSON ... josn 구조
type ObjJSON struct {
	Key         int    `json:"key"`          // key
	APIKey      int    `json:"api_key"`      // api_key
	Successed   bool   `json:"successed"`    // 성공여부
	ParamData   string `json:"param_data"`   // 요청 파라미터 정보
	Data        string `json:"data"`         // 응답 데이터
	Description string `json:"description"`  // description
	CreatedDate string `json:"created_date"` // created_date
}

// CreateJSON ...
func CreateJSON(c *gin.Context) {
}

// ReadJSON ...
func ReadJSON(c *gin.Context) {

}

// UpdateJSON ...
func UpdateJSON(c *gin.Context) {

}

// DelJSON ...
func DelJSON(c *gin.Context) {

}
