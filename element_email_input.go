package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
)

type EmailInput struct {
	slackType ElementType
	actionId  string

	initialEmail         string
	dispatchActionConfig composition.DispatchActionConfig
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
		slackType:            EmailInputElement,
		actionId:             actionId,
		dispatchActionConfig: composition.NewDispatchActionConfig(),

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
