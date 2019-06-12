package plans

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spetrunin/codingtest/config"
	"github.com/spetrunin/codingtest/models"
)

func ListHandler(c *gin.Context) {
	currency := c.Query("currency")

	if currency == "" {
		c.JSON(http.StatusOK, config.DefaultPlans)
		return
	}

	plans, err := updatePricing(currency, config.DefaultPlans)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, plans)
}

func updatePricing(cur string, plans []models.Plan) (out []models.Plan, err error) {
	currency := strings.ToUpper(cur)
	value, ok := config.Rates.Rates[currency]
	if !ok {
		return nil, errors.New("currency is not available")
	}

	for _, plan := range plans {
		plan.Price = plan.Price * value
		out = append(out, plan)
	}

	return
}
