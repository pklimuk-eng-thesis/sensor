package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PresenceSensorHandler struct {
	service PresenceSensorService
}

func NewPresenceSensorHandler(service PresenceSensorService) *PresenceSensorHandler {
	return &PresenceSensorHandler{service}
}

func (h *PresenceSensorHandler) IsPresenceDetected(c *gin.Context) {
	presenceDetected, err := h.service.IsPresenceDetected()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &presenceDetected)
}

func (h *PresenceSensorHandler) IsSensorEnabled(c *gin.Context) {
	sensorEnabled := h.service.IsSensorEnabled()
	c.IndentedJSON(http.StatusOK, &sensorEnabled)
}

func (h *PresenceSensorHandler) TogglePresenceDetected(c *gin.Context) {
	presenceDetected, err := h.service.TogglePresenceDetected()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, &presenceDetected)
}

func (h *PresenceSensorHandler) ToggleSensorEnabled(c *gin.Context) {
	sensorEnabled := h.service.ToggleSensorEnabled()
	c.IndentedJSON(http.StatusOK, &sensorEnabled)
}
