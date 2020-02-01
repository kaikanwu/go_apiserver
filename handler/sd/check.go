package sd

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const(
	B = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func HealthCheck(c *gin.Context) {
	message := "KK OK"
	c.String(http.StatusOK, "\n" +message)
}