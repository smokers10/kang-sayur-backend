package database

import (
	"kang-sayur-backend/infrastructure/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongoDB(t *testing.T) {
	ctx, cncl := helper.InitCtxTimeout()
	defer cncl()

	dbInit := MongoInit()
	db, er := dbInit.MongoDB()

	t.Run("Initialization Check", func(t *testing.T) {
		assert.Empty(t, er)
	})

	t.Run("Connection Check", func(t *testing.T) {
		erPing := db.Client().Ping(ctx, nil)

		assert.Empty(t, erPing)
	})
}
