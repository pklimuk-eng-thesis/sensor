package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	sService "github.com/pklimuk-eng-thesis/sensor/pkg/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockService struct {
	mock.Mock
}

func (m *mockService) Detected() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *mockService) IsSensorEnabled() bool {
	args := m.Called()
	return args.Bool(0)
}

func (m *mockService) ToggleDetected() (bool, error) {
	args := m.Called()
	return args.Bool(0), args.Error(1)
}

func (m *mockService) ToggleSensorEnabled() bool {
	args := m.Called()
	return args.Bool(0)
}

func TestSensorHandler_Detected_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewSensorHandler(mockSvc)

	r := gin.Default()
	r.GET(detectedEndpoint, handler.Detected)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("Detected").Return(true, nil)
		req, _ := http.NewRequest(http.MethodGet, detectedEndpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}

func TestSensorHandler_Detected_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewSensorHandler(mockSvc)

	r := gin.Default()
	r.GET(detectedEndpoint, handler.Detected)

	t.Run("service error", func(t *testing.T) {
		mockSvc.On("Detected").Return(false, sService.ErrSensorIsDisabled)
		req, _ := http.NewRequest(http.MethodGet, detectedEndpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, sService.ErrSensorIsDisabled.Error(), w.Body.String())
	})
}

func TestSensorHandler_IsSensorEnabled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewSensorHandler(mockSvc)

	r := gin.Default()
	r.GET(enabledEndpoint, handler.IsSensorEnabled)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("IsSensorEnabled").Return(true)
		req, _ := http.NewRequest(http.MethodGet, enabledEndpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}

func TestSensorHandler_ToggleDetected_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewSensorHandler(mockSvc)

	r := gin.Default()
	r.POST(detectedEndpoint, handler.ToggleDetected)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("ToggleDetected").Return(true, nil)
		req, _ := http.NewRequest(http.MethodPost, detectedEndpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}

func TestSensorHandler_ToggleDetected_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewSensorHandler(mockSvc)

	r := gin.Default()
	r.POST(detectedEndpoint, handler.ToggleDetected)

	t.Run("service error", func(t *testing.T) {
		mockSvc.On("ToggleDetected").Return(false, sService.ErrSensorIsDisabled)
		req, _ := http.NewRequest(http.MethodPost, detectedEndpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, sService.ErrSensorIsDisabled.Error(), w.Body.String())
	})
}

func TestSensorHandler_ToggleSensorEnabled(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockSvc := new(mockService)
	handler := NewSensorHandler(mockSvc)

	r := gin.Default()
	r.POST(enabledEndpoint, handler.ToggleSensorEnabled)

	t.Run("success", func(t *testing.T) {
		mockSvc.On("ToggleSensorEnabled").Return(true)
		req, _ := http.NewRequest(http.MethodPost, enabledEndpoint, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "true", w.Body.String())
	})
}
