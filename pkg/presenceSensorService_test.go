package pkg

import (
	"testing"
)

var sensorEnabledDetected = &PresenceSensor{SensorEnabled: true, PresenceDetected: true}
var sensorEnabledNotDetected = &PresenceSensor{SensorEnabled: true, PresenceDetected: false}
var sensorDisabledDetected = &PresenceSensor{SensorEnabled: false, PresenceDetected: true}
var sensorDisabledNotDetected = &PresenceSensor{SensorEnabled: false, PresenceDetected: false}

func Test_presenceSensorService_IsPresenceDetected(t *testing.T) {
	tests := []struct {
		name    string
		s       *presenceSensorService
		want    bool
		wantErr bool
	}{
		{name: "sensor enabled, presence detected",
			s:       &presenceSensorService{presenceSensor: sensorEnabledDetected},
			want:    true,
			wantErr: false},
		{name: "sensor enabled, presence not detected",
			s:       &presenceSensorService{presenceSensor: sensorEnabledNotDetected},
			want:    false,
			wantErr: false},
		{name: "sensor disabled, presence not detected",
			s:       &presenceSensorService{presenceSensor: sensorDisabledNotDetected},
			want:    false,
			wantErr: true},
		{name: "sensor disabled, presence detected",
			s:       &presenceSensorService{presenceSensor: sensorDisabledDetected},
			want:    false,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.IsPresenceDetected()
			if (err != nil) != tt.wantErr {
				t.Errorf("presenceSensorService.IsPresenceDetected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("presenceSensorService.IsPresenceDetected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_presenceSensorService_IsSensorEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    *presenceSensorService
		want bool
	}{
		{name: "sensor enabled, presence detected",
			s:    &presenceSensorService{presenceSensor: sensorEnabledDetected},
			want: true},
		{name: "sensor enabled, presence not detected",
			s:    &presenceSensorService{presenceSensor: sensorEnabledNotDetected},
			want: true},
		{name: "sensor disabled, presence not detected",
			s:    &presenceSensorService{presenceSensor: sensorDisabledNotDetected},
			want: false},
		{name: "sensor disabled, presence detected",
			s:    &presenceSensorService{presenceSensor: sensorDisabledDetected},
			want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.IsSensorEnabled(); got != tt.want {
				t.Errorf("presenceSensorService.IsSensorEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_presenceSensorService_TogglePresenceDetected(t *testing.T) {
	tests := []struct {
		name    string
		s       *presenceSensorService
		want    bool
		wantErr bool
	}{
		{name: "sensor enabled, presence detected",
			s:       &presenceSensorService{presenceSensor: sensorEnabledDetected},
			want:    false,
			wantErr: false},
		{name: "sensor enabled, presence not detected",
			s:       &presenceSensorService{presenceSensor: sensorEnabledNotDetected},
			want:    true,
			wantErr: false},
		{name: "sensor disabled, presence not detected",
			s:       &presenceSensorService{presenceSensor: sensorDisabledNotDetected},
			want:    false,
			wantErr: true},
		{name: "sensor disabled, presence detected",
			s:       &presenceSensorService{presenceSensor: sensorDisabledDetected},
			want:    false,
			wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.TogglePresenceDetected()
			if (err != nil) != tt.wantErr {
				t.Errorf("presenceSensorService.TogglePresenceDetected() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("presenceSensorService.TogglePresenceDetected() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_presenceSensorService_ToggleSensorEnabled(t *testing.T) {
	tests := []struct {
		name string
		s    *presenceSensorService
		want bool
	}{
		{name: "sensor enabled, presence detected",
			s:    &presenceSensorService{presenceSensor: sensorEnabledDetected},
			want: false},
		{name: "sensor enabled, presence not detected",
			s:    &presenceSensorService{presenceSensor: sensorEnabledNotDetected},
			want: false},
		{name: "sensor disabled, presence not detected",
			s:    &presenceSensorService{presenceSensor: sensorDisabledNotDetected},
			want: true},
		{name: "sensor disabled, presence detected",
			s:    &presenceSensorService{presenceSensor: sensorDisabledDetected},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ToggleSensorEnabled(); got != tt.want {
				t.Errorf("presenceSensorService.ToggleSensorEnabled() = %v, want %v", got, tt.want)
			}
		})
	}
}
