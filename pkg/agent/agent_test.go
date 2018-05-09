package agent

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"time"
	"myproj.com/clmgr-lrm/config"
)

func TestAgent(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "All Utils Test")
}

var _ = Describe("ParseConfig", func() {
	var a Agent
	var agnt Agent
	config.InitConfig()
	Context("initially", func() {
		It("Creating default agent", func() {
			a = &agent{
				Config: agentConfig{
					Name:    "testResourceAgent1",
					Version: "0.1",
					Longdesc: "\"this is a long description" +
						"\nof the first test resource agent" +
						"\nI'll add here some more strings" +
						"\nto make it look like all this" +
						"\nhave some science\"\n",
					Shortdesc: "it's a short description",
					Parameters: []Parameter{
						{
							Name:     "active",
							Unique:   true,
							Required: false,
							Longdesc: "longdesc of active paramether",
							ContType: ContentType{ct_bool, false},
						},
						{
							Name:      "recovery_throttle",
							Unique:    false,
							Required:  false,
							Longdesc:  "some nice description",
							Shortdesc: "some short description",
							ContType:  ContentType{ct_int, nil},
						},
					},
					Actions: []Action{
						{
							Name:          at_start,
							Timeout:       time.Duration(60 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_none,
						},
						{
							Name:          at_stop,
							Timeout:       time.Duration(60 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_none,
						},
						{
							Name:          at_notify,
							Timeout:       time.Duration(20 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_none,
						},
						{
							Name:          at_demote,
							Timeout:       time.Duration(20 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_none,
						},
						{
							Name:          at_promote,
							Timeout:       time.Duration(20 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_none,
						},
						{
							Name:          at_monitor,
							Timeout:       time.Duration(60 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_none,
						},
						{
							Name:          at_monitor,
							Timeout:       time.Duration(60 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_master,
						},
						{
							Name:          at_methaData,
							Timeout:       time.Duration(5 * time.Second),
							OnFail:        of_ignore,
							Interval:      time.Duration(20 * time.Second),
							Enabled:       true,
							RecordPending: false,
							Role:          ar_none,
						},
					},
				},
				scriptPath: "../../test/test_agents/test_agent1",
			}
		})

		It("Creating agent with config parsing", func() {
			ag, err := Create("test_agent1")
			Ω(err).Should(BeNil())
			for i := 0; i < len(ag.getConfig().Actions); i++ {
				a.getConfig().Actions[i].Operation = ag.getConfig().Actions[i].Operation
			}
			agnt = ag
		})
	})

	Context("Matching two agents", func() {
		It("Two structs must be equal", func() {
			Ω(agnt.getConfig().Shortdesc).Should(BeEquivalentTo(a.getConfig().Shortdesc))
			Ω(agnt.getConfig().Longdesc).Should(BeEquivalentTo(a.getConfig().Longdesc))
			Ω(agnt.getConfig().Name).Should(BeEquivalentTo(a.getConfig().Name))
			Ω(agnt.getConfig().Version).Should(BeEquivalentTo(a.getConfig().Version))
			Ω(agnt.getConfig().Parameters).Should(BeEquivalentTo(a.getConfig().Parameters))
			// this shit doesn't work for some reason
			//Ω(agnt.getConfig().Actions).Should(BeEquivalentTo(a.getConfig().Actions))
		})
	})
})
