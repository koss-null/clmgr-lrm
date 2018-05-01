package configuration

type (
	agentConfig struct {
		name       string
		version    string
		longdesc   string
		shortdesc  string
		parameters []Parameter
		actions    []Action
	}

	agent struct {
		agentConfig
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


func CreateAgent(agentName string) Agent {

}