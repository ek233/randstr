package randstr

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// Rand rand
type Rand struct {
	Pool *sync.Pool
}

var (
	seq      uint64 = 10000
	m               = NewRand()
	randlist        = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
)

// NewRand init Rand
func NewRand() *Rand {
	p := &sync.Pool{New: func() interface{} {
		return rand.New(rand.NewSource(getSeed()))
	},
	}
	mrand := &Rand{
		Pool: p,
	}
	return mrand
}

// get seed
func getSeed() int64 {
	//1592564536479936000
	seed := (atomic.AddUint64(&seq, 1)%1000 + 1000) * 1e15
	tn := time.Now().UnixNano() % 1e15
	return int64(seed) + tn
}

func (s *Rand) getrand() *rand.Rand {
	return s.Pool.Get().(*rand.Rand)
}
func (s *Rand) putrand(r *rand.Rand) {
	s.Pool.Put(r)
}

// Intn like math/rand.Intn
func (s *Rand) Intn(n int) int {
	r := s.getrand()
	defer s.putrand(r)

	return r.Intn(n)
}

// Read like math/rand.Read
func (s *Rand) Read(p []byte) (int, error) {
	r := s.getrand()
	defer s.putrand(r)

	return r.Read(p)
}

// RandomString creates a random string
func RandomString(len int) string {
	b := make([]byte, len)
	_, err := m.Read(b)
	if err != nil {
		return ""
	}
	for i := 0; i < len; i++ {
		b[i] = randlist[b[i]%(62)]
	}
	return *(*string)(unsafe.Pointer(&b))
}
