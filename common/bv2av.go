package common

import (
	"errors"
)

var table = "fZodR9XQDSUm21yCkr6zBqiveYah8bt4xsWpHnJE7jL5VG3guMTKNPAwcF"
var tr = make(map[byte]int)

func init() {
	for i := 0; i < 58; i++ {
		tr[table[i]] = i
	}
}

var s = []int{11, 10, 3, 8, 4, 6}
var xor = 177451812
var add = 8728348608

func DecodeBV(x string) (int, error) {
	if len(x) != 12 {
		return 0, errors.New("Invalid input: Input length should be 12 characters")
	}

	r := 0
	for i := 0; i < 6; i++ {
		if idx, ok := tr[x[s[i]]]; ok {
			r += idx * pow(58, i)
		} else {
			return 0, errors.New("Invalid input: Contains invalid character")
		}
	}
	return (r - add) ^ xor, nil
}

func EncodeBV(x int) (string, error) {
	if x < 0 || x >= 58*58*58*58*58*58 {
		return "", errors.New("Invalid input: Input value out of range")
	}

	x = (x ^ xor) + add
	r := []byte("BV1  4 1 7  ")
	for i := 0; i < 6; i++ {
		r[s[i]] = table[x/pow(58, i)%58]
	}
	return string(r), nil
}

func pow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}
