package sha3

import (
	"crypto/sha3"

	"github.com/snowmerak/rolling_log/lib/client/hasher"
)

type Sha256 struct{}

func NewSha256() *Sha256 {
	return &Sha256{}
}

var _ hasher.Hasher = (*Sha256)(nil)

func (s *Sha256) Hash(data []byte) []byte {
	hs := sha3.New256()
	hs.Write(data)
	return hs.Sum(nil)
}

func (s *Sha256) HashSize() int {
	return sha3.New256().Size()
}

type Sha384 struct{}

func NewSha384() *Sha384 {
	return &Sha384{}
}

var _ hasher.Hasher = (*Sha384)(nil)

func (s *Sha384) Hash(data []byte) []byte {
	hs := sha3.New384()
	hs.Write(data)
	return hs.Sum(nil)
}

func (s *Sha384) HashSize() int {
	return sha3.New384().Size()
}

type Sha512 struct{}

func NewSha512() *Sha512 {
	return &Sha512{}
}

var _ hasher.Hasher = (*Sha512)(nil)

func (s *Sha512) Hash(data []byte) []byte {
	hs := sha3.New512()
	hs.Write(data)
	return hs.Sum(nil)
}

func (s *Sha512) HashSize() int {
	return sha3.New512().Size()
}
