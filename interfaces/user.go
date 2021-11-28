package interfaces

type User interface {
	Id() string
	Name() string
	IsValid() string
	IsAuthenticated() bool
	Issuer() string
}
