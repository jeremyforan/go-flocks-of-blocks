package common

type optionalSetStates map[string]bool

// ButtonStyle these relate to the three
type ColorSchema string

// stringer
func (c ColorSchema) String() string {
	return string(c)
}

const (
	StyleDefault ColorSchema = "default"
	StylePrimary ColorSchema = "primary"
	StyleDanger  ColorSchema = "danger"
)

func RemoveDuplicateString(strSlice []string) []string {
	allKeys := make(map[string]bool)
	var list []string
	for _, item := range strSlice {
		item := string(item)
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
