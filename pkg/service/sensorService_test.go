package service

import (
	"testing"

	"github.com/pklimuk-eng-thesis/sensor/pkg/domain"
)

var sensorEnabledDetected = &domain.Sensor{Enabled: true, Detected: true}
var sensorEnabledNotDetected = &domain.Sensor{Enabled: true, Detected: false}
var sensorDisabledNotDetected = &domain.Sensor{Enabled: false, Detected: false}

func Test_SensorService_GetInfo(t *testing.T) {
	tests := []struct {
		name string
		s    sensorService
		want *domain.Sensor
	}{
		{name: "sensor enabled, detected",
			s:    sensorService{sensor: sensorEnabledDetected},
			want: &domain.Sensor{Enabled: true, Detected: true}},
		{name: "sensor enabled, not detected",
			s:    sensorService{sensor: sensorEnabledNotDetected},
			want: &domain.Sensor{Enabled: true, Detected: false}},
		{name: "sensor disabled, not detected",
			s:    sensorService{sensor: sensorDisabledNotDetected},
			want: &domain.Sensor{Enabled: false, Detected: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.GetInfo()
			if got.Detected != tt.want.Detected || got.Enabled != tt.want.Enabled {
				t.Errorf("sensorService.Detected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sensorService_ToggleDetected(t *testing.T) {
	tests := []struct {
		name    string
		s       sensorService
		want    *domain.Sensor
		wantErr bool
	}{
		{name: "sensor enabled, detected",
			s:       sensorService{sensor: sensorEnabledDetected},
			want:    &domain.Sensor{Enabled: true, Detected: false},
			wantErr: false},
		{name: "sensor enabled, not detected",
			s:       sensorService{sensor: sensorEnabledNotDetected},
			want:    &domain.Sensor{Enabled: true, Detected: true},
			wantErr: false},
		{name: "sensor disabled, not detected",
			s:       sensorService{sensor: sensorDisabledNotDetected},
			want:    &domain.Sensor{Enabled: false, Detected: false},
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToggleDetected()
			if (err != nil) != tt.wantErr {
				t.Errorf("sensorService.ToggleDetected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.Detected != tt.want.Detected || got.Enabled != tt.want.Enabled {
				t.Errorf("sensorService.ToggleDetected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sensorService_ToggleSensorEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    sensorService
		want *domain.Sensor
	}{
		{name: "sensor enabled, detected",
			s:    sensorService{sensor: sensorEnabledDetected},
			want: &domain.Sensor{Enabled: false, Detected: false}},
		{name: "sensor enabled, not detected",
			s:    sensorService{sensor: sensorEnabledNotDetected},
			want: &domain.Sensor{Enabled: false, Detected: false}},
		{name: "sensor disabled, not detected",
			s:    sensorService{sensor: sensorDisabledNotDetected},
			want: &domain.Sensor{Enabled: true, Detected: false}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.s.ToggleSensorEnabled()
			if got.Detected != tt.want.Detected || got.Enabled != tt.want.Enabled {
				t.Errorf("sensorService.ToggleSensorEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
