package flocksofblocks

type Foo struct {
	blockId
	actionId
}

// NewFoo
func NewFoo() Foo {
	return Foo{}
}

func (f Foo) SetBlockId(blockId string) Foo {
	f.blockId.SetValue(blockId)
	return f
}

func (f Foo) SetActionId(ActionId string) Foo {
	f.actionId.SetValue(ActionId)
	return f
}

func (f Foo) Template() string {
	return `{
		{{range $index, $field := .FieldsNames}}{{if $index}},{{end}}{{ .$field}}{{end}}
	}`
}

func (f Foo) String() string {
	return Render(f)
}
