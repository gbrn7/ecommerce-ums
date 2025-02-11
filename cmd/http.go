package cmd

import (
	"ecommerce-ums/helpers"
	"ecommerce-ums/internal/api"
	"ecommerce-ums/internal/interfaces"
	"ecommerce-ums/internal/repository"
	"ecommerce-ums/internal/services"

	"github.com/labstack/echo/v4"
)

func ServeHTTP() {
	d := dependencyInject()

	e := echo.New()
	e.GET("/healthcheck", d.HealthcheckAPI.Healthcheck)

	userV1 := e.Group("user/v1")
	userV1.POST("/register", d.UserAPI.RegisterUser)
	userV1.POST("/register/admin", d.UserAPI.RegisterAdmin)
	userV1.POST("/login", d.UserAPI.LoginUser)
	userV1.POST("/login/admin", d.UserAPI.LoginAdmin)
	userV1.PUT("/refresh-token", d.RefreshTokenAPI.RefreshToken, d.MiddlewareRefreshToken)
	userV1.GET("/profile", d.UserAPI.GetProfile, d.MiddlewareValidateAuth)
	userV1.DELETE("/logout", d.UserAPI.Logout, d.MiddlewareValidateAuth)

	e.Start(":" + helpers.GetEnv("PORT", "9000"))
}

type Dependency struct {
	UserRepository  interfaces.IUserRepository
	HealthcheckAPI  *api.HealthcheckAPI
	UserAPI         interfaces.IUserAPI
	RefreshTokenAPI interfaces.IRefreshTokenHandler
}

func dependencyInject() Dependency {
	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	userSvc := &services.UserService{
		UserRepo: userRepo,
	}

	UserApi := &api.UserAPI{
		UserService: userSvc,
	}

	refreshTokenSvc := &services.RefreshTokenService{
		UserRepo: userRepo,
	}

	refreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokenSvc,
	}

	return Dependency{
		UserRepository:  userRepo,
		HealthcheckAPI:  &api.HealthcheckAPI{},
		UserAPI:         UserApi,
		RefreshTokenAPI: refreshTokenAPI,
	}
}
