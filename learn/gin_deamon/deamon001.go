package gin_deamon

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWordGet(c *gin.Context)  {
	c.String(http.StatusOK, "rest api/v1/x")
}
func main() {
	router := gin.Default()

	router.Get("/api/v1/x", helloWordGet)

	router.run("127.0.0.1:8081")
}