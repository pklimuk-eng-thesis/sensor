package pkg

import "github.com/gin-gonic/gin"

func SetupRouter(r *gin.Engine, h *PresenceSensorHandler) {
	r.GET("/presenceDetected", h.IsPresenceDetected)
	r.GET("/sensorEnabled", h.IsSensorEnabled)
	r.POST("/presenceDetected", h.TogglePresenceDetected)
	r.POST("/sensorEnabled", h.ToggleSensorEnabled)
}
