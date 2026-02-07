package safe

// GlobalConfig represents a critical system configuration that is globally accessible.
// This mirrors the unsafe package but will be tested correctly.
var GlobalConfig = "default"

// GetConfig returns the current global configuration.
func GetConfig() string {
	return GlobalConfig
}
