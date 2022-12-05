package image

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
	"github.com/jeremyforan/go-flocks-of-blocks/composition/compositiontext"
	"net/url"
)

type Image struct {
	slackType block.BlockType
	imageUrl  *url.URL
	altText   string

	title   compositiontext.CompositionText
	blockId string

	optional imageOptions
}

type imageOptions struct {
	Title   bool
	BlockId bool
}

// NewImage creates a new image.
func NewImage(imageUrl *url.URL, altText string) Image {
	return Image{
		slackType: block.Image,
		imageUrl:  imageUrl,
		altText:   altText,
		optional: imageOptions{
			Title:   false,
			BlockId: false,
		},
	}
}

// SetTitle sets the title for the image.
func (i *Image) setTitle(title string) {
	i.title = compositiontext.NewPlainText(title)
	i.optional.Title = true
}

func (i *Image) removeTitle() {
	i.optional.Title = false
}

// AddTitle chain function to add title to an existing image
func (i Image) AddTitle(title string) Image {
	i.setTitle(title)
	return i
}

// RemoveTitle remove add title from image
func (i Image) RemoveTitle() Image {
	i.removeTitle()
	return i
}

// SetBlockId sets the block id for the block.
func (i *Image) setBlockId(blockId string) {
	i.blockId = blockId
	i.optional.BlockId = true
}

func (i *Image) removeBlockId() {
	i.optional.BlockId = false
}

// AddBlockId chain function to add block id to an existing image
func (i Image) AddBlockId(blockId string) Image {
	i.setBlockId(blockId)
	return i
}

// RemoveBlockId remove add block id from image
func (i Image) RemoveBlockId() Image {
	i.removeBlockId()
	return i
}

// abstraction is a helper struct to generate the abstraction for the image.
type imageAbstraction struct {
	Type     string
	ImageUrl string
	AltText  string
	Title    compositiontext.CompositionText
	BlockId  string
	Optional imageOptions
}

// generate abstraction from image
func (i Image) abstraction() imageAbstraction {
	return imageAbstraction{
		Type:     i.slackType.String(),
		ImageUrl: i.imageUrl.String(),
		AltText:  i.altText,
		Title:    i.title,
		BlockId:  i.blockId,
		Optional: i.optional,
	}
}

// Template returns the template for the image.
func (i imageAbstraction) Template() string {
	return `{
		"type": "{{.Type}}",
		"image_url": "{{.ImageUrl}}",
		"alt_text": "{{.AltText}}"{{if .Optional.Title}},
		"title": {{.Title.Render}}{{end}}{{if .Optional.BlockId}},
		"block_id": "{{.BlockId}}"{{end}}
	}`
}

// Render renders the image to a string.
func (i Image) Render() string {
	return common.Render(i.abstraction())
}
