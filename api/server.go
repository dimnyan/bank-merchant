package api


import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"admin": "password",
	}))

	authorized.POST("admin", func(c *gin.Context) {
		// user := c.MustGet(gin.AuthUserKey).(string)
		var json struct {
			Value string `json:"value" binding:"required"` 
		}

		if c.Bind(&json) == nil {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

    return r
}

func Run() {
    r := SetupRouter()
    r.Run(":8081")
}