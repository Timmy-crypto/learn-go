package algorithm

import (
	"fmt"
	"math"
	"testing"
)

func Atoi(s string) int64 {
	sLen := len(s)
	var r int64

	for i := 0; i < sLen; i++ {
		//	fmt.Println("the s[i] is:", s[i], s[i]-0x30, math.Pow10(sLen-i-1))
		r += (int64)(s[i]-0x30) * int64(math.Pow10(sLen-i-1))
		//	fmt.Println("the r is:", r)
	}
	return r
}

func Test_Atoi(t *testing.T) {
	s := "1234"

	r := Atoi(s)
	fmt.Println("the r is:", r)
}
