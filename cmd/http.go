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
	e.GET("/healthcheck", d.HealthCheckAPI.HealthCheck)

	userV1 := e.Group("/user/v1")
	userV1.POST("/register", d.UserAPI.RegisterUser)
	userV1.POST("/register/admin", d.UserAPI.RegisterAdmin)
	userV1.POST("/login", d.UserAPI.LoginUser)
	userV1.POST("/login/admin", d.UserAPI.LoginAdmin)
	userV1.GET("/profile", d.UserAPI.GetProfile, d.MiddlewareValidateAuth)
	userV1.PUT("/refresh-token", d.RefreshTokenAPI.RefreshToken, d.MiddlewareRefreshToken)

	e.Start(":" + helpers.GetEnv("PORT", "9001"))
}

type Dependency struct {
	HealthCheckAPI *api.HealthCheckAPI
	UserAPI        interfaces.IUserAPI
	UserRepository interfaces.IUserRepository
	RefreshTokenAPI interfaces.IRefreshTokenHandler
}

func dependencyInject() Dependency {
	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}
	userSvc := &services.UserService{
		UserRepo: userRepo,
	}
	userAPI := &api.UserAPI{
		UserService: userSvc,
	}
	refreshTokenSvc := &services.RefreshTokenService{
		UserRepo: userRepo,
	}
	refreshTokenAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshTokenSvc,
	}

	return Dependency{
		UserRepository: userRepo,
		HealthCheckAPI: &api.HealthCheckAPI{},
		UserAPI:        userAPI,
		RefreshTokenAPI: refreshTokenAPI,

	}
}

