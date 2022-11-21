package random

import (
	"math/rand"
	"time"
)

var seed = time.Now().UnixNano()

func SetSeed(seed_to_set int64) {
	rand.Seed(seed_to_set)
	seed = seed_to_set
}

func GetSeed() int64 {
	return seed
}

func RangeInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}
