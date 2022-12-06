package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"net/url"
	"testing"
)

func TestButton(t *testing.T) {
	t.Run("Create Button Element", func(t *testing.T) {

		u, err := url.Parse("http://bing.com/search?q=dotnet")
		if err != nil {
			t.Error(err)
		}

		button := NewButton("Click This", "button1")
		button.AddUrl(u).MakeStyleDanger()
		if err != nil {
			t.Error(err)
		}

		output := button.Render()

		fmt.Println(common.Pretty(output))
	})
}
