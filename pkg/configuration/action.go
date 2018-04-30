package configuration

import "time"

/*
	Action struct describes operations on resource that
	can be called by the lrm Agent
	name - any action supported by the agent
	timeout - how many wait before declaring the operation failed
	onFail - what to do if the action fails
	interval - how frequently to perform the operation
	enabled - ignore this operation definition if false
	recordPending - if true there comes a special record to check if the operation in progress
	role - specifies the role of resource on which the action will be performed
 */
type Action struct {
	name          actionType
	timeout       time.Duration
	onFail        onFailAction
	interval      time.Duration
	enabled       bool
	recordPending bool
	role          actionRole
}
