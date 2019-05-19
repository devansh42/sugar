//@author Devansh Gupta
//This file contains code for implementation of normal type conversion functions used in go
package sugar

import (
	"fmt"
	"strconv"
)

func getint(x interface{}) int {
	y := fmt.Sprint(x)
	i, _ := strconv.Atoi(y)
	return i
}
func getfloat(x interface{}) float64 {
	y := fmt.Sprint(x)
	i, _ := strconv.ParseFloat(y, 64)
	return i
}

func InterfaceToFloat32(x interface{}) float32 {
	return float32(getfloat(x))
}
func InterfaceToFloat64(x interface{}) float64 {
	return getfloat(x)
}

func InterfaceToInt64(x interface{}) int64 {
	return int64(getint(x))

}
func InterfaceToInt32(x interface{}) int32 {
	return int32(getint(x))
}
func InterfaceToInt(x interface{}) int {
	return getint(x)
}
func InterfaceToInt8(x interface{}) int8 {
	return int8(getint(x))
}

func InterfaceToInt16(x interface{}) int16 {
	return int16(getint(x))
}
