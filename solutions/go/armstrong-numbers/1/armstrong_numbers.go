package armstrongnumbers

import (
	"math/big"
	"strconv"
)

func IsNumber(n int) bool {
	number := strconv.Itoa(n)
	length := len(number)
	result := 0

	for _, v := range number {
		num, err := strconv.Atoi(string(v))
		if err != nil {
			panic("ops, this is not a number!")
		}
		result += pow(num, length)
	}
	return result == n
}

func pow(base int, exp int) int {
	b := big.NewInt(int64(base))
	e := big.NewInt(int64(exp))
	result := new(big.Int).Exp(b, e, nil)

	return int(result.Int64())
}
