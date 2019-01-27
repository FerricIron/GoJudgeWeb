package routers

import (
	"github.com/gin-gonic/gin"
)

func SetRouters(r *gin.Engine) {
	SetApiRouters(r.Group("/api"))
	SetDataRouters(r.Group("/data"))
	setAuthorizeRouters(r.Group("/authorize"))
}
