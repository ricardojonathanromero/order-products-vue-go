package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/app/service"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/errors"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"net/http"
)

// AuthHandler defines the interface for handling authentication-related HTTP requests.
type AuthHandler interface {
	HandleLogin(c echo.Context) error
	HandleRefreshToken(c echo.Context) error
}

// authHandlerImpl implements AuthHandler
type authHandlerImpl struct {
	authService service.AuthService
	log         logger.Logger
}

func New(authService service.AuthService, log logger.Logger) AuthHandler {
	return &authHandlerImpl{authService: authService, log: log}
}

// HandleLogin godoc
// @Summary Login customer
// @Description  login customer
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request	body	entities.LoginReq	true "request"
// @Success      200  {object}  entities.Token
// @Failure      400  {object}  errors.CustomError
// @Failure      409  {object}  errors.CustomError
// @Failure      403  {object}  errors.CustomError
// @Failure      424  {object}  errors.CustomError
// @Failure      500  {object}  errors.CustomError
// @Security		ApiKeyAuth
// @Router       /login [post]
func (a *authHandlerImpl) HandleLogin(c echo.Context) error {
	a.log.Debug("start handle login")
	ctx := c.Request().Context()

	var req entities.LoginReq

	a.log.Debug("binding request")
	if err := c.Bind(&req); err != nil {
		a.log.Errorf("error binding request: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.InvalidReqBind, err))
	}

	a.log.Debug("validating request")
	if err := c.Validate(req); err != nil {
		a.log.Errorf("error invalid fields: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.InvalidReqBind, err))
	}

	a.log.Info("processing login trx")
	tokenSession, err := a.authService.Login(ctx, &req)
	if err != nil {
		a.log.Errorf("error generating token session: %v", err)
		return c.JSON(http.StatusConflict, err)
	}

	a.log.Info("session created successfully")
	return c.JSON(http.StatusOK, tokenSession)
}

// HandleRefreshToken godoc
// @Summary Refresh Token
// @Description  refresh a session using refresh token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        request	body	entities.RefreshTokenReq	true "request"
// @Success      200  {object}  entities.Token
// @Failure      400  {object}  errors.CustomError
// @Failure      409  {object}  errors.CustomError
// @Failure      403  {object}  errors.CustomError
// @Failure      424  {object}  errors.CustomError
// @Failure      500  {object}  errors.CustomError
// @Security		ApiKeyAuth
// @Router       /refresh-token [post]
func (a *authHandlerImpl) HandleRefreshToken(c echo.Context) error {
	a.log.Debug("start handle refresh token")
	ctx := c.Request().Context()

	var req entities.RefreshTokenReq

	a.log.Debug("binding request")
	if err := c.Bind(&req); err != nil {
		a.log.Errorf("error binding request: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.InvalidReqBind, err))
	}

	a.log.Debug("validating request")
	if err := c.Validate(req); err != nil {
		a.log.Errorf("error invalid fields: %v", err)
		return c.JSON(http.StatusBadRequest, errors.NewError(errors.InvalidReqBind, err))
	}

	a.log.Info("processing renew session trx")
	tokenSession, err := a.authService.RenewSession(ctx, &req)
	if err != nil {
		a.log.Errorf("error re-generating token session: %v", err)
		return c.JSON(http.StatusConflict, err)
	}

	a.log.Info("session renewed successfully")
	return c.JSON(http.StatusOK, tokenSession)
}
