package helpers

import "math/rand"

func RandomValue() int {
	value := rand.Intn(100)
	return value
}