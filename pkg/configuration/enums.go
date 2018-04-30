package configuration

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

type onFailAction int16

const (
	of_ignore  onFailAction = iota
	of_block
	of_stop
	of_restart
	of_fence
	of_standby
)

type actionRole int16

const (
	ar_started actionRole = iota
	ar_stopped
	ar_master
	ar_slave
)

// Parameter
type contentType int16

const (
	ct_int    contentType = iota
	ct_float
	ct_string
	ct_bool
)