package element

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
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

		fmt.Println(flocksofblocks.Pretty(output))
	})
}
