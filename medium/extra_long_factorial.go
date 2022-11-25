package medium

import "math/big"

func extraLongFactorials(n int32) string {
	longFact := longFactorials(big.NewInt(int64(n)))
	return longFact.String()
}

func longFactorials(n *big.Int) *big.Int {
	longFactorial := &big.Int{}

	switch n.Cmp(&big.Int{}) {
	case -1, 0:
		longFactorial.SetInt64(1)
	default:
		longFactorial.Set(n)
		one := big.NewInt(1)
		longFactorial.Mul(longFactorial, longFactorials(n.Sub(n, one)))
	}

	return longFactorial
}
