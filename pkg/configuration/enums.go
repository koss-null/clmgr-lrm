package configuration

// Actions

type actionType string

const (
	at_start     actionType = "start"
	at_stop                 = "stop"
	at_notify               = "notify"
	at_reload               = "reload"
	at_promote              = "promote"
	at_demote               = "demote"
	at_monitor              = "monitor"
	at_methaData            = "metha-data"
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
	ar_started actionRole = "started"
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