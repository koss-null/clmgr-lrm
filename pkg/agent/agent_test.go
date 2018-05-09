package agent

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestAgent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "All Utils Test")
}

var _ = Describe("ParseConfig", func() {
	var a Agent
	Context("initially", func() {
		It("Creating default agent", func() {
			a = &agent{
				Config: agentConfig{
					Name:    "testResourceAgent1",
					Version: "0.1",
					Longdesc: "this is a long description" +
						"\nof the first test resource agent" +
						"\nI'll add here some more strings" +
						"\nto make it look like all this" +
						"\nhave some science",
					Shortdesc:  "it's a short description",
					Parameters: []Parameter{
						Parameter{
							Name: "active",
							Unique: true,
							Required: false,
							Longdesc: "longdesc of active paramether",
						}
					},
					Actions:    []Action{},
				},
				scriptPath: "/opt/clmgr/agents/test_agent1",
			}
		})
	})
})
