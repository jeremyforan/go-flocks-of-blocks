package image

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
