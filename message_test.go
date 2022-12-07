package flocksofblocks

import (
	"fmt"
	"net/url"
	"testing"
)

func TestMessage(t *testing.T) {
	t.Run("valid message", func(t *testing.T) {
		msg := NewMessage()

		url, err := url.Parse("http://google.ca")
		if err != nil {
			t.Error(err)
		}

		option3 := NewOption("Option 3", "value3")
		option3 = option3.SetDescription(NewPlainText("Option 3 description"))

		oGroup := NewOptionGroup("Select an option").AddOption(NewOption("Option 1", "value1")).AddOption(NewOption("Option 2", "value2"))
		fmt.Println(oGroup.Render())
		fmt.Println(Pretty(oGroup.Render()))

		button := NewButton("Click This", "button1").MakeStyleDanger().AddUrl(url)

		action := go_flocks_of_blocks.NewAction("Block").AddElement(button)

		div := go_flocks_of_blocks.NewDividerBlock()

		msg = msg.AddBlock(action).AddBlock(div)

		output := msg.Render()
		fmt.Println("Message output: \n\n", output)

		slackUrl := msg.GenerateKitBuilderUrl()

		fmt.Println("Message url: \n\n", slackUrl)

	})
}
