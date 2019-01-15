package routers

import (
	"github.com/gin-gonic/gin"
)
func SetRouters(r *gin.Engine) (){
	SetApiRouters(r.Group("/api"))
	SetWebRouters(r.Group("/web"))
	setAuthorizeRouters(r.Group("/authorize"))

}
