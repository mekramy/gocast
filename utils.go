// From html/template/content.go
// Copyright 2011 The Go Authors. All rights reserved.
package gocast

import (
	"fmt"
	"reflect"
)

// valueOf returns the value that the input interface{} points to,
// traversing through any number of pointer indirections. If the input
// is not a pointer or is nil, it returns the input itself. This function
// is useful for obtaining the underlying value of a pointer, regardless
// of how many levels of pointers there are.
//
// Parameters:
// - value: The input interface{} which may be a pointer or a value.
//
// Returns:
//   - The underlying value that the input points to, or the input itself
//     if it is not a pointer or is nil.
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
//
// Type Parameters:
// - T: The type that is being checked for implementation of the interface I.
// - I: The interface that is being checked against the type T.
//
// Returns:
// - bool: True if T implements I, otherwise false.
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
//
// Type Parameters:
// - T: The type whose name is to be obtained.
//
// Returns:
// - string: The name of the type T.
func typeName[T any]() string {
	var o T
	return reflect.TypeOf(o).String()
}
