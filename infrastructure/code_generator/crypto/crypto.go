package crypto

import (
	"crypto/rand"
	"io"
	codegenerator "kang-sayur-backend/infrastructure/code_generator"
)

type cryotoRand struct{}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

// Generate implements codegenerator.CodeGeneratorContract
func (*cryotoRand) Generate() string {
	b := make([]byte, 6)
	n, err := io.ReadAtLeast(rand.Reader, b, 6)
	if n != 6 {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func CryptoRand() codegenerator.CodeGeneratorContract {
	return &cryotoRand{}
}
