package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/sensor/pkg/domain"
	"github.com/pklimuk-eng-thesis/sensor/pkg/service"
	"github.com/stretchr/testify/assert"
)

func TestSensorHandler_GetInfo(t *testing.T) {
	// Create a new mock sensor service and set the expected behavior
	sensorService := new(service.MockSensorService)
	expectedSensor := &domain.Sensor{Enabled: true, Detected: true}
	sensorService.EXPECT().GetInfo().Return(expectedSensor)

	// Create a new sensor handler with the mock service
	handler := NewSensorHandler(sensorService)

	// Create a new Gin context and execute the handler
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.GetInfo(c)

	// Assert that the response status code is HTTP 200 OK and the JSON response body matches the expected sensor
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled":true,"detected":true}`, w.Body.String())
}

func TestSensorHandler_ToggleDetected(t *testing.T) {
	sensorService := new(service.MockSensorService)
	expectedSensor := &domain.Sensor{Enabled: true, Detected: false}
	sensorService.EXPECT().ToggleDetected().Return(expectedSensor, nil)

	handler := NewSensorHandler(sensorService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.ToggleDetected(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled":true,"detected":false}`, w.Body.String())
}

func TestSensorHandler_ToggleDetected_Error(t *testing.T) {
	sensorService := new(service.MockSensorService)
	sensorService.EXPECT().ToggleDetected().Return(nil, service.ErrSensorIsDisabled)

	handler := NewSensorHandler(sensorService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.ToggleDetected(c)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, service.ErrSensorIsDisabled.Error(), w.Body.String())
}

func TestSensorHandler_ToggleSensorEnabled(t *testing.T) {
	sensorService := new(service.MockSensorService)
	expectedSensor := &domain.Sensor{Enabled: false, Detected: false}
	sensorService.EXPECT().ToggleSensorEnabled().Return(expectedSensor)

	handler := NewSensorHandler(sensorService)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	handler.ToggleSensorEnabled(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"enabled":false,"detected":false}`, w.Body.String())
}
