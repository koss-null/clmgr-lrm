package resource

type (
	resource struct {
		Name string
		ResourceType

		start func() error
		stop func() error
		restart func() error
		monitor func() (<-chan []byte, <-chan error)
		metadata func([]byte) error
	}

	Resource interface {
		start() error
		stop() error
		restart() error
		monitor() (<-chan interface{}, <-chan error)
		metadata(ResourceMetha) error
	}
)


