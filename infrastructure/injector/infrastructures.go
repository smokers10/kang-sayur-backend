package injector

import (
	codegenerator "kang-sayur-backend/infrastructure/code_generator"
	"kang-sayur-backend/infrastructure/code_generator/crypto"
	"kang-sayur-backend/infrastructure/database"
	"kang-sayur-backend/infrastructure/encryption"
	bcrypt "kang-sayur-backend/infrastructure/encryption/bcrypt"
	"kang-sayur-backend/infrastructure/identifier"
	uuid "kang-sayur-backend/infrastructure/identifier/uuid"
	jsonwebtoken "kang-sayur-backend/infrastructure/json_web_token"
	mailer "kang-sayur-backend/infrastructure/mailer"
	"kang-sayur-backend/infrastructure/mailer/smtp"

	gojwt "kang-sayur-backend/infrastructure/json_web_token/jwt"

	"go.mongodb.org/mongo-driver/mongo"
)

type Infrastructures struct{}

func (sp *Infrastructures) DatabaseMongodb() *mongo.Database {
	db, err := database.MongoInit().MongoDB()
	if err != nil {
		panic(err)
	}

	return db
}

func (sp *Infrastructures) Encryption() *encryption.EncryptionContract {
	enc := bcrypt.Bcrypt()
	return &enc
}

func (sp *Infrastructures) Identifier() *identifier.IdentifierContract {
	id := uuid.UUID()
	return &id
}

func (sp *Infrastructures) JsonWebToken() *jsonwebtoken.JWTContract {
	jwt := gojwt.JsonWebToken()
	return &jwt
}

func (sp *Infrastructures) Mailer() *mailer.Contract {
	smtp := smtp.NativeSMTP()
	return &smtp
}

func (sp *Infrastructures) CodeGenerator() *codegenerator.CodeGeneratorContract {
	codegen := crypto.CryptoRand()
	return &codegen
}

func InfrastructureInjector() *Infrastructures {
	return &Infrastructures{}
}
