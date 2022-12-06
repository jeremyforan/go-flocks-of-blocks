package block

import (
	"github.com/jeremyforan/go-flocks-of-blocks"
	"github.com/jeremyforan/go-flocks-of-blocks/common"
)

type Action struct {
	slackType BlockType
	elements  []flocksofblocks.Element
	blockId   string

	optionals actionOptions
}

func (a Action) BlockRender() {}

func NewAction(blockId string) Action {
	return Action{
		slackType: ActionsBlock,
		elements:  []flocksofblocks.Element{},
		blockId:   blockId,
		optionals: actionOptions{
			blockId: false,
		},
	}
}

// SetBlockId sets the block id for the block.
func (a *Action) setBlockId(blockId string) {
	a.blockId = blockId
	a.optionals.blockId = true
}

func (a *Action) removeBlockId() {
	a.blockId = ""
	a.optionals.blockId = false
}

func (a *Action) addElement(element flocksofblocks.Element) {
	a.elements = append(a.elements, element)
}

type actionOptions struct {
	blockId bool
}

type actionAbstraction struct {
	Type     string
	Elements []flocksofblocks.Element
	BlockId  string

	Optional actionOptions
}

// AddBlockId chain function to add block id to an existing action block
func (a Action) AddBlockId(blockId string) Action {
	a.setBlockId(blockId)
	return a
}

// RemoveBlockId remove add block id from action block
func (a Action) RemoveBlockId() Action {
	a.removeBlockId()
	return a
}

// AddElement Add element to existing action block.
func (a Action) AddElement(element flocksofblocks.Element) Action {
	a.addElement(element)
	return a
}

// generate abstraction from action
func (a Action) abstraction() actionAbstraction {
	return actionAbstraction{
		Type:     string(a.slackType),
		Elements: a.elements,
		BlockId:  a.blockId,

		Optional: a.optionals,
	}
}

func (a actionAbstraction) Template() string {
	return `{
	"type": "{{.Type}}",
	
	"elements": [{{range $index, $element := .Elements}}{{if $index}},{{end}}{{$element.Render}}{{end}}]

{{if .BlockId}},
	"block_id": "{{.BlockId}}"
{{end}}
}`
}

// Render the block
func (a Action) Render() string {
	output := common.Render(a.abstraction())
	return common.Pretty(output)
}

type ActionType interface {
	Action()
}
