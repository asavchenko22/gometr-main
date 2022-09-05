package gometr

type Checker struct {
	items []Checkable
}

type Measurable interface {
	GetMetrics() string
}

type Checkable interface {
	Ping() error
	GetID() string
	Health() bool
	Add(check Checkable)
}
