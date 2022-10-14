package codegenerator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXLZD(t *testing.T) {
	otp := CodeGenXLZD().Generate()

	assert.NotEmpty(t, otp)
}
