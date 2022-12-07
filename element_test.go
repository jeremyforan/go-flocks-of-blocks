package flocksofblocks

import (
	"fmt"
	"net/url"
	"testing"
)

func TestButton(t *testing.T) {
	t.Run("Create Button Element", func(t *testing.T) {

		url, err := url.Parse("http://bing.com/search?q=dotnet")
		if err != nil {
			t.Error(err)
		}

		button := NewButton("Click This", "button1").MakeStyleDanger().AddUrl(url)
		button.AddUrl(url).MakeStyleDanger()
		if err != nil {
			t.Error(err)
		}

		output := button.Render()

		fmt.Println(Pretty(output))
	})
}

func TestNewCheckboxes(t *testing.T) {
	t.Run("NewCheckboxes", func(t *testing.T) {
		checkboxes := NewCheckboxes("that check check boom!!").AddInitialOption(NewOption("option1", "option1"))
		output := checkboxes.Render()
		fmt.Println(Pretty(output))
	})
}

func TestNewDatePicker(t *testing.T) {
	t.Run("NewDatePicker", func(t *testing.T) {
		datePicker := NewDatePicker("datepicker-action")

		output := datePicker.Render()
		fmt.Println(Pretty(output))

	})
}

func TestDateTimePicker(t *testing.T) {
	t.Run("NewDateTimePicker", func(t *testing.T) {
		dateTimePicker := NewDateTimePicker("datetimepicker-action")
		output := dateTimePicker.Render()
		fmt.Println(Pretty(output))
	})
}

func TestNewImage(t *testing.T) {
	t.Run("NewImage", func(t *testing.T) {
		urlString := "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png"
		urlParsed, err := url.Parse(urlString)
		if err != nil {
			t.Error(err)
		}
		image := NewImage(urlParsed, "Google Logo")
		output := image.Render()
		fmt.Println(output)
	})
}

func TestNewMultiSelectMenuWithConversationsList(t *testing.T) {
	t.Run("NewMultiSelectMenuWithConversationsList", func(t *testing.T) {
		menu := NewMultiSelectMenuWithConversationsList("actionID")
		confirm := NewConfirmationDialog("title", "text", "confirm", "deny")
		menu = menu.AddPlaceholder("placeholder").AddConfirmDialog(confirm)
		fmt.Println(menu.Section().Render())
	})
}

// todo: add massive amount of validation

func TestNewMultiSelectMenuWithExternalDataSource(t *testing.T) {
	t.Run("NewMultiSelectMenuWithExternalDataSource", func(t *testing.T) {
		menu := NewMultiSelectMenuWithExternalDataSource("actionId")

		opt := NewOption("initial", "value-1")
		menu = menu.AddInitialOption(opt).MaxSelectedItems(3).AddPlaceholder("Select items")

		fmt.Println(menu.Section())
	})
}

func TestNewMultiSelectMenuWithPublicChannelsSelect(t *testing.T) {
	t.Run("NewMultiSelectMenuWithPublicChannelsSelect", func(t *testing.T) {
		menu := NewMultiSelectMenuWithPublicChannelsSelect("actionID")
		fmt.Println(menu.Section().Render())
	})
}

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
		menu := NewMultiSelectMenuWithStaticOptions("text1234")

		opt := NewOption("Wait for it", "value-0")
		opt2 := NewOption("Initial", "value-1")
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

func TestNewMultiSelectMenuWithUserList(t *testing.T) {
	menu := NewMultiSelectMenuWithUserList("actionId")

	menu = menu.AddInitialUser("Sarah P").AddInitialUser("user2")

	menu = menu.AddPlaceholder("placeholder")

	fmt.Println(menu.Section())

}

func TestNewNumberInput(t *testing.T) {
	t.Run("NewNumberInput", func(t *testing.T) {
		number := NewNumberInput("actionId").FocusOnLoad().DecimalAllowed().MinValue(1).MaxValue(10).InitialValue(5).Placeholder("placeholder")

		output := number.Input("Label").Render()
		fmt.Println(output)
	})
}

func TestNewOverflowMenu(t *testing.T) {
	menu := NewOverflowMenu("actionId")

	opt1 := NewOption("option-1", "value-1")
	opt2 := NewOption("option-2", "value-a")
	opt3 := NewOption("option-3", "value-x")

	menu = menu.AddOption(opt1).AddOption(opt2).AddOption(opt3)

	output := menu.Render()

	fmt.Println(output)

}

func TestNewPlainTextInput(t *testing.T) {
	t.Run("NewPlainTextInput", func(t *testing.T) {
		pti := NewPlainTextInput("actionId").SetMaxLength(12).SetMinLength(3).FocusOnLoad()
		output := pti.Render()
		fmt.Println(output)
	})
}

func TestNewRadioButton(t *testing.T) {
	t.Run("NewRadioButton", func(t *testing.T) {
		rb := NewRadioButton("actionId").AddOption(NewOption("option-1", "value-1")).AddOption(NewOption("option-2", "value-a")).AddOption(NewOption("option-3", "value-x"))
		output := rb.Render()
		fmt.Println(output)
	})
}

func TestNewSelectMenuWithConversationsList(t *testing.T) {
	t.Run("NewSelectMenuWithConversationsList", func(t *testing.T) {
		menu := NewSelectMenuWithConversationsList("actionID")
		confirm := NewConfirmationDialog("title", "text", "confirm", "deny")
		menu = menu.AddPlaceholder("placeholder").AddConfirmDialog(confirm)
		fmt.Println(menu.Section().Render())
	})
}

// todo: add massive amount of validation

func TestNewSelectMenuWithExternalDataSource(t *testing.T) {
	t.Run("NewSelectMenuWithExternalDataSource", func(t *testing.T) {
		menu := NewSelectMenuWithExternalDataSource("actionId")

		opt := NewOption("initial", "value-1")
		menu = menu.AddInitialOption(opt).AddPlaceholder("placeholder").SetMinQueryLength(6)

		fmt.Println(menu.Section())
	})
}

func TestNewMultiSelectMenuWithPublicChannelsSelectone(t *testing.T) {
	t.Run("NewSelectMenuWithPublicChannelsSelect", func(t *testing.T) {
		menu := NewSelectMenuWithPublicChannelsSelect("actionID").EnableResponseUrlEnabled().FocusOnLoad()
		fmt.Println(menu.Section().Render())
	})
}

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

		opt := NewOption("Wait for it", "value-0")
		opt2 := NewOption("Initial", "value-1")

		optionGroup := NewOptionGroup("Group 1").AddOption(opt).AddOption(opt2)

		menu = menu.AddOptionGroup(optionGroup)

		// convert to section
		section := menu.Section().Render()
		if section == validMenuWithInitialAndPlaceholder {
			t.Errorf("Rendered section does not match expected JSON: %s", section)
		}

		fmt.Println(section)
	})
}

func TestNewSelectMenuWithUserList(t *testing.T) {
	menu := NewSelectMenuWithUserList("actionId")

	menu = menu.SetInitialUser("Sarah P")

	menu = menu.AddPlaceholder("placeholder")

	output := menu.Section().Render()

	fmt.Println(output)

}

func TestNewTimePicker(t *testing.T) {
	t.Run("NewTimePicker", func(t *testing.T) {
		tp := NewTimePicker("actionId")
		output := tp.Render()
		t.Log(output)
	})
}

func TestNewURLInput(t *testing.T) {
	t.Run("NewURLInput", func(t *testing.T) {
		website, err := url.Parse("https://www.google.com")
		if err != nil {
			t.Error(err)
		}

		ui := NewURLInput("actionId").UpdateActionId("New Id").UpdateInitialValue(website)
		output := ui.Render()
		t.Log(output)
	})
}
