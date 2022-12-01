package video

import (
	"go-flocks-of-blocks/block"
	"go-flocks-of-blocks/common"
	"go-flocks-of-blocks/composition/compositiontext"
	"net/url"
)

type Video struct {
	slackType    block.BlockType
	title        compositiontext.CompositionText
	thumbnailUrl *url.URL
	videoUrl     *url.URL
	altText      string

	// optionals
	authorName   string
	providerName string

	description compositiontext.CompositionText

	providerIconUrl *url.URL
	titleUrl        *url.URL
	blockId         string

	// optional
	optional optionalVideo
}

type optionalVideo struct {
	AuthorName      bool
	ProviderName    bool
	Description     bool
	ProviderIconUrl bool
	TitleUrl        bool
	BlockId         bool
}

func NewVideo(title string, thumbnailUrl *url.URL, videoUrl *url.URL, altText string) Video {
	return Video{
		slackType: block.Video,
		title:     compositiontext.NewPlainText(title),

		thumbnailUrl: thumbnailUrl,
		videoUrl:     videoUrl,
		altText:      altText,

		optional: optionalVideo{
			AuthorName:      false,
			ProviderName:    false,
			Description:     false,
			ProviderIconUrl: false,
			TitleUrl:        false,
			BlockId:         false,
		},
	}
}

// setAuthorName
func (v *Video) setAuthorName(authorName string) {
	v.authorName = authorName
	v.optional.AuthorName = true
}

// RemoveAuthorName
func (v *Video) removeAuthorName() {
	v.optional.AuthorName = false
}

// setProviderName
func (v *Video) setProviderName(providerName string) {
	v.providerName = providerName
	v.optional.ProviderName = true
}

// removeProviderName
func (v *Video) removeProviderName() {
	v.optional.ProviderName = false
}

// setDescription
func (v *Video) setDescription(description string) {
	v.description = compositiontext.NewPlainText(description).EnableEmoji()
	v.optional.Description = true
}

// removeDescription
func (v *Video) removeDescription() {
	v.optional.Description = false
}

// setProviderIconUrl
func (v *Video) setProviderIconUrl(providerIconUrl *url.URL) {
	v.providerIconUrl = providerIconUrl
	v.optional.ProviderIconUrl = true
}

// removeProviderIconUrl
func (v *Video) removeProviderIconUrl() {
	v.optional.ProviderIconUrl = false
}

// setTitleUrl
func (v *Video) setTitleUrl(titleUrl *url.URL) {
	v.titleUrl = titleUrl
	v.optional.TitleUrl = true
}

// removeTitleUrl
func (v *Video) removeTitleUrl() {
	v.optional.TitleUrl = false
}

// setBlockId
func (v *Video) setBlockId(blockId string) {
	v.blockId = blockId
	v.optional.BlockId = true
}

// removeBlockId
func (v *Video) removeBlockId() {
	v.optional.BlockId = false
}

// abstraction structure
type abstractionVideo struct {
	Type         string
	Title        compositiontext.CompositionText
	ThumbnailUrl string
	VideoUrl     string
	AltText      string

	AuthorName      string
	ProviderName    string
	Description     compositiontext.CompositionText
	ProviderIconUrl string
	TitleUrl        string
	BlockId         string

	Optional optionalVideo
}

// abstract
func (v Video) abstraction() abstractionVideo {
	providerIconUrl := ""
	if v.optional.ProviderIconUrl {
		providerIconUrl = v.providerIconUrl.String()
	}

	titleUrl := ""
	if v.optional.TitleUrl {
		titleUrl = v.titleUrl.String()
	}

	return abstractionVideo{
		Type:         string(v.slackType),
		Title:        v.title,
		ThumbnailUrl: v.thumbnailUrl.String(),
		VideoUrl:     v.videoUrl.String(),
		AltText:      v.altText,
		AuthorName:   v.authorName,
		ProviderName: v.providerName,
		Description:  v.description,

		ProviderIconUrl: providerIconUrl,
		TitleUrl:        titleUrl,
		BlockId:         v.blockId,
		Optional:        v.optional,
	}
}

// Render
func (v Video) Render() string {
	return common.Render(v.abstraction())
}

// Template
func (v abstractionVideo) Template() string {
	return `{
		"type": "{{.Type}}",
		"title": {{.Title.Render}},
		"thumbnail_url": "{{.ThumbnailUrl}}",
		"video_url": "{{.VideoUrl}}",
		"alt_text": "{{.AltText}}"{{if .Optional.AuthorName}},
		"author_name": "{{.AuthorName}}"{{end}}{{if .Optional.ProviderName}},
		"provider_name": "{{.ProviderName}}"{{end}}{{if .Optional.Description}},
		"description": {{.Description.Render}}{{end}}{{if .Optional.ProviderIconUrl}},
		"provider_icon_url": "{{.ProviderIconUrl}}"{{end}}{{if .Optional.TitleUrl}},
		"title_url": "{{.TitleUrl}}"{{end}}{{if .Optional.BlockId}},
		"block_id": "{{.BlockId}}"{{end}}
	}`
}
