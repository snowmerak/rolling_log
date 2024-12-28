package blake3

import (
	"lukechampine.com/blake3"

	"github.com/snowmerak/rolling_log/lib/client/hasher"
)

type Blake256 struct{}

func NewBlake256() *Blake256 {
	return &Blake256{}
}

var _ hasher.Hasher = (*Blake256)(nil)

func (s *Blake256) Hash(data []byte) []byte {
	hashed := blake3.Sum256(data)
	return hashed[:]
}

func (s *Blake256) HashSize() int {
	return 32
}

type Blake512 struct{}

func NewBlake512() *Blake512 {
	return &Blake512{}
}

var _ hasher.Hasher = (*Blake512)(nil)

func (s *Blake512) Hash(data []byte) []byte {
	hashed := blake3.Sum512(data)
	return hashed[:]
}

func (s *Blake512) HashSize() int {
	return 64
}

type Blake3 struct {
	size int
	key  []byte
}

func NewBlake3(size int, key []byte) *Blake3 {
	switch {
	case len(key) == 0:
		key = nil
	case len(key) < 32:
		key = append(key, make([]byte, 32-len(key))...)
	case len(key) > 32:
		key = key[:32]
	}
	return &Blake3{
		size: size,
		key:  key,
	}
}

var _ hasher.Hasher = (*Blake3)(nil)

func (s *Blake3) Hash(data []byte) []byte {
	hs := blake3.New(s.size, s.key)
	hs.Write(data)
	return hs.Sum(nil)
}

func (s *Blake3) HashSize() int {
	return s.size
}
