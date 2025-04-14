package helpers

import "math/rand"

func GenerateData() []int {
	rand.Seed(5)
	var d []int
	for i := 0; i < 10; i++ {
		d = append(d, rand.Intn(100))
	}
	return d
}
