package jwt

import (
	"fmt"
	"time"

	"github.com/charmingruby/swrc/internal/common/infra/security"
	"github.com/dgrijalva/jwt-go"
)

func NewJWTService(issuer, secretKey string) *JWTService {
	return &JWTService{
		issuer:    issuer,
		secretKey: secretKey,
	}
}

type JWTService struct {
	issuer    string
	secretKey string
}

type JWTClaim struct {
	Payload security.TokenPayload
	jwt.StandardClaims
}

func (s *JWTService) GenerateToken(p security.TokenPayload) (string, error) {
	tokenDuration := time.Duration(time.Minute * 60 * 24 * 7) //7 days

	claims := &JWTClaim{
		p,
		jwt.StandardClaims{
			Subject:   p.AccountID,
			Issuer:    s.issuer,
			ExpiresAt: time.Now().Local().Add(tokenDuration).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", nil
	}

	return tokenStr, nil
}

func (j *JWTService) isTokenValid(t *jwt.Token) (interface{}, error) {
	if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
		return nil, fmt.Errorf("invalid token %v", t)
	}

	return []byte(j.secretKey), nil
}

func (j *JWTService) ValidateToken(token string) (security.TokenPayload, error) {
	jwtToken, err := jwt.Parse(token, j.isTokenValid)
	if err != nil {
		return security.TokenPayload{}, err
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return security.TokenPayload{}, fmt.Errorf("unable to parse jwt claims")
	}

	payloadClaims, ok := claims["Payload"].(map[string]interface{})
	if !ok {
		return security.TokenPayload{}, fmt.Errorf("payload is missing or not a map")
	}

	payload := security.TokenPayload{
		AccountID: payloadClaims["account_id"].(string),
		Role:      payloadClaims["role"].(string),
		IsValid:   payloadClaims["is_valid"].(bool),
		Verified:  payloadClaims["verified"].(bool),
	}

	return payload, err
}
