package models

import (
	"context"
	"errors"

	internalErrors "github.com/devWaylander/pvz_store/pkg/errors"
	"github.com/devWaylander/pvz_store/pkg/log"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const authPrincipalKey contextKey = "authPrincipal"

type AuthPrincipal struct {
	UserID int64  `json:"uid"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

/*
  служат исключительно для передачи параметров из middlewares авторизации в имплементацию бизнес-хендлеров
  удовлетворяющих генерируемому api.StrictServerInterface
*/
// GetAuthPrincipal устанавливает пользователя в контексте по копии
func SetAuthPrincipal(ctx context.Context, principal AuthPrincipal) context.Context {
	return context.WithValue(ctx, authPrincipalKey, principal)
}

/*
  служит исключительно для передачи параметров из middlewares авторизации в имплементацию бизнес-хендлеров
  удовлетворяющих генерируемому api.StrictServerInterface
*/
// GetAuthPrincipal достаёт ссылку на пользователя из контекста
func GetAuthPrincipal(ctx context.Context) (*AuthPrincipal, error) {
	val := ctx.Value(authPrincipalKey)
	principal, ok := val.(AuthPrincipal)
	if !ok || principal == (AuthPrincipal{}) {
		err := errors.New(internalErrors.ErrDecodeCtx)
		log.Logger.Err(err).Msg("method GetAuthPrincipal")
		return nil, err
	}

	return &principal, nil
}

type Claims struct {
	AuthPrincipal
	jwt.RegisteredClaims
}

func NewClaims(userID int64, email, role string) Claims {
	return Claims{
		AuthPrincipal: AuthPrincipal{
			UserID: userID,
			Email:  email,
			Role:   role,
		},
	}
}
