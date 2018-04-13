package resource

type (
	// todo: think about ResourceMetha realisation structs
	// todo: probably add some additional funcs into interface
	resourceMetha struct {
		data []byte
	}

	ResourceMetha interface {
		Set(interface{})
		Get() interface{}
	}
)

// Accepts only []byte
func (rm *resourceMetha) Set(data interface{}) {
	rm.data = data.([]byte)
}

// Returns []byte convertable interface
func (rm *resourceMetha) Get() interface{} {
	return rm.data
}