package resource

import (
	"stand_up/internal/adapter/api/requests"
	"stand_up/internal/core/logger"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) StoreTask(c *gin.Context) {
	logger.Info("Get Store sprint called")

	request := requests.TaskRequest{}
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	resp, err := s.taskService.Save(request)
	if err != nil {
		logger.Error("Error saving sprint" + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resp})
}

func (s *HTTPHandler) GetAllTask(c *gin.Context) {
	logger.Info("Get tasks called")

	params := FlatUrlQuery(c.Request.URL.Query())
	resp, err := s.taskService.GetAll(params)

	if err != nil {
		logger.Error("Error getting sprint" + err.Error())
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": resp})
}
