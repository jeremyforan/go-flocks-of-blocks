package button

type buttonOptionals struct {
	Url                bool
	Value              bool
	Style              bool
	Confirm            bool
	AccessibilityLabel bool
}

// buttonConstructionOptions allows for optional parameters to be passed into the NewButton function.
type buttonConstructionOptions func(*Button)

////////////////////////////////////////////////////////////////////////////////////
// Button Abstraction

func (b *Button) abstraction() buttonAbstraction {
	url := ""
	if b.optionals.Url {
		url = b.url.String()
	}
	return buttonAbstraction{
		Type:               b.slackType.String(),
		Text:               b.text,
		ActionId:           b.actionId,
		Url:                url,
		Value:              b.value,
		Style:              b.style.String(),
		Confirm:            b.confirm,
		AccessibilityLabel: b.accessibilityLabel,
		Optionals:          b.optionals,
	}
}
