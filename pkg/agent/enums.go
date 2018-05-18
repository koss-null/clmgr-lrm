package agent

// Actions

type ActionType string

const (
	at_start    ActionType = "start"
	at_stop                = "stop"
	at_notify              = "notify"
	at_reload              = "reload"
	at_promote             = "promote"
	at_demote              = "demote"
	at_monitor             = "monitor"
	at_metaData            = "meta-data"
)

type onFailAction string

const (
	of_ignore  onFailAction = "ignore"
	of_block                = "block"
	of_stop                 = "stop"
	of_restart              = "restart"
	of_fence                = "fence"
	of_standby              = "standby"
)

type actionRole string

const (
	ar_none    actionRole = "none"
	ar_started            = "started"
	ar_stopped            = "stopped"
	ar_master             = "master"
	ar_slave              = "slave"
)

// Parameter
type contentType string

const (
	ct_int    contentType = "int"
	ct_float              = "float"
	ct_string             = "string"
	ct_bool               = "bool"
)

type ContentType struct {
	Ct  contentType `yaml:"type"`
	Def interface{} `yaml:"default,omitempty"`
}

func (ct *ContentType) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// creating anonimus struct to avoid recursion
	var a struct {
		Ct  contentType `yaml:"type"`
		Def interface{} `yaml:"default,omitempty"`
	}
	if err := unmarshal(&a); err != nil {
		return err
	}
	ct.Ct = a.Ct
	if a.Def == nil {
		ct.Def = nil
		return nil
	}
	switch a.Ct {
	case ct_int:
		ct.Def = a.Def.(int)
	case ct_float:
		ct.Def = a.Def.(float32)
	case ct_string:
		ct.Def = a.Def.(string)
	case ct_bool:
		ct.Def = a.Def.(bool)
	}
	return nil
}
