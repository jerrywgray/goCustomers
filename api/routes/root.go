package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jerrywgray/goCustomers/api/schemas"
	log "github.com/sirupsen/logrus"
)

func root(c *gin.Context) {
	results := sumResults{}
	db.Model(&schemas.Customers{}).Select("sum(count) as total").Scan(&results)
	log.Printf("sum: %d", results.Total)

	c.String(http.StatusOK, fmt.Sprintf(`Total entries: %d`, results.Total))
}
