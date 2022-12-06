package flocksofblocks

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/block/divider"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition"
	"github.com/jeremyforan/go-flocks-of-blocks/element"
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

		option3 := composition.NewOption("Option 3", "value3")
		option3 = option3.SetDescription(composition.NewPlainText("Option 3 description"))

		oGroup := composition.NewOptionGroup("Select an option").AddOption(composition.NewOption("Option 1", "value1")).AddOption(composition.NewOption("Option 2", "value2"))
		fmt.Println(oGroup.Render())
		fmt.Println(common.Pretty(oGroup.Render()))

		button := element.NewButton("Click This", "button1").MakeStyleDanger().AddUrl(url)

		action := block.NewAction("Block").AddElement(button)

		div := block.NewDividerBlock(divider.BlockId("divider1"))

		msg = msg.AddBlock(action).AddBlock(div)

		output := msg.Render()
		fmt.Println("Message output: \n\n", output)

		slackUrl := msg.GenerateKitBuilderUrl()

		fmt.Println("Message url: \n\n", slackUrl)

	})
}
