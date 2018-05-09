package agent

type Parameter struct {
	Name      string      `yaml:"name"`
	Unique    bool        `yaml:"unique"`
	Required  bool        `yaml:"required"`
	ContType  ContentType `yaml:"content-type"`
	Longdesc  string      `yaml:"longdesc,omitempty"`
	Shortdesc string      `yaml:"shortdesc,omitempty"`
}