package utils

import "strconv"

func StrToFloat32(f string) (floatString float32, err error) {
	s, err := strconv.ParseFloat(f, 32)
	if err == nil {
		floatString = float32(s)
	}
	return
}

func StrToInt(str string) (value int, err error) {
	value, err = strconv.Atoi(str)
	return
}

func Abs(x float32) float32 {
	if x < 0 {
		return -x
	}
	return x
}
