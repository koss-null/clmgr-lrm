package agent

import (
	"time"
)

/*
	Action struct describes operations on resource that
	can be called by the lrm Agent
	Name - any action supported by the agent
	timeout - how many wait before declaring the operation failed
	onFail - what to do if the action fails
	interval - how frequently to perform the operation
	enabled - ignore this operation definition if false
	recordPending - if true there comes a special record to check if the operation in progress
	role - specifies the role of resource on which the action will be performed
 */
type Action struct {
	Name          actionType    `yaml:"type"`
	Timeout       time.Duration `yaml:"timeout,omitempty"`
	OnFail        onFailAction  `yaml:"on-fail,omitempty"`
	Interval      time.Duration `yaml:"interval,omitempty"`
	Enabled       bool          `yaml:"enabled,omitempty"`
	RecordPending bool          `yaml:"record-pending,omitempty"`
	Role          actionRole    `yaml:"role,omitempty"`
	Operation     func() (interface{})
}

/*
	Performs unmarshalling Action with setting default values
 */
func (act *Action) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// creating anonimus struct to avoid recursion
	var a struct {
		Name          actionType    `yaml:"type"`
		Timeout       time.Duration `yaml:"timeout,omitempty"`
		OnFail        onFailAction  `yaml:"on-fail,omitempty"`
		Interval      time.Duration `yaml:"interval,omitempty"`
		Enabled       bool          `yaml:"enabled,omitempty"`
		RecordPending bool          `yaml:"record-pending,omitempty"`
		Role          actionRole    `yaml:"role,omitempty"`
	}
	// setting default
	a.Name = "defaultName"
	a.Timeout = 20 * time.Second
	a.OnFail = of_ignore
	a.Interval = 20 * time.Second
	a.Enabled = true
	a.RecordPending = false
	a.Role = ar_none
	if err := unmarshal(&a); err != nil {
		return err
	}
	// copying into aim struct fields
	*act = Action{a.Name, a.Timeout, a.OnFail, a.Interval,
		a.Enabled, a.RecordPending, a.Role, nil}
	return nil
}

func wrapError(f func() error) func() interface{} {
	return func() interface{} {
		return f()
	}
}

func (act *Action) personalizeAction(ag *agent) {
	for i := range ag.Config.Actions {
		switch ag.Config.Actions[i].Name {
		case at_start:
			ag.Config.Actions[i].Operation = wrapError(ag.Start)
		case at_stop:
			ag.Config.Actions[i].Operation = wrapError(ag.Stop)
		case at_monitor:
			ag.Config.Actions[i].Operation = ag.Monitor
		case at_notify:
			ag.Config.Actions[i].Operation = wrapError(ag.Notify)
		case at_reload:
			ag.Config.Actions[i].Operation = wrapError(ag.Reload)
		case at_promote:
			ag.Config.Actions[i].Operation = wrapError(ag.Promote)
		case at_demote:
			ag.Config.Actions[i].Operation = wrapError(ag.Demote)
		case at_metaData:
			ag.Config.Actions[i].Operation = ag.MetaData
		}
	}
}
