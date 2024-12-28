package signer

import "errors"

type Signer interface {
	Sign(data []byte) ([]byte, error)
	Marshal() ([]byte, error)
	Verifier() (Verifier, error)
	Unmarshal(data []byte) (Signer, error)
}

type Verifier interface {
	Verify(data []byte, signature []byte) (bool, error)
	Marshal() ([]byte, error)
	Unmarshal(data []byte) (Verifier, error)
}

var (
	ErrInvalidPublicKey = errors.New("invalid public key")
	ErrInvalidSignature = errors.New("invalid signature")
	ErrInvalidData      = errors.New("invalid data")
)
