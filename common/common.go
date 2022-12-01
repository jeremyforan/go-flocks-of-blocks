package common

type optionalSetStates map[string]bool

// ButtonStyle these relate to the three
type ColorSchema string

const (
	StyleDefault ColorSchema = "default"
	StylePrimary ColorSchema = "primary"
	StyleDanger  ColorSchema = "danger"
)
