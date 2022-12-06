package composition

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
)

type CompositionText struct {
	slackType CompositionType
	text      string
	emoji     bool
	verbatim  bool
}

// NewPlainText Create a new markdown object
func NewPlainText(text string) CompositionText {
	return newText("plain_text", text)
}

// NewMrkdwnText Create a new markdown compositiontext object
func NewMrkdwnText(text string) CompositionText {
	return newText("mrkdwn", text)
}

func newText(slackType CompositionType, text string) CompositionText {
	return CompositionText{
		slackType: slackType,
		text:      text,
	}
}

// SetEmoji set the emoji flag
func (m *CompositionText) setEmoji(emoji bool) {
	m.emoji = emoji
}

// EnableEmoji enable the emoji flag
func (m CompositionText) EnableEmoji() CompositionText {
	m.setEmoji(true)
	return m
}

// SetVerbatim set the verbatim flag
func (m *CompositionText) SetVerbatim(verbatim bool) {
	m.verbatim = verbatim
}

// Stringer that returns the compositiontext
func (m CompositionText) String() string {
	return m.text
}

// compositionTextAbstraction is used to render the composition compositiontext
type compositionTextAbstraction struct {
	Type     string
	Text     string
	Emoji    bool
	Verbatim bool
}

// abstraction for the text
func (m CompositionText) abstraction() compositionTextAbstraction {
	return compositionTextAbstraction{
		Type:     m.slackType.String(),
		Text:     m.text,
		Emoji:    m.emoji,
		Verbatim: m.verbatim,
	}
}

func (m compositionTextAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"text": "{{.Text}}"{{if .Emoji}},
	"emoji": {{.Emoji}}{{end}}{{if .Verbatim}}, 
	"verbatim": {{.Verbatim}}{{end}}
}`
}

// Render the compositiontext
func (m CompositionText) Render() string {
	return common.Render(m.abstraction())
}
