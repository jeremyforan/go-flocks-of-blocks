package flocksofblocks

func (b Button) Actions() Actions {
	return NewAction().AddElement(b)
}

type Actions struct {
	slackType BlockType
	elements  []Element

	blockId

	optionals actionOptions
}

type actionOptions struct {
	blockId bool
}

func NewAction() Actions {
	return Actions{
		slackType: ActionsBlock,
		elements:  []Element{},
		optionals: actionOptions{
			blockId: false,
		},
	}
}

func (a *Actions) addElement(element Element) {
	a.elements = append(a.elements, element)
}

type actionAbstraction struct {
	Type     string
	Elements []Element
	BlockId  string

	Optionals actionOptions
}

// AddBlockId chain function to add block id to an existing action block
func (a Actions) AddBlockId(blockId string) Actions {
	a.blockId.SetValue(blockId)
	return a
}

// RemoveBlockId remove add block id from action block
func (a Actions) RemoveBlockId() Actions {
	a.blockId.UnsetValue()
	return a
}

// AddElement Add element to existing action block.
func (a Actions) AddElement(element Element) Actions {
	a.addElement(element)
	return a
}

// generate abstraction from action
func (a Actions) abstraction() actionAbstraction {
	return actionAbstraction{
		Type:     string(a.slackType),
		Elements: a.elements,
		BlockId:  a.blockId.String(),

		Optionals: a.optionals,
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
func (a Actions) Render() string {
	raw := Render(a.abstraction())
	return Pretty(raw)
}

type ActionType interface {
	Action()
}
