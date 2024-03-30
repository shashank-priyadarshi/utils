package jwt

import (
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/shashank-priyadarshi/utilities/data/ports"
)

// Service Provider: Create JWT with required claims
// Parse & Verify JWT as well as JWT signer
// Refresh JWT

type Handler struct {
}

func Handle() (ports.Data, error) {
	return &Handler{}, nil
}

func (h *Handler) Generate() error {
	return nil
}

func (h *Handler) Validate() error {
	return nil
}

func (h *Handler) Refresh() error {
	return nil
}
