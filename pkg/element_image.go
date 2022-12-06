package flocksofblocks

import (
	"net/url"
)

type ImageElement struct {
	slackType ElementType
	imageUrl  *url.URL
	altText   string
}

func NewImageElement(imageUrl *url.URL, altText string) ImageElement {
	return ImageElement{
		slackType: ImageElementType,
		imageUrl:  imageUrl,
		altText:   altText,
	}
}

// SetImageUrl sets the image url for the image.
func (i *ImageElement) setImageUrl(imageUrl *url.URL) {
	i.imageUrl = imageUrl
}

// SetAltText sets the alt text for the image.
func (i *ImageElement) setAltText(altText string) {
	i.altText = altText
}

// abstractImage struct
type abstractImage struct {
	Type     string
	ImageUrl string
	AltText  string
}

// abstraction method
func (i *ImageElement) abstraction() abstractImage {
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
func (i ImageElement) Render() string {
	return Render(i.abstraction())
}
