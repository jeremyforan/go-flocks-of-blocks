package flocksofblocks

import "testing"

func TestNewActionBlock(t *testing.T) {
	t.Run("NewActionBlock", func(t *testing.T) {
		action := NewAction("Block")

		output := action.Render()
		t.Log(output)
	})
}
package flocksofblocks

import (
	"fmt"
	flocksofblocks "github.com/jeremyforan/go-flocks-of-blocks"
	"testing"
)

func TestFileRender(t *testing.T) {
	t.Run("valid file", func(t *testing.T) {
		file := flocksofblocks.NewFile("externalId", "source")
		file = file.AddBlockId("file1")
		output := file.Render()
		fmt.Println("File output: \n\n", output)
	})

}
package block

import (
	"fmt"
	"testing"
)

func TestHeader(t *testing.T) {
	t.Run("valid header", func(t *testing.T) {
		header := NewHeader("header text")
		output := header.Render()
		fmt.Println("Header output: \n\n", output)
	})
}
package block

import (
	"fmt"
	url2 "net/url"
	"testing"
)

func TestImage(t *testing.T) {
	t.Run("valid image", func(t *testing.T) {

		url, err := url2.Parse("https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
		if err != nil {
			t.Error(err)
		}

		img := NewImage(url, "Google Logo")
		output := img.Render()
		fmt.Println("Image output: \n\n", output)
	})

}
package flocksofblocks

import "testing"

func TestNewInputTest(t *testing.T) {
	t.Run("NewInputTest", func(t *testing.T) {
		input := NewInputTest("Block")

		output := input.Render()
		t.Log(output)
	}
}
package block

import (
	"fmt"
	"github.com/jeremyforan/go-flocks-of-blocks"
	"net/url"
	"testing"
)

func TestVideo(t *testing.T) {
	t.Run("NewVideo", func(t *testing.T) {
		thumbNailUrl, err := url.Parse("https://i.ytimg.com/vi/RRxQQxiM7AA/hqdefault.jpg")
		if err != nil {
			t.Error(err)
		}

		videoUrl, err := url.Parse("https://www.youtube.com/embed/RRxQQxiM7AA?feature=oembed&autoplay=1")
		if err != nil {
			t.Error(err)
		}

		titleUrl, err := url.Parse("https://www.youtube.com/watch?v=RRxQQxiM7AA")
		if err != nil {
			t.Error(err)
		}

		providerIconUrl, err := url.Parse("https://www.example.com/provider_icon.jpg")
		if err != nil {
			t.Error(err)
		}
		video := NewVideo("title", thumbNailUrl, videoUrl, "How to use Slack?")
		video = video.AddTitleUrl(titleUrl).AddProviderName("YouTube").AddAuthorName("Arcado Buendia").AddProviderIconUrl(providerIconUrl)
		video = video.AddDescription("Slack is a new way to communicate with your team. It's faster, better organized and more secure than email.")

		output := video.Render()
		fmt.Println(flocksofblocks.Pretty(output))
	})
}
