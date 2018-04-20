package numconvert

import (
	"errors"
	"math"
)

const (
	baseString = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func parse(number int64, hex int64, cpt []int64) []int64 {
	tail := number % hex
	cpt = append(cpt, tail)
	integ := math.Floor(float64(number) / float64(hex))
	if integ > 0 {
		return parse(int64(integ), hex, cpt)
	} else {
		return cpt
	}
}

func Convert(number int64, radix int64) (string, error) {
	if radix > int64(len(baseString)) {
		return "", errors.New("Over maximum radix")
	}
	parsed := parse(number, radix, []int64{})
	str := ""
	for i := len(parsed); i > 0; i-- {
		idx := parsed[i-1]
		str += baseString[idx : idx+1]
	}
	return str, nil
}
