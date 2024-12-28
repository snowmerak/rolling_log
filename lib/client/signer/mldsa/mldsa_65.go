package mldsa

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/circl/sign/mldsa/mldsa65"

	"github.com/snowmerak/rolling_log/lib/client/signer"
)

type Signer65 struct {
	privateKey *mldsa65.PrivateKey
}

func NewSigner65() (*Signer65, error) {
	_, privateKey, err := mldsa65.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}

	return &Signer65{privateKey: privateKey}, nil
}

var _ signer.Signer = (*Signer65)(nil)

func (s *Signer65) Sign(data []byte) ([]byte, error) {
	return s.privateKey.Sign(rand.Reader, data, crypto.Hash(0))
}

func (s *Signer65) Marshal() ([]byte, error) {
	return s.privateKey.Bytes(), nil
}

func (s *Signer65) Verifier() (signer.Verifier, error) {
	pubKey, ok := s.privateKey.Public().(*mldsa65.PublicKey)
	if !ok {
		return nil, signer.ErrInvalidPublicKey
	}

	return &Verifier65{publicKey: pubKey}, nil
}

func (s *Signer65) Unmarshal(data []byte) (signer.Signer, error) {
	privateKey := new(mldsa65.PrivateKey)
	if err := privateKey.UnmarshalBinary(data); err != nil {
		return nil, fmt.Errorf("unmarshal private key: %w", err)
	}

	return &Signer65{privateKey: privateKey}, nil
}

type Verifier65 struct {
	publicKey *mldsa65.PublicKey
}

var _ signer.Verifier = (*Verifier65)(nil)

func (v *Verifier65) Verify(data []byte, signature []byte) (bool, error) {
	return v.publicKey.Scheme().Verify(v.publicKey, data, signature, nil), nil
}

func (v *Verifier65) Marshal() ([]byte, error) {
	return v.publicKey.Bytes(), nil
}

func (v *Verifier65) Unmarshal(data []byte) (signer.Verifier, error) {
	publicKey := new(mldsa65.PublicKey)
	if err := publicKey.UnmarshalBinary(data); err != nil {
		return nil, fmt.Errorf("unmarshal public key: %w", err)
	}

	return &Verifier65{publicKey: publicKey}, nil
}