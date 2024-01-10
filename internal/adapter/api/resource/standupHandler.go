package resource

import (
	"stand_up/internal/adapter/api/requests"
	"stand_up/internal/core/logger"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) StoreStandupUpdate(c *gin.Context) {
	logger.Info("Get files called")

	request := requests.StandUpdateRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	resp, err := s.standupUpdate.Save(request)
	if err != nil {
		logger.Error("Error saving update" + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resp})
}

func (s *HTTPHandler) GetAllStandupUpdate(c *gin.Context) {
	logger.Info("Get update called")

	params := FlatUrlQuery(c.Request.URL.Query())
	resp, err := s.standupUpdate.GetAll(params)

	if err != nil {
		logger.Error("Error getting update" + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resp})
}
