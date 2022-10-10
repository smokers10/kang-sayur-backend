package identifier

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateID(t *testing.T) {
	uuid := UUID()

	assert.NotEmpty(t, uuid.GenerateID())
}
