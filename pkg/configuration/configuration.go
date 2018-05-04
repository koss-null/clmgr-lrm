package configuration

import (
	"strings"
	"io/ioutil"

	"github.com/ghodss/yaml"

	. "myproj.com/clmgr-lrm/config"
	"github.com/google/logger"
)

const configFormat = "yaml"

type (
	agentConfig struct {
		name       string      `yaml:"name"`
		version    string      `yaml:"vertsion,omitempty"`
		longdesc   string      `yaml:"longdesc,omitempty"`
		shortdesc  string      `yaml:"shortdesc,omitempty"`
		parameters []Parameter `yaml:"parameters,omitempty"`
		actions    []Action    `yaml:"actions,omitempty"`
	}

	agent struct {
		config     agentConfig
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
	CreateAgent() takes the name of agent, which is expected to
	be on default clmgr agent folder, parses it's config,
	the config name should be tha same as the name of a script
	with *.yaml extension at the end
 */
func CreateAgent(agentName string) (Agent, error) {
	agentPath := strings.Join([]string{Config.AgentPath, agentName}, "/")
	ag := agent{scriptPath: agentPath}

	if err := ag.ParseConfig(); err != nil {
		logger.Error("Failed to parse config for agent %s, err %s", agentName, err.Error())
		return nil, err
	}

	return &ag, nil
}

func (ag *agent) ParseConfig() error {
	configPath := strings.Join([]string{ag.scriptPath, configFormat}, ".")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		logger.Error("Can't read config (%s) for agent %s, err: %s", configPath, ag.scriptPath, err.Error())
		return err
	}

	if err := yaml.Unmarshal(data, &ag); err != nil {
		logger.Error("Can't unmarshall agent config %s, err: %s", configPath, err.Error())
		return err
	}
	return nil
}

func (ag *agent) Name() string {
	return ag.config.name
}

func (ag *agent) Version() string {
	return ag.config.version
}

func (ag *agent) LongDesc() string {
	return ag.config.longdesc
}

func (ag *agent) ShortDesc() string {
	return ag.config.shortdesc
}

func (ag *agent) Start() error {
	logger.Info("Resource %s Start op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Stop() error {
	logger.Info("Resource %s Stop op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Monitor() interface{} {
	logger.Info("Resource %s Monitor op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Notify() error {
	logger.Info("Resource %s Notify op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Reload() error {
	logger.Info("Resource %s Reload op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Promote() error {
	logger.Info("Resource %s Promote op is stubbed", ag.Name())
	return nil
}

func (ag *agent) Demote() error {
	logger.Info("Resource %s Demote op is stubbed", ag.Name())
	return nil
}

func (ag *agent) MethaData() interface{} {
	logger.Info("Resource %s MethaData op is stubbed", ag.Name())
	return nil
}
