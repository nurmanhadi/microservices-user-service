package security

import (
	"sync"
	"time"
	"user-service/pkg/enum"
	"user-service/src/config/env"
	"user-service/src/dto"

	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
	UserID string    `json:"user_id"`
	Role   enum.ROLE `json:"role"`
	jwt.RegisteredClaims
}

func JwtGenerateToken(userID string, role enum.ROLE) (*dto.AuthResponse, error) {
	var wg sync.WaitGroup
	var accessTokenString, refreshTokenString string
	var accessErr, refreshErr error

	wg.Add(2)

	go func() {
		defer wg.Done()
		accessClaims := JwtClaims{
			UserID: userID,
			Role:   role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(env.CONF.JWT.Access.Exp) * time.Minute)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
		accessTokenString, accessErr = accessToken.SignedString([]byte(env.CONF.JWT.Access.Secret))
	}()

	go func() {
		defer wg.Done()
		refreshClaims := JwtClaims{
			UserID: userID,
			Role:   role,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(env.CONF.JWT.Refresh.Exp) * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		}

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
		refreshTokenString, refreshErr = refreshToken.SignedString([]byte(env.CONF.JWT.Refresh.Secret))
	}()

	wg.Wait()

	if accessErr != nil {
		return nil, accessErr
	}
	if refreshErr != nil {
		return nil, refreshErr
	}

	return &dto.AuthResponse{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
