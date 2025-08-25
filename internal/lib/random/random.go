package random

import (
	"math/rand"
	"time"
)

func NewRandowString(length int) string {

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	res := make([]rune, length)
	for i := 0; i < length; i++ {
		letter := rune('a' + r.Intn(26))
		res[i] = letter
	}

	return string(res)
}
