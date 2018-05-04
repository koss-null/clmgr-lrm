package configuration

import (
	"strings"
	. "myproj.com/clmgr-lrm/config"

	"github.com/go-yaml/yaml"
	"github.com/google/logger"
	"io/ioutil"
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
	}
)

/*
	CreateAgent() takes the Name of agent, which is expected to
	be on default clmgr agent folder, parses it's Config,
	the Config Name should be tha same as the Name of a script
	with *.yaml extension at the end
 */
func CreateAgent(agentName string) (Agent, error) {
	agentPath := strings.Join([]string{Config.AgentPath, agentName}, "/")
	ag := agent{scriptPath: agentPath} // todo add here method to set default values

	if err := ag.ParseConfig(); err != nil {
		logger.Errorf("Failed to parse Config for agent %s, err %s", agentName, err.Error())
		return nil, err
	}

	return &ag, nil
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

	return nil
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

func (ag *agent) Start() error {
	logger.Infof("Resource %s Start op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Stop() error {
	logger.Infof("Resource %s Stop op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Monitor() interface{} {
	logger.Infof("Resource %s Monitor op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Notify() error {
	logger.Infof("Resource %s Notify op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Reload() error {
	logger.Infof("Resource %s Reload op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Promote() error {
	logger.Infof("Resource %s Promote op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Demote() error {
	logger.Infof("Resource %s Demote op is stubbed", ag.Name())
	return nil
}

func (ag *agent) MethaData() interface{} {
	logger.Infof("Resource %s MethaData op is stubbed", ag.Name())
	return nil
}
