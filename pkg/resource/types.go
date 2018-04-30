package resource

type ResourceType int

const (
	// todo: describe all resource types in comments
	primitive ResourceType = iota
	master
	slave
	clusterwide
	clone
)
