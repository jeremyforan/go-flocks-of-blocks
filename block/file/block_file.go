package file

import (
	"github.com/jeremyforan/go-flocks-of-blocks/block"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
)

type File struct {
	slackType  block.BlockType
	externalId string
	source     string
	blockId    string

	optional fileOptions
}

type fileOptions struct {
	BlockId bool
}

// NewFile creates a new file.
func NewFile(externalId string, source string) File {
	return File{
		slackType:  block.File,
		externalId: externalId,
		source:     source,
		optional: fileOptions{
			BlockId: false,
		},
	}
}

// SetBlockId sets the block id for the block.
func (f *File) setBlockId(blockId string) {
	f.blockId = blockId
	f.optional.BlockId = true
}

func (f *File) removeBlockId() {
	f.blockId = ""
	f.optional.BlockId = false
}

// AddBlockId chain function to add block id to an existing file
func (f File) AddBlockId(blockId string) File {
	f.setBlockId(blockId)
	return f
}

// RemoveBlockId remove add block id from file
func (f File) RemoveBlockId() File {
	f.removeBlockId()
	return f
}

// BlockRender is a dummy function to satisfy the Block interface.
func (f File) BlockRender() {}

// fileAbstraction is a helper struct to generate the abstraction for the file.
type fileAbstraction struct {
	Type       string
	ExternalId string
	Source     string
	BlockId    string
	Optional   fileOptions
}

// abstraction is a helper function to generate the abstraction for the file.
func (f File) abstraction() fileAbstraction {
	return fileAbstraction{
		Type:       string(f.slackType),
		ExternalId: f.externalId,
		Source:     f.source,
		BlockId:    f.blockId,

		Optional: f.optional,
	}
}

// template is a helper function to generate the template for the file.
func (f fileAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	"external_id": "{{.ExternalId}}",
	"source": "{{.Source}}"{{if .Optional.BlockId}},
	"block_id": "{{.BlockId}}"{{end}}
}`
}

// render is a helper function to generate the json for the file.
func (f File) Render() string {
	return common.Render(f.abstraction())
}
