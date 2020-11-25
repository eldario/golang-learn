package hardMapper

import (
	"math/rand"
	"testing"
)

type testCase struct {
	word string
	value  int
	result []int
}


func BenchmarkInsert(b *testing.B) {
	rand.Seed(1)
	var data = New()
	for i := 0; i < b.N; i++ {
		data.Insert("foo.bar")
	}
}

func BenchmarkGetFrequentUses(b *testing.B) {
	rand.Seed(1)
	var data = New()
	for i := 0; i < b.N; i++ {
		data.GetFrequentUses()
	}
}
