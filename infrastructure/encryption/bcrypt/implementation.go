package encryption

import (
	"kang-sayur-backend/infrastructure/encryption"

	"golang.org/x/crypto/bcrypt"
)

type implementation struct{}

// Compare implements encryption.EncryptionContract
func (*implementation) Compare(plaintext string, hashed_string string) (is_correct bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed_string), []byte(plaintext)); err != nil {
		return false
	}

	return true
}

// Hash implements encryption.EncryptionContract
func (*implementation) Hash(plaintext string) (hashed_string string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	return string(hashed)
}

func Bcrypt() encryption.EncryptionContract {
	return &implementation{}
}
