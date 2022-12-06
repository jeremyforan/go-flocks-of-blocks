package flocksofblocks

import (
	"bytes"
	"encoding/json"
	"text/template"
)

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

func removeDuplicateString(strSlice []string) []string {
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

type RenderableAbstraction interface {
	Template() string
}

func Render(asset RenderableAbstraction) string {
	var buff bytes.Buffer

	tmpl, err := template.New("").Parse(asset.Template())
	if err != nil {
		return ""
	}

	err = tmpl.Execute(&buff, asset)
	if err != nil {
		return ""
	}

	msg := buff.String()

	return msg
}

func Pretty(s string) string {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(s), "", "\t")
	if err != nil {
		return ""
	}

	return string(prettyJSON.Bytes())
}

func Valid(s string) bool {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, []byte(s), "", "\t")
	if err != nil {
		return false
	}
	return true
}
