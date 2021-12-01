package entities

import "github.com/google/uuid"

var _tenant *Tenant

type Tenant struct{}

func (t Tenant) Id() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return uuid.String()
}
