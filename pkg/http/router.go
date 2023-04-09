package http

import "github.com/gin-gonic/gin"

var detectedEndpoint = "/detected"
var enabledEndpoint = "/enabled"
var getInfoEndpoint = "/info"

func SetupRouter(r *gin.Engine, h *SensorHandler) {
	r.GET(getInfoEndpoint, h.GetInfo)
	r.POST(detectedEndpoint, h.ToggleDetected)
	r.POST(enabledEndpoint, h.ToggleSensorEnabled)
}
