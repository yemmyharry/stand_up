package resource

import (
	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) Routes(router *gin.Engine) {
	apirouter := router.Group("api/v1")
	apirouter.GET("/healthcheck", s.HealthCheck)
	apirouter.POST("/standup/update", s.StoreStandupUpdate)
	apirouter.GET("/standup/update", s.StoreStandupUpdate)
	apirouter.POST("/task", s.StoreTask)
	apirouter.GET("/task", s.GetAllTask)
	apirouter.POST("/sprint", s.StoreSprint)
	apirouter.GET("/sprint", s.GetAllSprint)
	router.NoRoute(func(c *gin.Context) { c.JSON(404, "no route") })
}
