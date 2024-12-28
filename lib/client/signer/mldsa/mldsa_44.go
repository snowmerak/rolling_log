package mldsa

import (
	"crypto"
	"crypto/rand"
	"fmt"

	"github.com/cloudflare/circl/sign/mldsa/mldsa44"

	"github.com/snowmerak/rolling_log/lib/client/signer"
)

type Signer44 struct {
	privateKey *mldsa44.PrivateKey
}

func NewSigner44() (*Signer44, error) {
	_, privateKey, err := mldsa44.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("generate key: %w", err)
	}

	return &Signer44{privateKey: privateKey}, nil
}

var _ signer.Signer = (*Signer44)(nil)

func (s *Signer44) Sign(data []byte) ([]byte, error) {
	return s.privateKey.Sign(rand.Reader, data, crypto.Hash(0))
}

func (s *Signer44) Marshal() ([]byte, error) {
	return s.privateKey.Bytes(), nil
}

func (s *Signer44) Verifier() (signer.Verifier, error) {
	pubKey, ok := s.privateKey.Public().(*mldsa44.PublicKey)
	if !ok {
		return nil, signer.ErrInvalidPublicKey
	}
	return &Verifier44{publicKey: pubKey}, nil
}

func (s *Signer44) Unmarshal(data []byte) (signer.Signer, error) {
	privateKey := new(mldsa44.PrivateKey)
	if err := privateKey.UnmarshalBinary(data); err != nil {
		return nil, fmt.Errorf("unmarshal private key: %w", err)
	}

	return &Signer44{privateKey: privateKey}, nil
}

type Verifier44 struct {
	publicKey *mldsa44.PublicKey
}

var _ signer.Verifier = (*Verifier44)(nil)

func (v *Verifier44) Verify(data []byte, signature []byte) (bool, error) {
	return v.publicKey.Scheme().Verify(v.publicKey, data, signature, nil), nil
}

func (v *Verifier44) Marshal() ([]byte, error) {
	return v.publicKey.Bytes(), nil
}

func (v *Verifier44) Unmarshal(data []byte) (signer.Verifier, error) {
	publicKey := new(mldsa44.PublicKey)
	if err := publicKey.UnmarshalBinary(data); err != nil {
		return nil, fmt.Errorf("unmarshal public key: %w", err)
	}

	return &Verifier44{publicKey: publicKey}, nil
}
