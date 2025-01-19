// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
package gocast

import (
	"fmt"
	"math"
	"reflect"
)

// valueOf returns the value that the input interface{} points to,
// traversing through any number of pointer indirections. If the input
// is not a pointer or is nil, it returns the input itself. This function
// is useful for obtaining the underlying value of a pointer, regardless
// of how many levels of pointers there are.
func valueOf(value any) any {
	// nil check
	if value == nil {
		return nil
	} else if typ := reflect.TypeOf(value); typ.Kind() == reflect.Ptr &&
		reflect.ValueOf(value).IsNil() {
		return nil
	}

	// parse value
	if typ := reflect.TypeOf(value); typ.Kind() != reflect.Ptr {
		return value
	}

	// parse pointer
	val := reflect.ValueOf(value)
	for val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	return val.Interface()
}

// stringOf takes an interface{} as input and returns an interface{}.
// It traverses through pointer indirections to find a value that implements either
// the fmt.Stringer or error interface. If the input is nil, it returns nil.
// If the input or any of its dereferenced values implement fmt.Stringer or error,
// it returns that value. Otherwise, it returns the final dereferenced value.
func stringOf(value any) any {
	// nil check
	if value == nil {
		return nil
	} else if typ := reflect.TypeOf(value); typ.Kind() == reflect.Ptr &&
		reflect.ValueOf(value).IsNil() {
		return nil
	}

	// parse
	errorType := reflect.TypeOf((*error)(nil)).Elem()
	stringerType := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	val := reflect.ValueOf(value)
	for !val.Type().Implements(stringerType) && !val.Type().Implements(errorType) && val.Kind() == reflect.Ptr && !val.IsNil() {
		val = val.Elem()
	}
	return val.Interface()
}

// isImplementsOf checks if a type T implements an interface I.
//
// This function is a generic utility that determines whether a type T
// implements an interface I at compile time. It can be useful for
// compile-time assertions and type checking.
func isImplementsOf[T any, I any]() bool {
	var o T
	_, ok := interface{}(o).(I)
	return ok
}

// typeName returns the name of the type T as a string.
//
// This function uses reflection to obtain the name of the type T. It is
// useful for debugging, logging, or any situation where you need to
// programmatically obtain the name of a type.
func typeName[T any]() string {
	var o T
	return reflect.TypeOf(o).String()
}

// intInRange checks if a given int64 value falls within the range of a specified integer type T.
//
// This function is useful for determining whether a value can be safely
// converted to a different integer type without overflow or underflow.
func intInRange[T int | int8 | int16 | int32 | int64](value int64) bool {
	var sample T
	switch any(sample).(type) {
	case int:
		return value >= math.MinInt && value <= math.MaxInt
	case int8:
		return value >= math.MinInt8 && value <= math.MaxInt8
	case int16:
		return value >= math.MinInt16 && value <= math.MaxInt16
	case int32:
		return value >= math.MinInt32 && value <= math.MaxInt32
	case int64:
		return value >= math.MinInt64 && value <= math.MaxInt64
	default:
		return false
	}
}

// uintInRange checks if a given uint64 value falls within the range of a specified unsigned integer type T.
//
// This function is useful for determining whether a value can be safely
// converted to a different unsigned integer type without overflow.
func uintInRange[T uint | uint8 | uint16 | uint32 | uint64](original int64, value uint64) bool {
	if original > 0 {
		var sample T
		switch any(sample).(type) {
		case uint:
			return value <= math.MaxUint
		case uint8:
			return value <= math.MaxUint8
		case uint16:
			return value <= math.MaxUint16
		case uint32:
			return value <= math.MaxUint32
		case uint64:
			return value <= math.MaxUint64
		}
	}
	return false
}

// floatInRange checks if a given float64 value falls within the range of a specified floating-point type T.
//
// This function is useful for determining whether a value can be safely
// converted to a different floating-point type without overflow.
func floatInRange[T float32 | float64](v float64) bool {
	var sample T
	switch any(sample).(type) {
	case float32:
		return v >= -math.MaxFloat32 && v <= math.MaxFloat32
	case float64:
		return v >= -math.MaxFloat64 && v <= math.MaxFloat64
	default:
		return false
	}
}
