package helps

import (
	"fmt"
	"strconv"
)

func PaddingZeros(num int64) string {
	s := fmt.Sprintf("%019d", num)
	return s
}

func PaddingZerosStr(num string) string {
	s := fmt.Sprintf("%019s", num)
	return s
}

func PaddingZerosIf(num interface{}) string {
	s := fmt.Sprintf("%019d", num)
	return s
}

func PaddingZerosFloat(num float64) string {
	s := fmt.Sprintf("%030.10f", num)
	return s
}

func PaddingZerosUint64(num uint64) string {
	s := fmt.Sprintf("%019d", num)
	return s
}

func S2I(s string) int64 {
	r, _ := strconv.ParseInt(s, 10, 64)
	return r
}

func PaddingZerosBigInt78(s string) string {
	return fmt.Sprintf("%078s", s)
}

//func PaddingZeroKeys(keys... string) {
//	for _, key := range keys {
//
//	}
//}
