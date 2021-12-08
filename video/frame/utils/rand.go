package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func NewRand(low int, high int) int {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return low + rand.Intn(high)%(high-low+1)
}

func CreateFourRand() string {
	return fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))
}
