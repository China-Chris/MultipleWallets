package handle

import "github.com/gin-gonic/gin"

func GetPing(c *gin.Context) {
	c.String(200, "pong")
}
