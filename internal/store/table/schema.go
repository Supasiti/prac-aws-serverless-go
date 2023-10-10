package table

import "errors"

const (
	User Schema = "USER"
)

var (
	ErrInvalidSchema = errors.New("invalid schema")
)

type Schema string

func (t Schema) String() string {
	return string(t)
}

func FromString(s string) (Schema, error) {
	switch s {
	case User.String():
		return User, nil
	}
	return "", ErrInvalidSchema
}
