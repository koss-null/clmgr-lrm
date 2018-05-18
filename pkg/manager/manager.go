package manager

import (
	. "myproj.com/clmgr-lrm/pkg/pool"
	"context"
	ops "myproj.com/clmgr-lrm/protobuf/compiled/protobuf/operations"
	"io"
	"myproj.com/clmgr-lrm/pkg/agent"
	"strings"
	"github.com/google/logger"
)

type (
	manager struct {
		AgentPool
	}

	Manager interface {
	}
)

func (m *manager) Run() {

}

func serialize(tp ops.Operation2Perform_OperationType) agent.ActionType {
	pbType := strings.ToLower(tp.String())
	switch pbType {
	case "metadata":
		return agent.ActionType("meta-data")
	default:
		return agent.ActionType(pbType)
	}
}

func (m *manager) Perform(ctx context.Context, stream ops.AgentOperationPerformer_PerformServer) (error) {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		ot := serialize(in.OpType)
		res, errCh := m.AgentPool.Do(in.Uuid, ot)
		for {
			or := new(ops.OperationResult)
			select {
			case err := <-errCh:
				logger.Errorf("Error during %s, on agent %s, err: %s", string(ot), in.Uuid, err.Error())
				or.Uuid = in.Uuid
				or.Error = err.Error()
				or.OpState = ops.OperationResult_OP_FAILED
				stream.Send(or)
				break
			case resp := <-res:
				or.Uuid = in.Uuid
				or.Metha = []byte(resp.(string))
				or.OpState = ops.OperationResult_OP_OK
				stream.Send(or)
				break
			default:
				or.Uuid = in.Uuid
				or.OpState = ops.OperationResult_OP_IN_PROGRESS
				stream.Send(or)
			}
		}
	}
}
