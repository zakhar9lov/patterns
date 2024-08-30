package takefirstn

import (
	"math/rand"
)

func GetRandomSliceInt(n int) []interface{} {
	slice := make([]interface{}, n)

	for i := 0; i < len(slice); i++ {
		slice[i] = rand.Intn(n)
	}

	return slice
}

func SendToChannel(slice []interface{}, ch chan<- interface{}) {
	defer close(ch)

	for _, v := range slice {
		ch <- v
	}

}
