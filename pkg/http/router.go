package http

import "github.com/gin-gonic/gin"

var detectedEndpoint = "/detected"
var enabledEndpoint = "/enabled"
var getInfoEndpoint = "/info"

func SetupRouter(r *gin.Engine, h *SensorHandler) {
	r.GET(getInfoEndpoint, h.GetInfo)
	r.PATCH(detectedEndpoint, h.ToggleDetected)
	r.PATCH(enabledEndpoint, h.ToggleSensorEnabled)
}
