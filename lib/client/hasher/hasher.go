package hasher

type Hasher interface {
	Hash(data []byte) []byte
}
