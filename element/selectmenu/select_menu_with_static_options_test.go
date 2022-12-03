package selectmenu

import (
	"fmt"
	"go-flocks-of-blocks/composition/option"
	"testing"
)

const (
	validMenuWithInitialandPlaceholder = `{
	"action_id": "text1234",
	"type": "multi_static_select",
	"options": [
		{
			"text": {
				"type": "plain_text",
				"text": "Wait for it"
			},
			"value": "value-0"
		},
		{
			"text": {
				"type": "plain_text",
				"text": "Initial"
			},
			"value": "value-1"
		}
	],
	"initial_options": [
		{
			"text": {
				"type": "plain_text",
				"text": "Initial"
			},
			"value": "value-1"
		}
	],
	"placeholder": {
		"type": "plain_text",
		"text": "Select items"
	}
}`
	validMenuInSection = ``
)

func TestNewMultiSelectMenuWithStaticOptions(t *testing.T) {
	t.Run("NewMultiSelectMenuWithStaticOptions", func(t *testing.T) {
		menu := NewSelectMenuWithStaticOptions("text1234")

		opt := option.NewOption("Wait for it", "value-0")
		opt2 := option.NewOption("Initial", "value-1")
		menu = menu.AddOption(opt).AddInitialOption(opt2).SetPlaceholder("Select items")

		output := menu.Render()
		if output != validMenuWithInitialandPlaceholder {
			t.Errorf("Rendered menu does not match expected JSON: %s", output)
		}

		// convert to section
		section := menu.Section().Render()
		if section == validMenuWithInitialandPlaceholder {
			t.Errorf("Rendered section does not match expected JSON: %s", section)
		}

		fmt.Println(section)
	})
}
