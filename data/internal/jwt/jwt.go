package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/shashank-priyadarshi/utilities/data/models"
	"github.com/shashank-priyadarshi/utilities/data/ports"
	"strings"
	"time"
)

type Handler struct {
	signingMethod jwt.SigningMethod
	signingKey    string
	defaultClaims map[string]interface{}
}

func Handle(config *models.Config) (ports.Data, error) {
	return &Handler{
		signingMethod: jwt.GetSigningMethod(config.SigningAlg),
		signingKey:    config.SigningKey,
		defaultClaims: config.DefaultClaims,
	}, nil
}

func (h *Handler) Generate(customClaims map[string]interface{}) (token string, err error) {

	newToken := jwt.New(h.signingMethod)
	claims := newToken.Claims.(jwt.MapClaims)

	for key, val := range h.defaultClaims {
		claims[key] = val
	}

	for key, val := range customClaims {
		claims[key] = val
	}

	token, err = newToken.SignedString(h.signingKey)

	return token, err
}

func (h *Handler) Validate(token string, customClaims map[string]interface{}) (err error) {

	tokenObj, err := h.getTokenObj(token)

	if err != nil {
		return err
	}

	_, err = h.validate(tokenObj, customClaims)

	return err
}

func (h *Handler) validate(token *jwt.Token, customClaims map[string]interface{}) (claims jwt.MapClaims, err error) {

	claims = token.Claims.(jwt.MapClaims)

	for key, val := range h.defaultClaims {

		var claimVal interface{}
		var ok bool

		if claimVal, ok = claims[key]; !ok {
			return claims, fmt.Errorf("missing default claim")
		}

		if val != claimVal {
			return claims, fmt.Errorf("default claim mismatch")
		}
	}

	for key, val := range customClaims {

		var claimVal interface{}
		var ok bool

		if claimVal, ok = claims[key]; !ok {
			return claims, fmt.Errorf("missing custom claim")
		}

		if val != claimVal {
			return claims, fmt.Errorf("custom claim mismatch")
		}
	}

	return claims, err
}

func (h *Handler) Refresh(token string, customClaims map[string]interface{}) (refreshedToken string, err error) {

	var tokenObj *jwt.Token
	tokenObj, err = h.getTokenObj(token)
	if err != nil {
		return refreshedToken, err
	}

	var claims jwt.MapClaims
	if claims, err = h.validate(tokenObj, customClaims); err != nil {
		return refreshedToken, err
	}

	claims["exp"] = time.Now().Add(15 * time.Minute)
	refreshedToken, err = tokenObj.SignedString(h.signingKey)

	return refreshedToken, err
}

func (h *Handler) getTokenObj(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		expectedSigningMethod, tokenSigningAlg := h.signingMethod.Alg(), token.Method.Alg()

		if !strings.EqualFold(expectedSigningMethod, tokenSigningAlg) {
			return nil, fmt.Errorf("mismatched signing methods")
		}

		return h.signingKey, nil
	})
}
