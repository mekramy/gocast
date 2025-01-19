package gocast

import (
	"encoding/json"
	"fmt"
)

type casterDriver struct {
	data any
}

func (driver casterDriver) IsNil() bool {
	return valueOf(driver.data) == nil
}

func (driver casterDriver) Interface() any {
	return driver.data
}

func (driver casterDriver) Unmarshal(out any) error {
	// Try direct unmarshal
	err := json.Unmarshal([]byte(fmt.Sprintf("%v", valueOf(driver.data))), out)
	if err == nil {
		return nil
	}

	// Try marshal and unmarshal
	bytes, err := json.Marshal(driver.data)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, out)
}

func (driver casterDriver) Bool() (bool, error) {
	return ToBool(driver.data)
}

func (driver casterDriver) BoolSafe(fallback bool) bool {
	val, err := driver.Bool()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int() (int, error) {
	return ToSigned[int](driver.data)
}

func (driver casterDriver) IntSafe(fallback int) int {
	val, err := driver.Int()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int8() (int8, error) {
	return ToSigned[int8](driver.data)
}

func (driver casterDriver) Int8Safe(fallback int8) int8 {
	val, err := driver.Int8()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int16() (int16, error) {
	return ToSigned[int16](driver.data)
}

func (driver casterDriver) Int16Safe(fallback int16) int16 {
	val, err := driver.Int16()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int32() (int32, error) {
	return ToSigned[int32](driver.data)
}

func (driver casterDriver) Int32Safe(fallback int32) int32 {
	val, err := driver.Int32()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int64() (int64, error) {
	return ToSigned[int64](driver.data)
}

func (driver casterDriver) Int64Safe(fallback int64) int64 {
	val, err := driver.Int64()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint() (uint, error) {
	return ToUnsigned[uint](driver.data)
}

func (driver casterDriver) UintSafe(fallback uint) uint {
	val, err := driver.Uint()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint8() (uint8, error) {
	return ToUnsigned[uint8](driver.data)
}

func (driver casterDriver) Uint8Safe(fallback uint8) uint8 {
	val, err := driver.Uint8()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint16() (uint16, error) {
	return ToUnsigned[uint16](driver.data)
}

func (driver casterDriver) Uint16Safe(fallback uint16) uint16 {
	val, err := driver.Uint16()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint32() (uint32, error) {
	return ToUnsigned[uint32](driver.data)
}

func (driver casterDriver) Uint32Safe(fallback uint32) uint32 {
	val, err := driver.Uint32()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint64() (uint64, error) {
	return ToUnsigned[uint64](driver.data)
}

func (driver casterDriver) Uint64Safe(fallback uint64) uint64 {
	val, err := driver.Uint64()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Float32() (float32, error) {
	return ToFloat[float32](driver.data)
}

func (driver casterDriver) Float32Safe(fallback float32) float32 {
	val, err := driver.Float32()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Float64() (float64, error) {
	return ToFloat[float64](driver.data)
}

func (driver casterDriver) Float64Safe(fallback float64) float64 {
	val, err := driver.Float64()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) String() (string, error) {
	return ToString(driver.data)
}

func (driver casterDriver) StringSafe(fallback string) string {
	val, err := driver.String()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Slice() ([]any, error) {
	return ToSlice(driver.data)
}

func (driver casterDriver) SliceSafe(fallback []any) []any {
	if res, err := driver.Slice(); err == nil {
		return res
	}
	return fallback
}

func (driver casterDriver) BoolSlice() ([]bool, error) {
	return ToBoolSlice(driver.data)
}

func (driver casterDriver) BoolSliceSafe(fallback []bool) []bool {
	if res, err := driver.BoolSlice(); err == nil {
		return res
	}
	return fallback
}

func (driver casterDriver) IntSlice() ([]int, error) {
	return ToSignedSlice[int](driver.data)
}

func (driver casterDriver) IntSliceSafe(fallback []int) []int {
	val, err := driver.IntSlice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int8Slice() ([]int8, error) {
	return ToSignedSlice[int8](driver.data)

}

func (driver casterDriver) Int8SliceSafe(fallback []int8) []int8 {
	val, err := driver.Int8Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int16Slice() ([]int16, error) {
	return ToSignedSlice[int16](driver.data)
}

func (driver casterDriver) Int16SliceSafe(fallback []int16) []int16 {
	val, err := driver.Int16Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int32Slice() ([]int32, error) {
	return ToSignedSlice[int32](driver.data)
}

func (driver casterDriver) Int32SliceSafe(fallback []int32) []int32 {
	val, err := driver.Int32Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Int64Slice() ([]int64, error) {
	return ToSignedSlice[int64](driver.data)
}

func (driver casterDriver) Int64SliceSafe(fallback []int64) []int64 {
	val, err := driver.Int64Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) UintSlice() ([]uint, error) {
	return ToUnsignedSlice[uint](driver.data)
}

func (driver casterDriver) UintSliceSafe(fallback []uint) []uint {
	val, err := driver.UintSlice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint8Slice() ([]uint8, error) {
	return ToUnsignedSlice[uint8](driver.data)
}

func (driver casterDriver) Uint8SliceSafe(fallback []uint8) []uint8 {
	val, err := driver.Uint8Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint16Slice() ([]uint16, error) {
	return ToUnsignedSlice[uint16](driver.data)
}

func (driver casterDriver) Uint16SliceSafe(fallback []uint16) []uint16 {
	val, err := driver.Uint16Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint32Slice() ([]uint32, error) {
	return ToUnsignedSlice[uint32](driver.data)
}

func (driver casterDriver) Uint32SliceSafe(fallback []uint32) []uint32 {
	val, err := driver.Uint32Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Uint64Slice() ([]uint64, error) {
	return ToUnsignedSlice[uint64](driver.data)
}

func (driver casterDriver) Uint64SliceSafe(fallback []uint64) []uint64 {
	val, err := driver.Uint64Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Float32Slice() ([]float32, error) {
	return ToFloatSlice[float32](driver.data)
}

func (driver casterDriver) Float32SliceSafe(fallback []float32) []float32 {
	val, err := driver.Float32Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) Float64Slice() ([]float64, error) {
	return ToFloatSlice[float64](driver.data)
}

func (driver casterDriver) Float64SliceSafe(fallback []float64) []float64 {
	val, err := driver.Float64Slice()
	if err != nil {
		return fallback
	}
	return val
}

func (driver casterDriver) StringSlice() ([]string, error) {
	return ToStringSlice(driver.data)
}

func (driver casterDriver) StringSliceSafe(fallback []string) []string {
	val, err := driver.StringSlice()
	if err != nil {
		return fallback
	}
	return val
}
