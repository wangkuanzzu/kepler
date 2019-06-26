package router

import (
	"github.com/QOSGroup/kepler/server/handler/key"
	"github.com/QOSGroup/kepler/server/handler/qcp"
	"github.com/QOSGroup/kepler/server/handler/qsc"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.Use(func(c *gin.Context) {

		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, HEAD, OPTIONS, PUT, PATCH, DELETE")

		if c.Request.Method != "OPTIONS" {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusOK)
		}
	})
	key.Register(r)
	qcp.Register(r)
	qsc.Register(r)
}
