package json

import "encoding/json"

var Null = NewValue[any](nil)

type Value[T any] struct {
	v T
}

func NewValue[T any](v T) Value[T] {
	return Value[T]{v: v}
}

func (o Value[T]) MarshalJSON() ([]byte, error) {
	data, err := json.Marshal(o.v)
	if err != nil {
		return nil, err
	}

	return data, nil
}
