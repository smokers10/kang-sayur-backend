package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCryptoRand(t *testing.T) {
	otp := CryptoRand().Generate()

	for i := 0; i < 3; i++ {
		fmt.Println(otp)
	}

	assert.NotEmpty(t, otp)
}
