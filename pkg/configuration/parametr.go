package configuration

type Parameter struct {
	name        string
	unique      bool
	required    bool
	contentType contentType
	longdesc    string
	shortdesc   string
}
