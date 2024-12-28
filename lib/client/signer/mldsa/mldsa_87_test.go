package mldsa_test

import (
	"bytes"
	"crypto/rand"
	"testing"

	"github.com/snowmerak/rolling_log/lib/client/signer/mldsa"
)

func TestSigner87_SignAndVerify(t *testing.T) {
	pvk, err := mldsa.NewSigner87()
	if err != nil {
		t.Fatalf("NewSigner44: %v", err)
	}

	data := make([]byte, 32)
	rand.Read(data)

	signature, err := pvk.Sign(data)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	pbk, err := pvk.Verifier()
	if err != nil {
		t.Fatalf("Verifier: %v", err)
	}

	ok, err := pbk.Verify(data, signature)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}

	if !ok {
		t.Fatalf("Verify: failed")
	}
}

func TestSigner87_MarshalAndUnmarshal(t *testing.T) {
	pvk, err := mldsa.NewSigner87()
	if err != nil {
		t.Fatalf("NewSigner44: %v", err)
	}

	input := make([]byte, 32)
	rand.Read(input)

	signature, err := pvk.Sign(input)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	data, err := pvk.Marshal()
	if err != nil {
		t.Fatalf("Marshal: %v", err)
	}

	pvk2, err := new(mldsa.Signer87).Unmarshal(data)
	if err != nil {
		t.Fatalf("UnmarshalSigner44: %v", err)
	}

	signature2, err := pvk2.Sign(input)
	if err != nil {
		t.Fatalf("Sign: %v", err)
	}

	if !bytes.Equal(signature, signature2) {
		t.Fatalf("Signatures do not match")
	}

	pbk, err := pvk2.Verifier()
	if err != nil {
		t.Fatalf("Verifier: %v", err)
	}

	ok, err := pbk.Verify(input, signature)
	if err != nil {
		t.Fatalf("Verify: %v", err)
	}

	if !ok {
		t.Fatalf("Verify: failed")
	}
}
