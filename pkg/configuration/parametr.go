package configuration

type Parameter struct {
	Name      string      `yaml:"name"`
	Unique    bool        `yaml:"unique"`
	Required  bool        `yaml:"required"`
	ContType  contentType `yaml:"contentType"`
	Longdesc  string      `yaml:"longdesc,omitempty"`
	Shortdesc string      `yaml:"shrotdesc,omitempty"`
}