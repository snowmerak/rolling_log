package blake3_test

import (
	"testing"

	"github.com/snowmerak/rolling_log/lib/client/hasher/blake3"
)

func TestBlake256_Hash(t *testing.T) {
	hs := blake3.NewBlake256()
	hashed := hs.Hash([]byte("hello world"))
	if len(hashed) != hs.HashSize() {
		t.Errorf("invalid hash size: current (%d), expected (%d)", len(hashed), hs.HashSize())
	}
}

func TestBlake512_Hash(t *testing.T) {
	hs := blake3.NewBlake512()
	hashed := hs.Hash([]byte("hello world"))
	if len(hashed) != hs.HashSize() {
		t.Errorf("invalid hash size: current (%d), expected (%d)", len(hashed), hs.HashSize())
	}
}

func TestBlake3_Hash(t *testing.T) {
	hs := blake3.NewBlake3(64, nil)
	hashed := hs.Hash([]byte("hello world"))
	if len(hashed) != hs.HashSize() {
		t.Errorf("invalid hash size: current (%d), expected (%d)", len(hashed), hs.HashSize())
	}
}
