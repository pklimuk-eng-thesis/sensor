package pkg

import (
	"testing"

	"github.com/pklimuk-eng-thesis/sensor/pkg/domain"
)

var sensorEnabledDetected = &domain.Sensor{Enabled: true, Detected: true}
var sensorEnabledNotDetected = &domain.Sensor{Enabled: true, Detected: false}
var sensorDisabledDetected = &domain.Sensor{Enabled: false, Detected: true}
var sensorDisabledNotDetected = &domain.Sensor{Enabled: false, Detected: false}

func Test_SensorService_Detected(t *testing.T) {
	tests := []struct {
		name    string
		s       *sensorService
		want    bool
		wantErr bool
	}{
		{name: "sensor enabled, detected",
			s:       &sensorService{sensor: sensorEnabledDetected},
			want:    true,
			wantErr: false},
		{name: "sensor enabled, not detected",
			s:       &sensorService{sensor: sensorEnabledNotDetected},
			want:    false,
			wantErr: false},
		{name: "sensor disabled, not detected",
			s:       &sensorService{sensor: sensorDisabledNotDetected},
			want:    false,
			wantErr: true},
		{name: "sensor disabled, detected",
			s:       &sensorService{sensor: sensorDisabledDetected},
			want:    false,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Detected()
			if (err != nil) != tt.wantErr {
				t.Errorf("sensorService.Detected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sensorService.Detected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sensorService_IsSensorEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    *sensorService
		want bool
	}{
		{name: "sensor enabled, detected",
			s:    &sensorService{sensor: sensorEnabledDetected},
			want: true},
		{name: "sensor enabled, not detected",
			s:    &sensorService{sensor: sensorEnabledNotDetected},
			want: true},
		{name: "sensor disabled, not detected",
			s:    &sensorService{sensor: sensorDisabledNotDetected},
			want: false},
		{name: "sensor disabled, detected",
			s:    &sensorService{sensor: sensorDisabledDetected},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSensorEnabled(); got != tt.want {
				t.Errorf("sensorService.IsSensorEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sensorService_ToggleDetected(t *testing.T) {
	tests := []struct {
		name    string
		s       *sensorService
		want    bool
		wantErr bool
	}{
		{name: "sensor enabled, detected",
			s:       &sensorService{sensor: sensorEnabledDetected},
			want:    false,
			wantErr: false},
		{name: "sensor enabled, not detected",
			s:       &sensorService{sensor: sensorEnabledNotDetected},
			want:    true,
			wantErr: false},
		{name: "sensor disabled, not detected",
			s:       &sensorService{sensor: sensorDisabledNotDetected},
			want:    false,
			wantErr: true},
		{name: "sensor disabled, detected",
			s:       &sensorService{sensor: sensorDisabledDetected},
			want:    false,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.ToggleDetected()
			if (err != nil) != tt.wantErr {
				t.Errorf("sensorService.ToggleDetected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("sensorService.ToggleDetected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sensorService_ToggleSensorEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    *sensorService
		want bool
	}{
		{name: "sensor enabled, detected",
			s:    &sensorService{sensor: sensorEnabledDetected},
			want: false},
		{name: "sensor enabled, not detected",
			s:    &sensorService{sensor: sensorEnabledNotDetected},
			want: false},
		{name: "sensor disabled, not detected",
			s:    &sensorService{sensor: sensorDisabledNotDetected},
			want: true},
		{name: "sensor disabled, detected",
			s:    &sensorService{sensor: sensorDisabledDetected},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToggleSensorEnabled(); got != tt.want {
				t.Errorf("sensorService.ToggleSensorEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
