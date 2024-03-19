package jwt

import (
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/shashank-priyadarshi/utilities/data/ports"
	loggerPort "github.com/shashank-priyadarshi/utilities/logger/ports"
)

// Service Provider: Create JWT with required claims
// Parse & Verify JWT as well as JWT signer
// Refresh JWT

type JWT struct {
	log loggerPort.Logger
}

func NewJWTHandler(log loggerPort.Logger) (ports.Data, error) {
	return &JWT{log: log}, nil
}

func (j *JWT) Generate() error {
	return nil
}

func (j *JWT) Validate() error {
	return nil
}

func (j *JWT) Refresh() error {
	return nil
}
