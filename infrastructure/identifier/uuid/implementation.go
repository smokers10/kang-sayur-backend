package identifier

import (
	"kang-sayur-backend/infrastructure/identifier"

	"github.com/google/uuid"
)

type implementation struct{}

// GenerateID implements identifier.IdentifierContract
func (*implementation) GenerateID() (ID string) {
	rand_id, _ := uuid.NewRandom()
	return rand_id.String()
}

func UUID() identifier.IdentifierContract {
	return &implementation{}
}
