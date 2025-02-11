package api

import (
	"ecommerce-ums/constants"
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/interfaces"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefrehTokenService
}

func (api *RefreshTokenHandler) RefreshToken(e echo.Context) error {
	var (
		log = helpers.Logger
	)

	refreshToken := e.Request().Header.Get("Authorization")

	token := e.Get("token")

	tokenClaim, ok := token.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse to claimToken", nil)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}

	resp, err := api.RefreshTokenService.RefreshToken(e.Request().Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Error("failed on refresh token service", err)
		return helpers.SendResponseHTTP(e, http.StatusInternalServerError, constants.ErrServerError, nil)

	}

	return helpers.SendResponseHTTP(e, http.StatusOK, constants.SuccessMessage, resp)
}
