package domain

type Sensor struct {
	Detected bool `json:"detected"`
	Enabled  bool `json:"enabled"`
}
