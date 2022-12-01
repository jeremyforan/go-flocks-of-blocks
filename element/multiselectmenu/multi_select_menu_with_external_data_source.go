package multiselectmenu

import (
	"go-flocks-of-blocks/composition/compositiontext"
	"go-flocks-of-blocks/composition/confirmationdialog"
	"go-flocks-of-blocks/composition/option"
	"go-flocks-of-blocks/element"
)

type MultiSelectMenuWithExternalDataSource struct {
	slackType element.ElementType
	actionID  string
	options   []option.Option

	minQueryLength   int
	initialOptions   []option.Option
	confirm          confirmationdialog.ConfirmationDialog
	maxSelectedItems int
	focusOnLoad      bool
	placeholder      compositiontext.CompositionText

	optionals multiSelectMenuWithStaticOptionsOptions
}

type MultiSelectMenuWithExternalDataSourceOptions struct {
	MinQueryLength   bool
	InitialOptions   bool
	Confirm          bool
	MaxSelectedItems bool
	FocusOnLoad      bool
	Placeholder      bool
}

func NewMultiSelectMenuWithExternalDataSource(actionId string) MultiSelectMenuWithExternalDataSource {
	return MultiSelectMenuWithExternalDataSource{
		slackType: element.MultiSelectMenuWithExternalDataSource,
		actionID:  actionId,
		options:   []option.Option{},
		optionals: multiSelectMenuWithStaticOptionsOptions{
			OptionGroups:     false,
			InitialOptions:   false,
			Confirm:          false,
			MaxSelectedItems: false,
			FocusOnLoad:      false,
			Placeholder:      false,
		},
	}
}
