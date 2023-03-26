package pkg

import (
	"errors"

	"github.com/pklimuk-eng-thesis/sensor/pkg/domain"
)

var ErrSensorIsDisabled = errors.New("Sensor is disabled")

type SensorService interface {
	Detected() (bool, error)
	IsSensorEnabled() bool
	ToggleDetected() (bool, error)
	ToggleSensorEnabled() bool
}

type sensorService struct {
	sensor *domain.Sensor
}

func NewSensorService(sensor *domain.Sensor) SensorService {
	return &sensorService{
		sensor: sensor,
	}
}

func (s *sensorService) Detected() (bool, error) {
	if !s.sensor.Enabled {
		return false, ErrSensorIsDisabled
	}
	return s.sensor.Detected, nil
}

func (s *sensorService) IsSensorEnabled() bool {
	return s.sensor.Enabled
}

func (s *sensorService) ToggleDetected() (bool, error) {
	if !s.sensor.Enabled {
		return false, ErrSensorIsDisabled
	}
	s.sensor.Detected = !s.sensor.Detected
	return s.sensor.Detected, nil
}

func (s *sensorService) ToggleSensorEnabled() bool {
	s.sensor.Enabled = !s.sensor.Enabled
	return s.sensor.Enabled
}
