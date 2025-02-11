package interfaces

import (
	"context"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/models"

	"github.com/labstack/echo/v4"
)

type IRefrehTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(echo.Context) error
}
