package encryption

type EncryptionContract interface {
	Hash(plaintext string) (hashed_string string)

	Compare(plaintext string, hashed_string string) (is_correct bool)
}
