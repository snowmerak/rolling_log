package sha3_test

import (
	"testing"

	"github.com/snowmerak/rolling_log/lib/client/hasher/sha3"
)

func TestSha256_Hash(t *testing.T) {
	h := sha3.NewSha256()
	data := []byte("hello world")

	hashed := h.Hash(data)
	if len(hashed) != h.HashSize() {
		t.Errorf("invalid hash size: current (%d), expected (%d)", len(hashed), h.HashSize())
	}
}

func TestSha384_Hash(t *testing.T) {
	h := sha3.NewSha384()
	data := []byte("hello world")

	hashed := h.Hash(data)
	if len(hashed) != h.HashSize() {
		t.Errorf("invalid hash size: current (%d), expected (%d)", len(hashed), h.HashSize())
	}
}

func TestSha512_Hash(t *testing.T) {
	h := sha3.NewSha512()
	data := []byte("hello world")

	hashed := h.Hash(data)
	if len(hashed) != h.HashSize() {
		t.Errorf("invalid hash size: current (%d), expected (%d)", len(hashed), h.HashSize())
	}
}
