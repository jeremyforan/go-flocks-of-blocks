package composition

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"net/url"
)

type Option struct {
	// Required
	text  CompositionText
	value string

	// Optional
	description CompositionText
	url         url.URL

	optionals optional
}

type optional struct {
	Description bool
	Url         bool
}

type optionOption func(*Option)

// NewOption creates a new option.
func NewOption(text string, value string) Option {
	return Option{
		text:  NewPlainText(text),
		value: value,
	}
}

func (o *Option) setDescription(description CompositionText) {
	o.description = description
	o.optionals.Description = true
}

func (o *Option) setUrl(url url.URL) {
	o.url = url
	o.optionals.Url = true
}

func (o Option) SetDescription(description CompositionText) Option {
	o.setDescription(description)
	return o
}

func (o *Option) SetUrl(url url.URL) *Option {
	o.setUrl(url)
	return o
}

func (o *Option) RemoveDescription() *Option {
	o.optionals.Description = false
	return o
}

func (o *Option) RemoveUrl() *Option {
	o.optionals.Url = false
	return o
}

// optionAbstraction is used to render the option
type optionAbstraction struct {
	Text        CompositionText
	Value       string
	Description CompositionText
	Url         string
	Optionals   optional
}

// create an option obstraction for rendering
func (o Option) abstraction() optionAbstraction {
	url := ""
	if o.optionals.Url {
		url = o.url.String()
	}
	return optionAbstraction{
		Text:        o.text,
		Value:       o.value,
		Description: o.description,
		Url:         url,
		Optionals:   o.optionals,
	}
}

func (o optionAbstraction) Template() string {
	return `{
	"text": {{ .Text.Render}},
	"value": "{{.Value}}"{{if .Optionals.Description}},
	"description": {{.Description.Render}}{{end}}{{if .Optionals.Url}},
	"url": "{{.Url}}"{{end}}	
}`
}

// Render renders the option to a string.
func (o Option) Render() string {
	return common.Render(o.abstraction())
}
