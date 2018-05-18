package pool

import (
	. "myproj.com/clmgr-lrm/pkg/agent"
	"github.com/google/uuid"
	"github.com/google/logger"
	"errors"
)

type (
	agentPool struct {
		pool map[uuid.UUID]Wrapper
	}

	AgentPool interface {
		Do(id string, act ActionType) (chan interface{}, chan error)
		Add(Agent) (uuid.UUID, error)
		Remove(id uuid.UUID)
	}
)

func NewPool() AgentPool {
	return &agentPool{make(map[uuid.UUID]Wrapper)}
}

func (p *agentPool) Do(id string, act ActionType) (chan interface{}, chan error) {
	res, errCh := make(chan interface{}), make(chan error)

	go func() {
		uuid, err := uuid.Parse(id)
		if err != nil {
			logger.Error("Can't parse agent uuid")
			errCh <- err
			return
		}
		ag, ok := p.pool[uuid]
		if !ok {
			logger.Error("Can't find uuid in agentPool")
			errCh <- errors.New("no such uuid")
			return
		}

		r, e := ag.Do(act)
		if e != nil {
			errCh <- e
		}
		if res != nil {
			res <- r
		}
		return
	}()
	return res, errCh
}


func (p *agentPool) Add(a Agent) (uuid.UUID, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		logger.Errorf("Can't generate UUID for agent, err: %s", err.Error())
		return [16]byte{}, err
	}
	p.pool[id] = a
	return id, nil
}

func (p *agentPool) Remove(id uuid.UUID) {
	delete(p.pool, id)
}