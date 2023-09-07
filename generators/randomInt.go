package randomint

import (
	"math/rand"
)

func GenerateRandomInt(amount int) []int {
	arr := make([]int, amount)
	for i := 0; i < amount; i++ {
		randomNumber := rand.Intn(10_000)
		arr[i] = randomNumber
	}
	return arr
}
