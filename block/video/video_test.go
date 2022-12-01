package video

import (
	"fmt"
	"go-flocks-of-blocks/common"
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
		fmt.Println(common.Pretty(output))
	})
}
