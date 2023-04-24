package expressiontree

import (
	"github.com/marcusljx/eevee3/lib/utils"
	"reflect"
)

type Operation[T any] func(T, T) T

func (o Operation[T]) String() string {
	tt := reflect.TypeOf(o)
	return utils.RebuildSignature(tt)
}
