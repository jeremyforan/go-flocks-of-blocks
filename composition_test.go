package flocksofblocks

import (
	"fmt"
	"testing"
)

func TestNewConfirmationDialog(t *testing.T) {
	t.Run("NewConfirmationDialog", func(t *testing.T) {
		confirm := NewConfirmationDialog("title", "text", "confirm", "deny")
		if confirm.Render() == "" {
			t.Error("failed to render")
		}

		confirm = confirm.SetStyle(StylePrimary)
		if confirm.Render() == "" {
			t.Error("failed to render")
		}
		fmt.Println(confirm.Render())
	})
}

func TestNewDispatchActionConfig(t *testing.T) {
	t.Run("NewDispatchActionConfig", func(t *testing.T) {
		dispatchActionConfig := NewDispatchActionConfig().OnEnterPressed().OnCharacterEntered()

		output := dispatchActionConfig.Render()
		fmt.Println(output)
	})
}

const (
	outputValidFilter = `"filter": {
	"include": ["im","mpim"],
	"exclude_external_shared_channels": true
}`
)

func TestNewFilter(t *testing.T) {
	t.Run("should return a new filter", func(t *testing.T) {
		filter := NewFilter()
		if filter.Render() != "" {
			t.Errorf("Expected empty string, got %s", filter.Render())
		}
		filter = filter.IncludeIM().IncludeMPIM().ExcludeExternalSharedChannels()
		if filter.Render() != outputValidFilter {
			t.Errorf("Expected %s, got %s", outputValidFilter, filter.Render())
		}
		filter = filter.ExcludeBotUsers().IncludeIM().IncludeMPIM().ExcludeExternalSharedChannels()
		output := filter.Render()
		fmt.Println(output)
	})
}
