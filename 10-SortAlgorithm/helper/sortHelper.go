package helper

import (
	"math/rand"
	"time"
)

func GetRandomSlice(sliceLen int) []int {
	var randomSlice []int
	newRandSeed:=rand.NewSource(time.Now().Unix())
	newRand:=rand.New(newRandSeed)
	for i := 0; i < sliceLen; i++ {
		randomSlice = append(randomSlice, newRand.Int()%100)
	}
	return randomSlice
}
