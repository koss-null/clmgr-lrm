package configuration

type (
	agent_config struct {
		name       string
		version    string
		longdesc   string
		shortdesc  string
		parameters []Parameter
		actions    []Action
	}

	agent struct {
	}

	Agent interface {
	}
)
