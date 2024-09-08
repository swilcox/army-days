package days

import "time"

type Config struct {
	UseArmyButtDays bool `json:"useArmyButtDays,omitempty" yaml:"useArmyButtDays,omitempty"`
	ShowCompleted   bool `json:"showCompleted,omitempty" yaml:"showCompleted,omitempty"`
}

// NewConfig creates a new Config with default values
func NewConfig() Config {
	return Config{
		UseArmyButtDays: false,
		ShowCompleted:   false,
	}
}

// UnmarshalYAML implements the yaml.Unmarshaler interface
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type Alias Config
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	*c = NewConfig() // Set defaults
	return unmarshal(aux)
}

type Entry struct {
	Title string `json:"title" yaml:"title"`
	Date  string `json:"date" yaml:"date"`
}

type ConfigData struct {
	Config  Config  `json:"config" yaml:"config"`
	Entries []Entry `json:"entries" yaml:"entries"`
}

type OutputData struct {
	Title string    `json:"title" yaml:"title"`
	Date  time.Time `json:"date" yaml:"date"`
	Days  float32   `json:"days" yaml:"days"`
}
