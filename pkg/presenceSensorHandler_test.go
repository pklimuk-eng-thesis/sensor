package pkg

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) IsPresenceDetected() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *mockService) IsSensorEnabled() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *mockService) TogglePresenceDetected() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *mockService) ToggleSensorEnabled() bool {
	args := m.Called()
	return args.Bool(0)
}

func TestPresenceSensorHandler_IsPresenceDetected_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewPresenceSensorHandler(mockSvc)

	r := gin.Default()
	r.GET("/presenceDetected", handler.IsPresenceDetected)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("IsPresenceDetected").Return(true, nil)
		req, _ := http.NewRequest(http.MethodGet, "/presenceDetected", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}

func TestPresenceSensorHandler_IsPresenceDetected_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewPresenceSensorHandler(mockSvc)

	r := gin.Default()
	r.GET("/presenceDetected", handler.IsPresenceDetected)

	t.Run("service error", func(t *testing.T) {
		mockSvc.On("IsPresenceDetected").Return(false, ErrSensorIsDisabled)
		req, _ := http.NewRequest(http.MethodGet, "/presenceDetected", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, ErrSensorIsDisabled.Error(), w.Body.String())
	})
}

func TestPresenceSensorHandler_IsSensorEnabled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewPresenceSensorHandler(mockSvc)

	r := gin.Default()
	r.GET("/sensorEnabled", handler.IsSensorEnabled)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("IsSensorEnabled").Return(true)
		req, _ := http.NewRequest(http.MethodGet, "/sensorEnabled", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}

func TestPresenceSensorHandler_TogglePresenceDetected_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewPresenceSensorHandler(mockSvc)

	r := gin.Default()
	r.POST("/presenceDetected", handler.TogglePresenceDetected)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("TogglePresenceDetected").Return(true, nil)
		req, _ := http.NewRequest(http.MethodPost, "/presenceDetected", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}

func TestPresenceSensorHandler_TogglePresenceDetected_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewPresenceSensorHandler(mockSvc)

	r := gin.Default()
	r.POST("/presenceDetected", handler.TogglePresenceDetected)

	t.Run("service error", func(t *testing.T) {
		mockSvc.On("TogglePresenceDetected").Return(false, ErrSensorIsDisabled)
		req, _ := http.NewRequest(http.MethodPost, "/presenceDetected", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, ErrSensorIsDisabled.Error(), w.Body.String())
	})
}

func TestPresenceSensorHandler_ToggleSensorEnabled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewPresenceSensorHandler(mockSvc)

	r := gin.Default()
	r.POST("/sensorEnabled", handler.ToggleSensorEnabled)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("ToggleSensorEnabled").Return(true)
		req, _ := http.NewRequest(http.MethodPost, "/sensorEnabled", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}
