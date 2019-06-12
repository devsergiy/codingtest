package subscriptions

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spetrunin/codingtest/models"
)

func SubscribeHandler(c *gin.Context) {
	var subscription models.Subscription

	data, _ := c.GetRawData()
	fmt.Println(string(data))

	err := json.Unmarshal(data, &subscription)
	if err != nil {
		log.Output(1, err.Error())
	}

	if ok, errors := subscription.Valid(); !ok {
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	c.JSON(http.StatusOK, subscription)
}
