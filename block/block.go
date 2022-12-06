package block

type BlockType string

const (
	ActionsBlock BlockType = "actions"
	ContextBlock BlockType = "context"
	DividerBlock BlockType = "divider"
	FileBlock    BlockType = "file"
	HeaderBlock  BlockType = "header"
	ImageBlock   BlockType = "image"
	InputBlock   BlockType = "input"
	SectionBlock BlockType = "section"
	VideoBlock   BlockType = "video"
)

// stringer for BlockType
func (b BlockType) String() string {
	return string(b)
}

type Block interface {
	BlockRender()
}
