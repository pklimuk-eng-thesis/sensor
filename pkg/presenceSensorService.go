package pkg

import (
	"errors"
)

var ErrSensorIsDisabled = errors.New("Sensor is disabled")

type PresenceSensorService interface {
	IsPresenceDetected() (bool, error)
	IsSensorEnabled() bool
	TogglePresenceDetected() (bool, error)
	ToggleSensorEnabled() bool
}

type presenceSensorService struct {
	presenceSensor *PresenceSensor
}

func NewPresenceSensorService(sensor *PresenceSensor) PresenceSensorService {
	return &presenceSensorService{
		presenceSensor: sensor,
	}
}

func (s *presenceSensorService) IsPresenceDetected() (bool, error) {
	if !s.presenceSensor.SensorEnabled {
		return false, ErrSensorIsDisabled
	}
	return s.presenceSensor.PresenceDetected, nil
}

func (s *presenceSensorService) IsSensorEnabled() bool {
	return s.presenceSensor.SensorEnabled
}

func (s *presenceSensorService) TogglePresenceDetected() (bool, error) {
	if !s.presenceSensor.SensorEnabled {
		return false, ErrSensorIsDisabled
	}
	s.presenceSensor.PresenceDetected = !s.presenceSensor.PresenceDetected
	return s.presenceSensor.PresenceDetected, nil
}

func (s *presenceSensorService) ToggleSensorEnabled() bool {
	s.presenceSensor.SensorEnabled = !s.presenceSensor.SensorEnabled
	return s.presenceSensor.SensorEnabled
}
