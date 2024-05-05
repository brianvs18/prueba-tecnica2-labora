package common

import "math/rand"

func GetRandomObject[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}
