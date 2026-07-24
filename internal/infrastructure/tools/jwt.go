package tools

import (
	"errors"
	"fmt"
	"platform/internal/infrastructure/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type JWTPayload struct {
	Subject     string // user id
	Issuer      string // "platform" or "admin" or "portal"
	Audience    string // "access" or "refresh"
	Role        string // "user" or "admin" or "manager"
	LastLoginAt *int64
}

type AppClaims struct {
	Role        string `json:"r"`
	LastLoginAt *int64 `json:"llt,omitempty"`
	jwt.RegisteredClaims
}

type JwtService struct {
	secret         []byte
	expiredAccess  time.Duration
	expiredRefresh time.Duration
	log            *zap.Logger
}

func NewJwtService(cfg *configs.JWTConfig, logger *zap.Logger) (*JwtService, error) {
	secret, err := cfg.GetSecret()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize jwt service: %w", err)
	}
	return &JwtService{
		secret:         secret,
		expiredAccess:  cfg.ExpiredAccess,
		expiredRefresh: cfg.ExpiredRefresh,
		log:            logger,
	}, nil
}

func (j *JwtService) CreateToken(p JWTPayload) (string, error) {
	now := time.Now().UTC()

	// Автоматично визначаємо TTL всередині сервісу на основі призначення токена
	var ttl time.Duration
	switch p.Audience {
	case "access":
		ttl = j.expiredAccess
	case "refresh":
		ttl = j.expiredRefresh
	default:
		j.log.Info("unknown audience", zap.String("audience", p.Audience))
		ttl = 15 * time.Minute
	}

	claims := AppClaims{
		Role:        p.Role,
		LastLoginAt: p.LastLoginAt,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   p.Subject,
			Issuer:    p.Issuer,
			Audience:  jwt.ClaimStrings{p.Audience},
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JwtService) ParseAndValidateClaims(tokenString string, expectedAudience string, expectedIssuer string) (*AppClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*AppClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid claims format or token invalid")
	}

	// Валідація Audience
	aud, err := claims.RegisteredClaims.GetAudience()
	if err != nil || len(aud) == 0 || aud[0] != expectedAudience {
		return nil, fmt.Errorf("unexpected audience: expected %s", expectedAudience)
	}

	if claims.Issuer != expectedIssuer {
		return nil, errors.New("invalid issuer")
	}

	if claims.Subject == "" {
		return nil, errors.New("missing subject")
	}

	return claims, nil
}
