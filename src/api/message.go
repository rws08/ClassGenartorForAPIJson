package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ReturnError ... error message
func ReturnError(c *gin.Context, err error) {
	ret := map[string]string{}
	ret["error_msg"] = err.Error()
	c.JSON(http.StatusBadRequest, ret)
}
