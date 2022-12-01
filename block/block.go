package block

type BlockType string

const (
	Actions BlockType = "actions"
	Context BlockType = "context"
	Divider BlockType = "divider"
	File    BlockType = "file"
	Header  BlockType = "header"
	Image   BlockType = "image"
	Input   BlockType = "input"
	Section BlockType = "section"
	Video   BlockType = "video"
)

// stringer for BlockType
func (b BlockType) String() string {
	return string(b)
}
