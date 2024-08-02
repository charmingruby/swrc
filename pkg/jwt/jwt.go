package jwt

import (
	"fmt"
	"time"

	"github.com/charmingruby/swrc/internal/account/infra/security"
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

type Payload struct {
	AccountID string `json:"account_id"`
	Role      string `json:"role"`
	IsValid   bool   `json:"is_valid"`
	Verified  bool   `json:"verified"`
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

func (j *JWTService) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, j.isTokenValid)

	return err == nil
}

func (j *JWTService) RetriveTokenPayload(token string) (*Payload, error) {
	t, err := jwt.Parse(token, j.isTokenValid)
	if err != nil {
		return nil, err
	}

	claims, ok := t.Claims.(jwt.MapClaims)

	if !ok {
		return nil, fmt.Errorf("unable to parse jwt claims")
	}

	payload := &Payload{
		AccountID: claims["account_id"].(string),
		Role:      claims["role"].(string),
		IsValid:   claims["role"].(bool),
		Verified:  claims["verified"].(bool),
	}

	return payload, err
}
