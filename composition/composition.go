package composition

type Composition interface {
	compositeRender() string
}

type CompositionType string

// stringer
func (c CompositionType) String() string {
	return string(c)
}

const (
	PlainText CompositionType = "plain_text"
	Mrkdwn    CompositionType = "mrkdwn"
)
