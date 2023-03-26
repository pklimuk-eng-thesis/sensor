package http

import "github.com/gin-gonic/gin"

var detectedEndpoint = "/detected"
var enabledEndpoint = "/enabled"

func SetupRouter(r *gin.Engine, h *SensorHandler) {
	r.GET(detectedEndpoint, h.Detected)
	r.GET(enabledEndpoint, h.IsSensorEnabled)
	r.POST(detectedEndpoint, h.ToggleDetected)
	r.POST(enabledEndpoint, h.ToggleSensorEnabled)
}
