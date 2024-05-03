package getrand

import (
	"math/rand"
	"strconv"
	"time"
)

func GetRand(stringQuantity, numberQuantity int) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	randomString := []rune{}
	alfa := []rune("abcdefghijklmnopqrstuvwxyz")

	for i := 0; i < r.Intn(3)+stringQuantity-2; i++ {
		randomString = append(randomString, alfa[r.Intn(len(alfa))])
	}
	for i := 0; i < r.Intn(3)+numberQuantity-2; i++ {
		randomString = append(randomString, rune(strconv.Itoa(r.Intn(10))[0]))
	}

	return string(randomString)
}
