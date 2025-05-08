package domain

import (
	"time"

	"github.com/google/uuid"
)

type Function struct {
	ID      uuid.UUID
	Name    string
	Author  string
	Code    []byte
	Version int
	Created time.Time
}

func NewFunction(name string, author string, code []byte) Function {
	return Function{
		ID:      uuid.New(),
		Name:    name,
		Author:  author,
		Code:    code,
		Version: 0,
		Created: time.Now(),
	}
}
