package gocast

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToBool casts an interface to a bool type.
func ToBool(value interface{}) (bool, error) {
	value = valueOf(value)
	switch val := value.(type) {
	case nil:
		return false, nilErr()
	case BoolErrorProvider:
		return val.Bool()
	case BoolProvider:
		return val.Bool(), nil
	case bool:
		return val, nil
	case int:
		return val != 0, nil
	case int8:
		return val != 0, nil
	case int16:
		return val != 0, nil
	case int32:
		return val != 0, nil
	case int64:
		return val != 0, nil
	case uint:
		return val != 0, nil
	case uint8:
		return val != 0, nil
	case uint16:
		return val != 0, nil
	case uint32:
		return val != 0, nil
	case uint64:
		return val != 0, nil
	case float32:
		return val != 0, nil
	case float64:
		return val != 0, nil
	case string:
		v, err := strconv.ParseBool(val)
		if err != nil {
			return false, typeError("bool")
		}
		return v, nil
	default:
		v, err := strconv.ParseBool(fmt.Sprintf("%v", value))
		if err != nil {
			return false, typeError("bool")
		}
		return v, nil
	}
}

// ToSigned casts an interface to a signed integer type.
func ToSigned[T int | int8 | int16 | int32 | int64](value interface{}) (T, error) {
	value = valueOf(value)
	msg := typeError(typeName[T]())

	// Check provider
	if isImplementsOf[T, int]() {
		switch val := value.(type) {
		case IntErrorProvider:
			v, err := val.Int()
			return T(v), err
		case IntProvider:
			return T(val.Int()), nil
		}
	} else if isImplementsOf[T, int8]() {
		switch val := value.(type) {
		case Int8ErrorProvider:
			v, err := val.Int8()
			return T(v), err
		case Int8Provider:
			return T(val.Int8()), nil
		}
	} else if isImplementsOf[T, int16]() {
		switch val := value.(type) {
		case Int16ErrorProvider:
			v, err := val.Int16()
			return T(v), err
		case Int16Provider:
			return T(val.Int16()), nil
		}
	} else if isImplementsOf[T, int32]() {
		switch val := value.(type) {
		case Int32ErrorProvider:
			v, err := val.Int32()
			return T(v), err
		case Int32Provider:
			return T(val.Int32()), nil
		}
	} else if isImplementsOf[T, int64]() {
		switch val := value.(type) {
		case Int64ErrorProvider:
			v, err := val.Int64()
			return T(v), err
		case Int64Provider:
			return T(val.Int64()), nil
		}
	}

	// Cast
	switch val := value.(type) {
	case nil:
		return 0, nilErr()
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case int:
		return T(val), nil
	case int8:
		return T(val), nil
	case int16:
		return T(val), nil
	case int32:
		return T(val), nil
	case int64:
		return T(val), nil
	case uint:
		return T(val), nil
	case uint8:
		return T(val), nil
	case uint16:
		return T(val), nil
	case uint32:
		return T(val), nil
	case uint64:
		return T(val), nil
	case float32:
		return T(val), nil
	case float64:
		return T(val), nil
	case string:
		i, err := strconv.ParseInt(val, 0, 0)
		if err == nil {
			return T(i), nil
		}

		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return T(f), nil
		}

		return 0, msg
	default:
		i, err := strconv.ParseInt(fmt.Sprintf("%v", val), 0, 0)
		if err == nil {
			return T(i), nil
		}

		f, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 64)
		if err == nil {
			return T(f), nil
		}

		return 0, msg
	}
}

// ToUnsigned casts an interface to a unsigned integer type.
func ToUnsigned[T uint | uint8 | uint16 | uint32 | uint64](value interface{}) (T, error) {
	value = valueOf(value)
	msg := typeError(typeName[T]())

	// Check provider
	if isImplementsOf[T, uint]() {
		switch val := value.(type) {
		case UintErrorProvider:
			v, err := val.Uint()
			return T(v), err
		case UintProvider:
			return T(val.Uint()), nil
		}
	} else if isImplementsOf[T, uint8]() {
		switch val := value.(type) {
		case Uint8ErrorProvider:
			v, err := val.Uint8()
			return T(v), err
		case Uint8Provider:
			return T(val.Uint8()), nil
		}
	} else if isImplementsOf[T, uint16]() {
		switch val := value.(type) {
		case Uint16ErrorProvider:
			v, err := val.Uint16()
			return T(v), err
		case Uint16Provider:
			return T(val.Uint16()), nil
		}
	} else if isImplementsOf[T, uint32]() {
		switch val := value.(type) {
		case Uint32ErrorProvider:
			v, err := val.Uint32()
			return T(v), err
		case Uint32Provider:
			return T(val.Uint32()), nil
		}
	} else if isImplementsOf[T, uint64]() {
		switch val := value.(type) {
		case Uint64ErrorProvider:
			v, err := val.Uint64()
			return T(v), err
		case Uint64Provider:
			return T(val.Uint64()), nil
		}
	}

	// Cast
	switch val := value.(type) {
	case nil:
		return 0, nilErr()
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case int:
		if val < 0 {
			return 0, msg
		}
		return T(val), nil
	case int8:
		if val < 0 {
			return 0, msg
		}
		return T(val), nil
	case int16:
		if val < 0 {
			return 0, msg
		}
		return T(val), nil
	case int32:
		if val < 0 {
			return 0, msg
		}
		return T(val), nil
	case int64:
		if val < 0 {
			return 0, msg
		}
		return T(val), nil
	case uint:
		return T(val), nil
	case uint8:
		return T(val), nil
	case uint16:
		return T(val), nil
	case uint32:
		return T(val), nil
	case uint64:
		return T(val), nil
	case float32:
		if val < 0 {
			return 0, msg
		}
		return T(val), nil
	case float64:
		if val < 0 {
			return 0, msg
		}
		return T(val), nil
	case string:
		i, err := strconv.ParseInt(val, 0, 0)
		if i < 0 {
			return 0, msg
		} else if err == nil {
			return T(i), nil
		}

		f, err := strconv.ParseFloat(val, 64)
		if f < 0 {
			return 0, msg
		} else if err == nil {
			return T(f), nil
		}

		return 0, msg
	default:
		i, err := strconv.ParseInt(fmt.Sprintf("%v", val), 0, 0)
		if i < 0 {
			return 0, msg
		} else if err == nil {
			return T(i), nil
		}

		f, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 64)
		if f < 0 {
			return 0, msg
		} else if err == nil {
			return T(f), nil
		}

		return 0, msg
	}
}

// ToFloat casts an interface to a float type.
func ToFloat[T float32 | float64](value interface{}) (T, error) {
	value = valueOf(value)
	msg := typeError(typeName[T]())

	// Check provider
	if isImplementsOf[T, float32]() {
		switch val := value.(type) {
		case Float32ErrorProvider:
			v, err := val.Float32()
			return T(v), err
		case Float32Provider:
			return T(val.Float32()), nil
		}
	} else if isImplementsOf[T, float64]() {
		switch val := value.(type) {
		case Float64ErrorProvider:
			v, err := val.Float64()
			return T(v), err
		case Float64Provider:
			return T(val.Float64()), nil
		}
	}

	// Cast
	switch val := value.(type) {
	case nil:
		return 0, nilErr()
	case bool:
		if val {
			return 1, nil
		}
		return 0, nil
	case int:
		return T(val), nil
	case int8:
		return T(val), nil
	case int16:
		return T(val), nil
	case int32:
		return T(val), nil
	case int64:
		return T(val), nil
	case uint:
		return T(val), nil
	case uint8:
		return T(val), nil
	case uint16:
		return T(val), nil
	case uint32:
		return T(val), nil
	case uint64:
		return T(val), nil
	case float32:
		return T(val), nil
	case float64:
		return T(val), nil
	case string:
		f, err := strconv.ParseFloat(val, 64)
		if err == nil {
			return T(f), nil
		}

		i, err := strconv.ParseInt(val, 0, 0)
		if err == nil {
			return T(i), nil
		}

		return 0, msg
	default:
		f, err := strconv.ParseFloat(fmt.Sprintf("%v", val), 64)
		if err == nil {
			return T(f), nil
		}

		i, err := strconv.ParseInt(fmt.Sprintf("%v", val), 0, 0)
		if err == nil {
			return T(i), nil
		}

		return 0, msg
	}
}

// ToString casts an interface to a string type.
func ToString(value interface{}) (string, error) {
	value = valueOf(value)
	switch val := value.(type) {
	case nil:
		return "", nilErr()
	case StringErrorProvider:
		return val.String()
	case StringProvider:
		return val.String(), nil
	case bool:
		return strconv.FormatBool(val), nil
	case int:
		return strconv.FormatInt(int64(val), 10), nil
	case int8:
		return strconv.FormatInt(int64(val), 10), nil
	case int16:
		return strconv.FormatInt(int64(val), 10), nil
	case int32:
		return strconv.FormatInt(int64(val), 10), nil
	case int64:
		return strconv.FormatInt(val, 10), nil
	case uint:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint32:
		return strconv.FormatUint(uint64(val), 10), nil
	case uint64:
		return strconv.FormatUint(uint64(val), 10), nil
	case float32:
		return strconv.FormatUint(uint64(val), 10), nil
	case float64:
		return strconv.FormatUint(uint64(val), 10), nil
	case string:
		return val, nil
	default:
		return "", typeError("bool")
	}
}

// ToSlice casts an interface{} to a []interface{} type.
func ToSlice(value interface{}) ([]interface{}, error) {
	var res []interface{}

	switch val := value.(type) {
	case []interface{}:
		return append(res, val...), nil
	case []map[string]interface{}:
		for _, u := range val {
			res = append(res, u)
		}
		return res, nil
	default:
		return res, typeError("[]interface{}")
	}
}

// ToBoolSlice casts an interface to a []bool type.
func ToBoolSlice(i interface{}) ([]bool, error) {
	if i == nil {
		return []bool{}, nilErr()
	}

	switch v := i.(type) {
	case []bool:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]bool, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToBool(s.Index(j).Interface())
			if err != nil {
				return []bool{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []bool{}, typeError("[]bool")
	}
}

// ToSignedSlice casts an interface to a signed integer slice type.
func ToSignedSlice[T int | int8 | int16 | int32 | int64](i interface{}) ([]T, error) {
	if i == nil {
		return []T{}, nilErr()
	}

	switch v := i.(type) {
	case []T:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]T, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToSigned[T](s.Index(j).Interface())
			if err != nil {
				return []T{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []T{}, typeError("[]" + typeName[T]())
	}
}

// ToUnsignedSlice casts an interface to a unsigned integer slice type.
func ToUnsignedSlice[T uint | uint8 | uint16 | uint32 | uint64](i interface{}) ([]T, error) {
	if i == nil {
		return []T{}, nilErr()
	}

	switch v := i.(type) {
	case []T:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]T, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToUnsigned[T](s.Index(j).Interface())
			if err != nil {
				return []T{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []T{}, typeError("[]" + typeName[T]())
	}
}

// ToFloatSlice casts an interface to a float slice type.
func ToFloatSlice[T float32 | float64](i interface{}) ([]T, error) {
	if i == nil {
		return []T{}, nilErr()
	}

	switch v := i.(type) {
	case []T:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]T, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToFloat[T](s.Index(j).Interface())
			if err != nil {
				return []T{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []T{}, typeError("[]" + typeName[T]())
	}
}

// ToStringSlice casts an interface to a []string type.
func ToStringSlice(i interface{}) ([]string, error) {
	if i == nil {
		return []string{}, nilErr()
	}

	switch v := i.(type) {
	case []string:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]string, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := ToString(s.Index(j).Interface())
			if err != nil {
				return []string{}, err
			}
			a[j] = val
		}
		return a, nil
	default:
		return []string{}, typeError("[]string")
	}
}
