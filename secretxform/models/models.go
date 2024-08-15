package models

type (
	Type      int
	Algorithm string

	Config struct {
		Type                   Type
		SigningAlg, SigningKey string
		CurveBits              int
		DefaultClaims          map[string]interface{}
	}
)

func (a Algorithm) String() string {
	return string(a)
}
