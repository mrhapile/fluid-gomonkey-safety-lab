package unsafe

// GlobalConfig represents a critical system configuration that is globally accessible.
// This simulated dependency mimics Fluid's reliance on global runtime state.
var GlobalConfig = "default"

// GetConfig returns the current global configuration.
func GetConfig() string {
	return GlobalConfig
}
