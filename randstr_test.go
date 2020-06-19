package randstr

import "testing"

func TestRandString(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(RandomString(16))
	}
}

func BenchmarkRandString(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			RandomString(16)
		}
	})
}
