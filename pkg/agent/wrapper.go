package agent

import (
	. "myproj.com/clmgr-lrm/pkg/common"
	"errors"
)

type Wrapper interface {
	Do(Action) (interface{}, error)
}

func (ag *agent) Do(act Action) (interface{}, error) {
	if _, item := GetFromSlice(InterfaceSlice(ag.Config.Actions), act); item != nil {
		res := item.(Action).Operation()
		if IsError(res) {
			return nil, res.(error)
		}
		return res, nil
	}
	return nil, errors.New("the action doesn't provided for this agent")
}