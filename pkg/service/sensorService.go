package service

import (
	"errors"

	"github.com/pklimuk-eng-thesis/sensor/pkg/domain"
)

var ErrSensorIsDisabled = errors.New("Sensor is disabled")

//go:generate --name SensorService --output mock_sensorService.go
type SensorService interface {
	GetInfo() *domain.Sensor
	ToggleDetected() (*domain.Sensor, error)
	ToggleSensorEnabled() *domain.Sensor
}

type sensorService struct {
	sensor *domain.Sensor
}

func NewSensorService(sensor *domain.Sensor) SensorService {
	return &sensorService{
		sensor: sensor,
	}
}

func (s *sensorService) GetInfo() *domain.Sensor {
	return s.sensor
}

func (s *sensorService) ToggleDetected() (*domain.Sensor, error) {
	if !s.sensor.Enabled {
		s.sensor.Detected = false
		return s.sensor, ErrSensorIsDisabled
	}
	s.sensor.Detected = !s.sensor.Detected
	return s.sensor, nil
}

func (s *sensorService) ToggleSensorEnabled() *domain.Sensor {
	s.sensor.Enabled = !s.sensor.Enabled
	if !s.sensor.Enabled {
		s.sensor.Detected = false
	}
	return s.sensor
}
