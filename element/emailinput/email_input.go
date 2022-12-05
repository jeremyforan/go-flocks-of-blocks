package emailinput

import (
	"github.com/jeremyforan/go-flocks-of-blocks/composition/dispatchactionconfig"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
)

type EmailInput struct {
	slackType element.ElementType
	actionId  string

	initialEmail         string
	dispatchActionConfig dispatchactionconfig.DispatchActionConfig
	focusOnLoad          bool
	placeholder          string

	options EmailInputOptions
}

type EmailInputOptions struct {
	InitialEmail         bool
	DispatchActionConfig bool
	FocusOnLoad          bool
	Placeholder          bool
}

func NewEmailInput(actionId string) EmailInput {
	return EmailInput{
		slackType:            element.EmailInput,
		actionId:             actionId,
		dispatchActionConfig: dispatchactionconfig.NewDispatchActionConfig(),

		options: EmailInputOptions{
			InitialEmail:         false,
			DispatchActionConfig: false,
			FocusOnLoad:          false,
			Placeholder:          false,
		},
	}
}

// SetInitialEmail sets the initial email for the email input.
func (e *EmailInput) setInitialEmail(initialEmail string) {
	e.initialEmail = initialEmail
	e.options.InitialEmail = true
}

// RemoveInitialEmail removes the initial email for the email input.
func (e *EmailInput) removeInitialEmail() {
	e.options.InitialEmail = false
}

// todo: email input not implemented yet
