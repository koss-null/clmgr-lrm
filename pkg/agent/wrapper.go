package agent

import (
	. "myproj.com/clmgr-lrm/pkg/common"
	"errors"
)

type Wrapper interface {
	Do(ActionType) (interface{}, error)
}

func (ag *agent) Do(act ActionType) (interface{}, error) {
	if _, item := GetFromSliceF(ToInterface(ag.Config.Actions), act,
		func(x interface{}) interface{} {
			return x.(Action).Name
		}); item != nil {
		res := item.(Action).Operation()
		if IsError(res) {
			return nil, res.(error)
		}
		return res, nil
	}
	return nil, errors.New("the action doesn't provided for this agent")
}
