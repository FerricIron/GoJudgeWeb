package authorize

import "github.com/gin-gonic/gin"

type loginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
func Login(c *gin.Context)  {
	c.JSON(200,gin.H{
		"status":"test",
		"message":"test",
	})
}
