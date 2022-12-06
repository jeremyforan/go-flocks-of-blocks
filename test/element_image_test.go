package element

import (
	"fmt"
	"net/url"
	"testing"
)

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
