package flocksofblocks

import (
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"net/url"
)

type Image struct {
	slackType ElementType
	imageUrl  *url.URL
	altText   string
}

func NewImage(imageUrl *url.URL, altText string) Image {
	return Image{
		slackType: ImageElement,
		imageUrl:  imageUrl,
		altText:   altText,
	}
}

// SetImageUrl sets the image url for the image.
func (i *Image) setImageUrl(imageUrl *url.URL) {
	i.imageUrl = imageUrl
}

// SetAltText sets the alt text for the image.
func (i *Image) setAltText(altText string) {
	i.altText = altText
}

// abstractImage struct
type abstractImage struct {
	Type     string
	ImageUrl string
	AltText  string
}

// abstraction method
func (i *Image) abstraction() abstractImage {
	return abstractImage{
		Type:     i.slackType.String(),
		ImageUrl: i.imageUrl.String(),
		AltText:  i.altText,
	}
}

// Template method
func (i abstractImage) Template() string {
	return `{
		"type": "{{.Type}}",
		"image_url": "{{.ImageUrl}}",
		"alt_text": "{{.AltText}}"
	}`
}

// Render method
func (i Image) Render() string {
	return common.Render(i.abstraction())
}
