package flocksofblocks

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/block/action"
	"github.com/jeremyforan/go-flocks-of-blocks/block/divider"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/option"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/optiongroup"
	"github.com/jeremyforan/go-flocks-of-blocks/element/button"
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

		option3 := option.NewOption("Option 3", "value3")
		option3 = option3.SetDescription(compositiontext.NewPlainText("Option 3 description"))

		oGroup := optiongroup.NewOptionGroup("Select an option").AddOption(option.NewOption("Option 1", "value1")).AddOption(option.NewOption("Option 2", "value2"))
		fmt.Println(oGroup.Render())
		fmt.Println(common.Pretty(oGroup.Render()))

		button := button.NewButton("Click This", "button1").MakeStyleDanger().AddUrl(url)

		action := action.NewAction("Block").AddElement(button)

		div := divider.NewDividerBlock(divider.BlockId("divider1"))

		msg = msg.AddBlock(action).AddBlock(div)

		output := msg.Render()
		fmt.Println("Message output: \n\n", output)
	})
}
