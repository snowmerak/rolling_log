package mldsa

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/circl/sign/mldsa/mldsa87"

	"github.com/snowmerak/rolling_log/lib/client/signer"
)

type Signer87 struct {
	privateKey *mldsa87.PrivateKey
}

func NewSigner87() (*Signer87, error) {
	_, privateKey, err := mldsa87.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}

	return &Signer87{privateKey: privateKey}, nil
}

var _ signer.Signer = (*Signer87)(nil)

func (s *Signer87) Sign(data []byte) ([]byte, error) {
	return s.privateKey.Sign(rand.Reader, data, crypto.Hash(0))
}

func (s *Signer87) Marshal() ([]byte, error) {
	return s.privateKey.Bytes(), nil
}

func (s *Signer87) Verifier() (signer.Verifier, error) {
	pubKey, ok := s.privateKey.Public().(*mldsa87.PublicKey)
	if !ok {
		return nil, signer.ErrInvalidPublicKey
	}

	return &Verifier87{publicKey: pubKey}, nil
}

func (s *Signer87) Unmarshal(data []byte) (signer.Signer, error) {
	privateKey := new(mldsa87.PrivateKey)
	if err := privateKey.UnmarshalBinary(data); err != nil {
		return nil, fmt.Errorf("unmarshal private key: %w", err)
	}

	return &Signer87{privateKey: privateKey}, nil
}

type Verifier87 struct {
	publicKey *mldsa87.PublicKey
}

var _ signer.Verifier = (*Verifier87)(nil)

func (v *Verifier87) Verify(data []byte, signature []byte) (bool, error) {
	return v.publicKey.Scheme().Verify(v.publicKey, data, signature, nil), nil
}

func (v *Verifier87) Marshal() ([]byte, error) {
	return v.publicKey.Bytes(), nil
}

func (v *Verifier87) Unmarshal(data []byte) (signer.Verifier, error) {
	publicKey := new(mldsa87.PublicKey)
	if err := publicKey.UnmarshalBinary(data); err != nil {
		return nil, fmt.Errorf("unmarshal public key: %w", err)
	}

	return &Verifier87{publicKey: publicKey}, nil
}
