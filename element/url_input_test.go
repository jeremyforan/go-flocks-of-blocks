package element

import (
	"net/url"
	"testing"
)

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
