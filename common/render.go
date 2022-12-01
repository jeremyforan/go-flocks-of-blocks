package common

import (
	"bytes"
	"encoding/json"
	"text/template"
)

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
