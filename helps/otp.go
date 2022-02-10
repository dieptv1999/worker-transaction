package helps

import (
	"crypto/rand"
	"log"
	"math/big"
)

func RandOTP6Digits() int64 {
	max := big.NewInt(999999)
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		log.Printf("[RandOTP6Digits] = %#v \n", err)
		return 0
	}
	return n.Int64()
}
