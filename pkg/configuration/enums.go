package configuration

import (
	"github.com/ghodss/yaml"
	. "myproj.com/clmgr-lrm/pkg/common"
	"errors"
)

// Actions

type actionType int16

const (
	at_start     actionType = iota
	at_stop
	at_notify
	at_reload
	at_promote
	at_demote
	at_monitor
	at_methaData
)

func (a *actionType) GetActionType(at string) error {
	switch at {
	case "start":
		*a = at_start
	case "stop":
		*a = at_stop
	case "notify":
		*a = at_notify
	case "reload":
		*a = at_reload
	case "promote":
		*a = at_promote
	case "demote":
		*a = at_demote
	case "monitor":
		*a = at_monitor
	case "methadata":
		*a = at_methaData
	default:
		Logger.Error("Can't parse action type")
		return errors.New("failed to parse action type")
	}

	return nil
}

type onFailAction int16

const (
	of_ignore  onFailAction = iota
	of_block
	of_stop
	of_restart
	of_fence
	of_standby
)

func (a *onFailAction) UnmarshallYAML(b []byte) error {
	param := make(map[string]string)
	err := yaml.Unmarshal(b, &param)
	if err != nil {
		Logger.Error("Can't unmarshal onfail action from yaml, err: %s", err.Error())
		return err
	}

	switch param["on-fail"] {
	case "ignore":
		*a = of_ignore
	case "block":
		*a = of_block
	case "stop":
		*a = of_stop
	case "restart":
		*a = of_restart
	case "fence":
		*a = of_fence
	case "standby":
		*a = of_standby
	default:
		Logger.Error("Can't unmarshal onfail actions from yaml")
		return errors.New("failed to unmarshal onfail action")
	}

	return nil
}

type actionRole int16

const (
	ar_started actionRole = iota
	ar_stopped
	ar_master
	ar_slave
)

func (a *actionRole) UnmarshallYAML(b []byte) error {
	param := make(map[string]string)
	err := yaml.Unmarshal(b, &param)
	if err != nil {
		Logger.Error("Can't unmarshal parameters from yaml, err: %s", err.Error())
		return err
	}

	switch param["role"] {
	case "started":
		*a = ar_started
	case "stopped":
		*a = ar_stopped
	case "master":
		*a = ar_master
	case "slave":
		*a = ar_slave
	default:
		Logger.Error("Can't unmarshal actions from yaml")
		return errors.New("failed to unmarshal action")
	}

	return nil
}

// Parameter
type contentType int16

const (
	ct_int    contentType = iota
	ct_float
	ct_string
	ct_bool
)

func (t *contentType) UnmarshallYAML(b []byte) error {
	param := make(map[string]string)
	err := yaml.Unmarshal(b, &param)
	if err != nil {
		Logger.Error("Can't unmarshal parameters from yaml, err: %s", err.Error())
		return err
	}

	switch param["type"] {
	case "int":
		*t = ct_int
	case "float":
		*t = ct_float
	case "string":
		*t = ct_string
	case "bool":
		*t = ct_bool
	default:
		Logger.Error("Can't unmarshal parameters from yaml")
		return errors.New("failed to unmarshal parameter")
	}

	return nil
}
