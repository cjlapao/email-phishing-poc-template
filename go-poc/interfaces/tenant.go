package interfaces

type Tenant interface {
	Id() string
	IsValid() string
	CorrelationId() string
}
