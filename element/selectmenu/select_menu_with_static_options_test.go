package selectmenu

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/option"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/optiongroup"
	"testing"
)

const (
	validMenuWithInitialAndPlaceholder = `{
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
)

func TestNewSelectMenuWithStaticOptions(t *testing.T) {
	t.Run("NewMultiSelectMenuWithStaticOptions", func(t *testing.T) {
		menu := NewSelectMenuWithStaticOptions("text1234")

		opt := option.NewOption("Wait for it", "value-0")
		opt2 := option.NewOption("Initial", "value-1")

		optionGroup := optiongroup.NewOptionGroup("Group 1").AddOption(opt).AddOption(opt2)

		menu = menu.AddOptionGroup(optionGroup)

		// convert to section
		section := menu.Section().Render()
		if section == validMenuWithInitialAndPlaceholder {
			t.Errorf("Rendered section does not match expected JSON: %s", section)
		}

		fmt.Println(section)
	})
}
