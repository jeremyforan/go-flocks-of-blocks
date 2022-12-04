package dispatchactionconfig

import "go-flocks-of-blocks/common"

type DispatchActionTypes string

const (
	OnEnterPressed     DispatchActionTypes = "on_enter_pressed"
	OnCharacterEntered DispatchActionTypes = "on_character_entered"
)

type DispatchActionConfig struct {
	triggerActionsOn []DispatchActionTypes
}

type abstractionDispatchActionConfig struct {
	TriggerActionsOn []string
}

func NewDispatchActionConfig() DispatchActionConfig {
	return DispatchActionConfig{
		triggerActionsOn: []DispatchActionTypes{},
	}
}

func (d *DispatchActionConfig) setTriggerActionsOn(triggerActionsOn DispatchActionTypes) {
	d.triggerActionsOn = append(d.triggerActionsOn, triggerActionsOn)
}

func (d *DispatchActionConfig) removeTriggerActionsOn() {
	d.triggerActionsOn = []DispatchActionTypes{}
}

// OnEnterPressed chain function to add on_enter_pressed to an existing dispatch action config
func (d DispatchActionConfig) OnEnterPressed() DispatchActionConfig {
	d.setTriggerActionsOn(OnEnterPressed)
	return d
}

// OnCharacterEntered chain function to add on_character_entered to an existing dispatch action config
func (d DispatchActionConfig) OnCharacterEntered() DispatchActionConfig {
	d.setTriggerActionsOn(OnCharacterEntered)
	return d
}

// RemoveTriggerActionsOn remove add trigger actions on from dispatch action config
func (d DispatchActionConfig) RemoveTriggerActions() DispatchActionConfig {
	d.removeTriggerActionsOn()
	return d
}

// Template generates the template for the block
func (d abstractionDispatchActionConfig) Template() string {
	return `{
"trigger_actions_on": [{{range $index, $element := .TriggerActionsOn}}{{if $index}}, {{end}}"{{$element}}"{{end}}]
}`
}

// abstraction
func (d DispatchActionConfig) abstraction() abstractionDispatchActionConfig {
	return abstractionDispatchActionConfig{
		TriggerActionsOn: removeDuplicateStr(d.triggerActionsOn),
	}
}

// Render the block
func (d DispatchActionConfig) Render() string {
	output := common.Render(d.abstraction())
	return common.Pretty(output)
}

func removeDuplicateStr(strSlice []DispatchActionTypes) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		item := string(item)
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

// todo: make the Dispatch Config an interface
