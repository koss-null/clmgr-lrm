package agent

import (
	"strings"
	"io/ioutil"

	. "myproj.com/clmgr-lrm/config"

	"github.com/google/logger"
	"gopkg.in/yaml.v2"
)

const configFormat = "yaml"

type (
	agentConfig struct {
		Name       string      `yaml:"name"`
		Version    string      `yaml:"version,omitempty"`
		Longdesc   string      `yaml:"longdesc,omitempty"`
		Shortdesc  string      `yaml:"shortdesc,omitempty"`
		Parameters []Parameter `yaml:"parameters,omitempty"`
		Actions    []Action    `yaml:"actions,omitempty"`
	}

	agent struct {
		Config     agentConfig `yaml:",inline"`
		scriptPath string
	}

	Agent interface {
		ParseConfig() error
		getConfig() agentConfig

		Name() string
		Version() string
		LongDesc() string
		ShortDesc() string

		Start() error
		Stop() error
		Monitor() interface{}
		Notify() error
		Reload() error
		Promote() error
		Demote() error
		MethaData() interface{}

		Do(Action) (interface{}, error)
	}
)

func defaultConfig() agentConfig {
	return agentConfig{Version: "none"}
}

func defaultAgent(path string) Agent {
	return &agent{
		scriptPath: path,
		Config:     defaultConfig(),
	}
}

/*
	CreateAgent() takes the Name of agent, which is expected to
	be on default clmgr agent folder, parses it's Config,
	the Config Name should be tha same as the Name of a script
	with *.yaml extension at the end
 */
func Create(agentName string) (Agent, error) {
	agentPath := strings.Join([]string{Config.AgentPath, agentName}, "/")
	ag := defaultAgent(agentPath)

	if err := ag.ParseConfig(); err != nil {
		logger.Errorf("Failed to parse Config for agent %s, err %s", agentName, err.Error())
		return nil, err
	}

	return ag, nil
}

func (ag *agent) ParseConfig() error {
	configPath := strings.Join([]string{ag.scriptPath, configFormat}, ".")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Errorf("Can't read Config for agent %s, err: %s", ag.scriptPath, err.Error())
		return err
	}

	if err := yaml.Unmarshal(data, ag); err != nil {
		logger.Errorf("Can't unmarshall agent Config %s, err: %s", ag.scriptPath, err.Error())
		return err
	}

	// accotiating actions with current agent
	for i := range ag.Config.Actions {
		ag.Config.Actions[i].personalizeAction(ag)
	}

	return nil
}

func (ag *agent) getConfig() agentConfig {
	return ag.Config
}

func (ag *agent) Name() string {
	return ag.Config.Name
}

func (ag *agent) Version() string {
	return ag.Config.Version
}

func (ag *agent) LongDesc() string {
	return ag.Config.Longdesc
}

func (ag *agent) ShortDesc() string {
	return ag.Config.Shortdesc
}
