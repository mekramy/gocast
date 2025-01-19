package gocast

// SliceProvider is an interface that provides a method to return a slice of any type.
type SliceProvider interface {
	Slice() []any
}

// SliceErrorProvider is an interface that provides a method to return a slice of any type with an error.
type SliceErrorProvider interface {
	Slice() ([]any, error)
}

// BoolProvider is an interface that provides a method to return a boolean value.
type BoolProvider interface {
	Bool() bool
}

// BoolErrorProvider is an interface that provides a method to return a boolean value with an error.
type BoolErrorProvider interface {
	Bool() (bool, error)
}

// BoolSliceProvider is an interface that provides a method to return a slice of boolean values.
type BoolSliceProvider interface {
	BoolSlice() []bool
}

// BoolSliceErrorProvider is an interface that provides a method to return a slice of boolean values with an error.
type BoolSliceErrorProvider interface {
	BoolSlice() ([]bool, error)
}

// IntProvider is an interface that provides a method to return an integer value.
type IntProvider interface {
	Int() int
}

// IntErrorProvider is an interface that provides a method to return an integer value with an error.
type IntErrorProvider interface {
	Int() (int, error)
}

// IntSliceProvider is an interface that provides a method to return a slice of integer values.
type IntSliceProvider interface {
	IntSlice() []int
}

// IntSliceErrorProvider is an interface that provides a method to return a slice of integer values with an error.
type IntSliceErrorProvider interface {
	IntSlice() ([]int, error)
}

// Int8Provider is an interface that provides a method to return an int8 value.
type Int8Provider interface {
	Int8() int8
}

// Int8ErrorProvider is an interface that provides a method to return an int8 value with an error.
type Int8ErrorProvider interface {
	Int8() (int8, error)
}

// Int8SliceProvider is an interface that provides a method to return a slice of int8 values.
type Int8SliceProvider interface {
	Int8Slice() []int8
}

// Int8SliceErrorProvider is an interface that provides a method to return a slice of int8 values with an error.
type Int8SliceErrorProvider interface {
	Int8Slice() ([]int8, error)
}

// Int16Provider is an interface that provides a method to return an int16 value.
type Int16Provider interface {
	Int16() int16
}

// Int16ErrorProvider is an interface that provides a method to return an int16 value with an error.
type Int16ErrorProvider interface {
	Int16() (int16, error)
}

// Int16SliceProvider is an interface that provides a method to return a slice of int16 values.
type Int16SliceProvider interface {
	Int16Slice() []int16
}

// Int16SliceErrorProvider is an interface that provides a method to return a slice of int16 values with an error.
type Int16SliceErrorProvider interface {
	Int16Slice() ([]int16, error)
}

// Int32Provider is an interface that provides a method to return an int32 value.
type Int32Provider interface {
	Int32() int32
}

// Int32ErrorProvider is an interface that provides a method to return an int32 value with an error.
type Int32ErrorProvider interface {
	Int32() (int32, error)
}

// Int32SliceProvider is an interface that provides a method to return a slice of int32 values.
type Int32SliceProvider interface {
	Int32Slice() []int32
}

// Int32SliceErrorProvider is an interface that provides a method to return a slice of int32 values with an error.
type Int32SliceErrorProvider interface {
	Int32Slice() ([]int32, error)
}

// Int64Provider is an interface that provides a method to return an int64 value.
type Int64Provider interface {
	Int64() int64
}

// Int64ErrorProvider is an interface that provides a method to return an int64 value with an error.
type Int64ErrorProvider interface {
	Int64() (int64, error)
}

// Int64SliceProvider is an interface that provides a method to return a slice of int64 values.
type Int64SliceProvider interface {
	Int64Slice() []int64
}

// Int64SliceErrorProvider is an interface that provides a method to return a slice of int64 values with an error.
type Int64SliceErrorProvider interface {
	Int64Slice() ([]int64, error)
}

// UintProvider is an interface that provides a method to return a uint value.
type UintProvider interface {
	Uint() uint
}

// UintErrorProvider is an interface that provides a method to return a uint value with an error.
type UintErrorProvider interface {
	Uint() (uint, error)
}

// UintSliceProvider is an interface that provides a method to return a slice of uint values.
type UintSliceProvider interface {
	UintSlice() []uint
}

// UintSliceErrorProvider is an interface that provides a method to return a slice of uint values with an error.
type UintSliceErrorProvider interface {
	UintSlice() ([]uint, error)
}

// Uint8Provider is an interface that provides a method to return a uint8 value.
type Uint8Provider interface {
	Uint8() uint8
}

// Uint8ErrorProvider is an interface that provides a method to return a uint8 value with an error.
type Uint8ErrorProvider interface {
	Uint8() (uint8, error)
}

// Uint8SliceProvider is an interface that provides a method to return a slice of uint8 values.
type Uint8SliceProvider interface {
	Uint8Slice() []uint8
}

// Uint8SliceErrorProvider is an interface that provides a method to return a slice of uint8 values with an error.
type Uint8SliceErrorProvider interface {
	Uint8Slice() ([]uint8, error)
}

// Uint16Provider is an interface that provides a method to return a uint16 value.
type Uint16Provider interface {
	Uint16() uint16
}

// Uint16ErrorProvider is an interface that provides a method to return a uint16 value with an error.
type Uint16ErrorProvider interface {
	Uint16() (uint16, error)
}

// Uint16SliceProvider is an interface that provides a method to return a slice of uint16 values.
type Uint16SliceProvider interface {
	Uint16Slice() []uint16
}

// Uint16SliceErrorProvider is an interface that provides a method to return a slice of uint16 values with an error.
type Uint16SliceErrorProvider interface {
	Uint16Slice() ([]uint16, error)
}

// Uint32Provider is an interface that provides a method to return a uint32 value.
type Uint32Provider interface {
	Uint32() uint32
}

// Uint32ErrorProvider is an interface that provides a method to return a uint32 value with an error.
type Uint32ErrorProvider interface {
	Uint32() (uint32, error)
}

// Uint32SliceProvider is an interface that provides a method to return a slice of uint32 values.
type Uint32SliceProvider interface {
	Uint32Slice() []uint32
}

// Uint32SliceErrorProvider is an interface that provides a method to return a slice of uint32 values with an error.
type Uint32SliceErrorProvider interface {
	Uint32Slice() ([]uint32, error)
}

// Uint64Provider is an interface that provides a method to return a uint64 value.
type Uint64Provider interface {
	Uint64() uint64
}

// Uint64ErrorProvider is an interface that provides a method to return a uint64 value with an error.
type Uint64ErrorProvider interface {
	Uint64() (uint64, error)
}

// Uint64SliceProvider is an interface that provides a method to return a slice of uint64 values.
type Uint64SliceProvider interface {
	Uint64Slice() []uint64
}

// Uint64SliceErrorProvider is an interface that provides a method to return a slice of uint64 values with an error.
type Uint64SliceErrorProvider interface {
	Uint64Slice() ([]uint64, error)
}

// Float32Provider is an interface that provides a method to return a float32 value.
type Float32Provider interface {
	Float32() float32
}

// Float32ErrorProvider is an interface that provides a method to return a float32 value with an error.
type Float32ErrorProvider interface {
	Float32() (float32, error)
}

// Float32SliceProvider is an interface that provides a method to return a slice of float32 values.
type Float32SliceProvider interface {
	Float32Slice() []float32
}

// Float32SliceErrorProvider is an interface that provides a method to return a slice of float32 values with an error.
type Float32SliceErrorProvider interface {
	Float32Slice() ([]float32, error)
}

// Float64Provider is an interface that provides a method to return a float64 value.
type Float64Provider interface {
	Float64() float64
}

// Float64ErrorProvider is an interface that provides a method to return a float64 value with an error.
type Float64ErrorProvider interface {
	Float64() (float64, error)
}

// Float64SliceProvider is an interface that provides a method to return a slice of float64 values.
type Float64SliceProvider interface {
	Float64Slice() []float64
}

// Float64SliceErrorProvider is an interface that provides a method to return a slice of float64 values with an error.
type Float64SliceErrorProvider interface {
	Float64Slice() ([]float64, error)
}

// StringProvider is an interface that provides a method to return a string value.
type StringProvider interface {
	String() string
}

// StringErrorProvider is an interface that provides a method to return a string value with an error.
type StringErrorProvider interface {
	String() (string, error)
}

// StringSliceProvider is an interface that provides a method to return a slice of string values.
type StringSliceProvider interface {
	StringSlice() []string
}

// StringSliceErrorProvider is an interface that provides a method to return a slice of string values with an error.
type StringSliceErrorProvider interface {
	StringSlice() ([]string, error)
}
