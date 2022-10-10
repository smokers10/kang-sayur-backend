package encryption

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncryption(t *testing.T) {
	b := Bcrypt()
	result := "$2a$10$wfU9PuV3T/z9IBkI8F75SeYv2AF2hfPltvkAeNxVJUGXUq7rzMUoS" // test123

	t.Run("Test Hash", func(t *testing.T) {
		hashed_string := b.Hash("test123")
		fmt.Println(hashed_string)
		assert.NotEmpty(t, hashed_string)
	})

	t.Run("Test Compare", func(t *testing.T) {
		t.Run("Wrong Comparation", func(t *testing.T) {
			cr := b.Compare("test_encrypt123", result)

			assert.Equal(t, false, cr)
		})

		t.Run("Correct Comparation", func(t *testing.T) {
			cr := b.Compare("test123", result)

			assert.Equal(t, true, cr)
		})
	})
}
