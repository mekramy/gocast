package gocast_test

import (
	"reflect"
	"testing"

	"github.com/mekramy/gocast"
)

func TestToBool(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected bool
		err      bool
	}{
		{true, true, false},
		{false, false, false},
		{1, true, false},
		{0, false, false},
		{"true", true, false},
		{"false", false, false},
		{"invalid", false, true},
		{nil, false, true},
	}

	for _, test := range tests {
		result, err := gocast.ToBool(test.input)
		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToBool(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if result != test.expected {
			t.Errorf("ToBool(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToSigned(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected int
		err      bool
	}{
		{true, 1, false},
		{false, 0, false},
		{1, 1, false},
		{0, 0, false},
		{"123", 123, false},
		{"invalid", 0, true},
		{nil, 0, true},
	}

	for _, test := range tests {
		result, err := gocast.ToSigned[int](test.input)
		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToSigned(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if result != test.expected {
			t.Errorf("ToSigned(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToUnsigned(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected uint
		err      bool
	}{
		{true, 1, false},
		{false, 0, false},
		{1, 1, false},
		{-1, 0, true},
		{"123", 123, false},
		{"invalid", 0, true},
		{nil, 0, true},
	}

	for _, test := range tests {
		result, err := gocast.ToUnsigned[uint](test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToUnsigned(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if result != test.expected {
			t.Errorf("ToUnsigned(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToFloat(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected float64
		err      bool
	}{
		{true, 1.0, false},
		{false, 0.0, false},
		{1, 1.0, false},
		{0, 0.0, false},
		{"123.45", 123.45, false},
		{"invalid", 0.0, true},
		{nil, 0.0, true},
	}

	for _, test := range tests {
		result, err := gocast.ToFloat[float64](test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToFloat(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if result != test.expected {
			t.Errorf("ToFloat(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
		err      bool
	}{
		{true, "true", false},
		{false, "false", false},
		{1, "1", false},
		{0, "0", false},
		{"hello", "hello", false},
		{nil, "", true},
	}

	for _, test := range tests {
		result, err := gocast.ToString(test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToString(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if result != test.expected {
			t.Errorf("ToString(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []interface{}
		err      bool
	}{
		{[]interface{}{1, 2, 3}, []interface{}{1, 2, 3}, false},
		{[]map[string]interface{}{{"key": "value"}}, []interface{}{map[string]interface{}{"key": "value"}}, false},
		{"invalid", []interface{}{}, true},
	}

	for _, test := range tests {
		result, err := gocast.ToSlice(test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToSlice(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ToSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToBoolSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []bool
		err      bool
	}{
		{[]bool{true, false}, []bool{true, false}, false},
		{[]interface{}{true, false}, []bool{true, false}, false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := gocast.ToBoolSlice(test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToBoolSlice(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ToBoolSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToSignedSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []int
		err      bool
	}{
		{[]int{1, 2, 3}, []int{1, 2, 3}, false},
		{[]interface{}{1, 2, 3}, []int{1, 2, 3}, false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := gocast.ToSignedSlice[int](test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToSignedSlice(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ToSignedSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToUnsignedSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []uint
		err      bool
	}{
		{[]uint{1, 2, 3}, []uint{1, 2, 3}, false},
		{[]interface{}{1, 2, 3}, []uint{1, 2, 3}, false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := gocast.ToUnsignedSlice[uint](test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToUnsignedSlice(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ToUnsignedSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToFloatSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []float64
		err      bool
	}{
		{[]float64{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}, false},
		{[]interface{}{1.1, 2.2, 3.3}, []float64{1.1, 2.2, 3.3}, false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := gocast.ToFloatSlice[float64](test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToFloatSlice(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ToFloatSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func TestToStringSlice(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected []string
		err      bool
	}{
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, false},
		{[]interface{}{"a", "b", "c"}, []string{"a", "b", "c"}, false},
		{"invalid", nil, true},
	}

	for _, test := range tests {
		result, err := gocast.ToStringSlice(test.input)

		if test.err && err != nil {
			continue
		}

		if (err != nil) != test.err {
			t.Errorf("ToStringSlice(%v) error = %v, expected error = %v", test.input, err, test.err)
		}

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ToStringSlice(%v) = %v, expected %v", test.input, result, test.expected)
		}
	}
}
