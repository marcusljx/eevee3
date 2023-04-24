package utils

import (
	"fmt"
	"reflect"
	"strings"
)

func RebuildSignature(tt reflect.Type) string {
	if tt.Kind() != reflect.Func {
		panic(fmt.Sprintf("expected a type with Kind()==reflect.Func, got %s instead", tt.String()))
	}
	var b strings.Builder
	b.WriteString("func(")
	for i := 0; i < tt.NumIn(); i++ {
		b.WriteString(tt.In(i).String())
		if i != tt.NumIn()-1 {
			b.WriteString(", ")
		}
	}
	b.WriteString(") ")

	if tt.NumOut() == 1 {
		b.WriteString(tt.Out(0).String())
		return b.String()
	}

	b.WriteRune('(')
	for i := 0; i < tt.NumOut(); i++ {
		b.WriteString(tt.Out(i).String())
		if i != tt.NumOut()-1 {
			b.WriteString(", ")
		}
	}
	b.WriteRune(')')

	return b.String()
}
