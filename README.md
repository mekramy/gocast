# GoCast Package

Package `gocast` provides utilities for type casting and conversion in Go. It defines an interface `Caster` with methods for converting values to various Go types, including primary types like `bool`, `int`, `uint`, `float`, and `string`, as well as slices of these types. The package also includes functions for casting values to these types and handling errors related to type conversion.

## Installation

To install the `gocast` package, use the following command:

```sh
go get github.com/mekramy/gocast
```

Then, import the package in your Go code:

```go
import (
    "github.com/mekramy/gocast"
)
```

## Functions

The package provides functions for casting values to various types and handling errors related to type conversion.

### ToBool

`func ToBool(value interface{}) (bool, error)`

Casts an interface to a `bool` type.

### ToSigned

`func ToSigned[T int | int8 | int16 | int32 | int64](value interface{}) (T, error)`

Casts an interface to a signed integer type.

### ToUnsigned

`func ToUnsigned[T uint | uint8 | uint16 | uint32 | uint64](value interface{}) (T, error)`

Casts an interface to an unsigned integer type.

### ToFloat

`func ToFloat[T float32 | float64](value interface{}) (T, error)`

Casts an interface to a float type.

### ToString

`func ToString(value interface{}) (string, error)`

Casts an interface to a `string` type.

### ToSlice

`func ToSlice(value interface{}) ([]interface{}, error)`

Casts an interface to a `[]interface{}` type.

### ToBoolSlice

`func ToBoolSlice(i interface{}) ([]bool, error)`

Casts an interface to a `[]bool` type.

### ToSignedSlice

`func ToSignedSlice[T int | int8 | int16 | int32 | int64](i interface{}) ([]T, error)`

Casts an interface to a signed integer slice type.

### ToUnsignedSlice

`func ToUnsignedSlice[T uint | uint8 | uint16 | uint32 | uint64](i interface{}) ([]T, error)`

Casts an interface to an unsigned integer slice type.

### ToFloatSlice

`func ToFloatSlice[T float32 | float64](i interface{}) ([]T, error)`

Casts an interface to a float slice type.

### ToStringSlice

`func ToStringSlice(i interface{}) ([]string, error)`

Casts an interface to a `[]string` type.

### Functions Usage

```go
package main

import (
    "fmt"
    "github.com/mekramy/gocast"
)

func main() {
    // Example of ToBool
    boolValue, err := gocast.ToBool("true")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Bool value:", boolValue) // output: true
    }

    // Example of ToSigned
    intValue, err := gocast.ToSigned[int]("-123.32")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Int value:", intValue) // output: 123
    }

    // Example of ToString
    stringValue, err := gocast.ToString(123)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("String value:", stringValue) // output: "123"
    }
}
```

## Caster Interface

The `Caster` interface provides methods for type casting and conversion. It includes methods for checking if a value is nil, retrieving the value as an interface, and converting the value to primary Go types such as `bool`, `int`, `uint`, `float`, and `string`. Each type conversion method has a corresponding safe method that returns a fallback value in case of an error, and methods for converting to slices of each type.

### Methods

- `IsNil() bool`: Checks if the value is nil.
- `Interface() any`: Returns the value as an `interface{}`.
- `Unmarshal(out any) error`: Unmarshals the value using a JSON decoder.
- `Bool() (bool, error)`: Converts the value to a `bool`.
- `BoolSafe(fallback bool) bool`: Converts the value to a `bool`, returning a fallback value in case of an error.
- `Int() (int, error)`: Converts the value to an `int`.
- `IntSafe(fallback int) int`: Converts the value to an `int`, returning a fallback value in case of an error.
- `Int8() (int8, error)`: Converts the value to an `int8`.
- `Int8Safe(fallback int8) int8`: Converts the value to an `int8`, returning a fallback value in case of an error.
- `Int16() (int16, error)`: Converts the value to an `int16`.
- `Int16Safe(fallback int16) int16`: Converts the value to an `int16`, returning a fallback value in case of an error.
- `Int32() (int32, error)`: Converts the value to an `int32`.
- `Int32Safe(fallback int32) int32`: Converts the value to an `int32`, returning a fallback value in case of an error.
- `Int64() (int64, error)`: Converts the value to an `int64`.
- `Int64Safe(fallback int64) int64`: Converts the value to an `int64`, returning a fallback value in case of an error.
- `Uint() (uint, error)`: Converts the value to a `uint`.
- `UintSafe(fallback uint) uint`: Converts the value to a `uint`, returning a fallback value in case of an error.
- `Uint8() (uint8, error)`: Converts the value to a `uint8`.
- `Uint8Safe(fallback uint8) uint8`: Converts the value to a `uint8`, returning a fallback value in case of an error.
- `Uint16() (uint16, error)`: Converts the value to a `uint16`.
- `Uint16Safe(fallback uint16) uint16`: Converts the value to a `uint16`, returning a fallback value in case of an error.
- `Uint32() (uint32, error)`: Converts the value to a `uint32`.
- `Uint32Safe(fallback uint32) uint32`: Converts the value to a `uint32`, returning a fallback value in case of an error.
- `Uint64() (uint64, error)`: Converts the value to a `uint64`.
- `Uint64Safe(fallback uint64) uint64`: Converts the value to a `uint64`, returning a fallback value in case of an error.
- `Float32() (float32, error)`: Converts the value to a `float32`.
- `Float32Safe(fallback float32) float32`: Converts the value to a `float32`, returning a fallback value in case of an error.
- `Float64() (float64, error)`: Converts the value to a `float64`.
- `Float64Safe(fallback float64) float64`: Converts the value to a `float64`, returning a fallback value in case of an error.
- `String() (string, error)`: Converts the value to a `string`.
- `StringSafe(fallback string) string`: Converts the value to a `string`, returning a fallback value in case of an error.
- `Slice() ([]any, error)`: Returns the value as a slice of `interface{}`.
- `SliceSafe(fallback []any) []any`: Returns the value as a slice of `interface{}`, returning a fallback value in case of an error.
- `BoolSlice() ([]bool, error)`: Converts the value to a slice of `bool`.
- `BoolSliceSafe(fallback []bool) []bool`: Converts the value to a slice of `bool`, returning a fallback value in case of an error.
- `IntSlice() ([]int, error)`: Converts the value to a slice of `int`.
- `IntSliceSafe(fallback []int) []int`: Converts the value to a slice of `int`, returning a fallback value in case of an error.
- `Int8Slice() ([]int8, error)`: Converts the value to a slice of `int8`.
- `Int8SliceSafe(fallback []int8) []int8`: Converts the value to a slice of `int8`, returning a fallback value in case of an error.
- `Int16Slice() ([]int16, error)`: Converts the value to a slice of `int16`.
- `Int16SliceSafe(fallback []int16) []int16`: Converts the value to a slice of `int16`, returning a fallback value in case of an error.
- `Int32Slice() ([]int32, error)`: Converts the value to a slice of `int32`.
- `Int32SliceSafe(fallback []int32) []int32`: Converts the value to a slice of `int32`, returning a fallback value in case of an error.
- `Int64Slice() ([]int64, error)`: Converts the value to a slice of `int64`.
- `Int64SliceSafe(fallback []int64) []int64`: Converts the value to a slice of `int64`, returning a fallback value in case of an error.
- `UintSlice() ([]uint, error)`: Converts the value to a slice of `uint`.
- `UintSliceSafe(fallback []uint) []uint`: Converts the value to a slice of `uint`, returning a fallback value in case of an error.
- `Uint8Slice() ([]uint8, error)`: Converts the value to a slice of `uint8`.
- `Uint8SliceSafe(fallback []uint8) []uint8`: Converts the value to a slice of `uint8`, returning a fallback value in case of an error.
- `Uint16Slice() ([]uint16, error)`: Converts the value to a slice of `uint16`.
- `Uint16SliceSafe(fallback []uint16) []uint16`: Converts the value to a slice of `uint16`, returning a fallback value in case of an error.
- `Uint32Slice() ([]uint32, error)`: Converts the value to a slice of `uint32`.
- `Uint32SliceSafe(fallback []uint32) []uint32`: Converts the value to a slice of `uint32`, returning a fallback value in case of an error.
- `Uint64Slice() ([]uint64, error)`: Converts the value to a slice of `uint64`.
- `Uint64SliceSafe(fallback []uint64) []uint64`: Converts the value to a slice of `uint64`, returning a fallback value in case of an error.
- `Float32Slice() ([]float32, error)`: Converts the value to a slice of `float32`.
- `Float32SliceSafe(fallback []float32) []float32`: Converts the value to a slice of `float32`, returning a fallback value in case of an error.
- `Float64Slice() ([]float64, error)`: Converts the value to a slice of `float64`.
- `Float64SliceSafe(fallback []float64) []float64`: Converts the value to a slice of `float64`, returning a fallback value in case of an error.
- `StringSlice() ([]string, error)`: Converts the value to a slice of `string`.
- `StringSliceSafe(fallback []string) []string`: Converts the value to a slice of `string`, returning a fallback value in case of an error.

### Caster Usage

```go
package main

import (
    "fmt"
    "github.com/mekramy/gocast"
)

// Parse int
var value interface{} = "asdfasdf"
caster := gocast.New(value)
intValue, err := caster.Int() // output: error
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Int value:", intValue)
}

// Parse bool
var value interface{} = "1"
caster := gocast.New(value)
boolValue, err := caster.Bool()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Bool value:", boolValue) // output: true
}

// Parse string
var value interface{} = 123.392
caster := gocast.New(value)
stringValue, err := caster.String()
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("String value:", stringValue) // output: "123.392"
}
```

## Errors

The package defines error handling functions for type conversion errors.

### IsNilError

`func IsNilError(err error) bool`

Checks if the provided error is a nil error. Returns true if the error is not nil and is a nil error.

### IsCastError

`func IsCastError(err error) bool`

Checks if the provided error is a casting error. Returns true if the error is not nil and is a casting error.

### Providers

The library also defines several provider interfaces for different types and slices. These interfaces provide methods to converting custom types like standard go `Stringer` interface.

#### Slice Providers

- `SliceProvider`: Provides a method to return a slice of any type.
- `SliceErrorProvider`: Provides a method to return a slice of any type with an error.

#### Bool Providers

- `BoolProvider`: Provides a method to return a boolean value.
- `BoolErrorProvider`: Provides a method to return a boolean value with an error.
- `BoolSliceProvider`: Provides a method to return a slice of boolean values.
- `BoolSliceErrorProvider`: Provides a method to return a slice of boolean values with an error.

#### Int Providers

- `IntProvider`: Provides a method to return an integer value.
- `IntErrorProvider`: Provides a method to return an integer value with an error.
- `IntSliceProvider`: Provides a method to return a slice of integer values.
- `IntSliceErrorProvider`: Provides a method to return a slice of integer values with an error.

#### Int8 Providers

- `Int8Provider`: Provides a method to return an int8 value.
- `Int8ErrorProvider`: Provides a method to return an int8 value with an error.
- `Int8SliceProvider`: Provides a method to return a slice of int8 values.
- `Int8SliceErrorProvider`: Provides a method to return a slice of int8 values with an error.

#### Int16 Providers

- `Int16Provider`: Provides a method to return an int16 value.
- `Int16ErrorProvider`: Provides a method to return an int16 value with an error.
- `Int16SliceProvider`: Provides a method to return a slice of int16 values.
- `Int16SliceErrorProvider`: Provides a method to return a slice of int16 values with an error.

#### Int32 Providers

- `Int32Provider`: Provides a method to return an int32 value.
- `Int32ErrorProvider`: Provides a method to return an int32 value with an error.
- `Int32SliceProvider`: Provides a method to return a slice of int32 values.
- `Int32SliceErrorProvider`: Provides a method to return a slice of int32 values with an error.

#### Int64 Providers

- `Int64Provider`: Provides a method to return an int64 value.
- `Int64ErrorProvider`: Provides a method to return an int64 value with an error.
- `Int64SliceProvider`: Provides a method to return a slice of int64 values.
- `Int64SliceErrorProvider`: Provides a method to return a slice of int64 values with an error.

#### Uint Providers

- `UintProvider`: Provides a method to return a uint value.
- `UintErrorProvider`: Provides a method to return a uint value with an error.
- `UintSliceProvider`: Provides a method to return a slice of uint values.
- `UintSliceErrorProvider`: Provides a method to return a slice of uint values with an error.

#### Uint8 Providers

- `Uint8Provider`: Provides a method to return a uint8 value.
- `Uint8ErrorProvider`: Provides a method to return a uint8 value with an error.
- `Uint8SliceProvider`: Provides a method to return a slice of uint8 values.
- `Uint8SliceErrorProvider`: Provides a method to return a slice of uint8 values with an error.

#### Uint16 Providers

- `Uint16Provider`: Provides a method to return a uint16 value.
- `Uint16ErrorProvider`: Provides a method to return a uint16 value with an error.
- `Uint16SliceProvider`: Provides a method to return a slice of uint16 values.
- `Uint16SliceErrorProvider`: Provides a method to return a slice of uint16 values with an error.

#### Uint32 Providers

- `Uint32Provider`: Provides a method to return a uint32 value.
- `Uint32ErrorProvider`: Provides a method to return a uint32 value with an error.
- `Uint32SliceProvider`: Provides a method to return a slice of uint32 values.
- `Uint32SliceErrorProvider`: Provides a method to return a slice of uint32 values with an error.

#### Uint64 Providers

- `Uint64Provider`: Provides a method to return a uint64 value.
- `Uint64ErrorProvider`: Provides a method to return a uint64 value with an error.
- `Uint64SliceProvider`: Provides a method to return a slice of uint64 values.
- `Uint64SliceErrorProvider`: Provides a method to return a slice of uint64 values with an error.

#### Float32 Providers

- `Float32Provider`: Provides a method to return a float32 value.
- `Float32ErrorProvider`: Provides a method to return a float32 value with an error.
- `Float32SliceProvider`: Provides a method to return a slice of float32 values.
- `Float32SliceErrorProvider`: Provides a method to return a slice of float32 values with an error.

#### Float64 Providers

- `Float64Provider`: Provides a method to return a float64 value.
- `Float64ErrorProvider`: Provides a method to return a float64 value with an error.
- `Float64SliceProvider`: Provides a method to return a slice of float64 values.
- `Float64SliceErrorProvider`: Provides a method to return a slice of float64 values with an error.

#### String Providers

- `StringProvider`: Provides a method to return a string value.
- `StringErrorProvider`: Provides a method to return a string value with an error.
- `StringSliceProvider`: Provides a method to return a slice of string values.
- `StringSliceErrorProvider`: Provides a method to return a slice of string values with an error.
