package codegenerator

import (
	codegenerator "kang-sayur-backend/infrastructure/code_generator"
	"time"

	"github.com/xlzd/gotp"
)

type xlzd struct{}

// Generate implements codegenerator.CodeGeneratorContract
func (*xlzd) Generate() string {
	totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
	return totp.At(int64(time.Now().Unix()))
}

func CodeGenXLZD() codegenerator.CodeGeneratorContract {
	return &xlzd{}
}
