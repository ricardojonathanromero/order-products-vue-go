package service

import (
	"context"
	"github.com/auth0/go-auth0/authentication"
	"github.com/auth0/go-auth0/authentication/oauth"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/constants"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/auth/pkg/domain/entities"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/logger"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/transform"
	"github.com/ricardojonathanromero/order-products-vue-go/backend/utilities/utils"
)

type AuthService interface {
	Login(ctx context.Context, req *entities.LoginReq) (*entities.Token, error)
	RenewSession(ctx context.Context, req *entities.RefreshTokenReq) (*entities.Token, error)
}

type authServiceImpl struct {
	client   *authentication.Authentication
	log      logger.Logger
	scope    string
	audience string
	realm    string
}

func New(client *authentication.Authentication, log logger.Logger) AuthService {
	return &authServiceImpl{
		client:   client,
		log:      log,
		scope:    utils.GetEnv(constants.Auth0Scope, constants.DefaultAuth0Score),
		audience: utils.GetEnv(constants.Auth0Audience, constants.Empty),
		realm:    utils.GetEnv(constants.Auth0Realm, constants.Empty),
	}
}

func (s *authServiceImpl) Login(ctx context.Context, req *entities.LoginReq) (*entities.Token, error) {
	var opts oauth.IDTokenValidationOptions
	s.log.Debug("start service Login")

	body := oauth.LoginWithPasswordRequest{
		Username: req.Username,
		Password: req.Password,
		Scope:    s.scope,
		Audience: s.audience,
		Realm:    s.realm,
	}

	s.log.Debug("auth0 body request")
	s.log.Debug(transform.StructToString(body))

	s.log.Debug("trying to create session")
	session, err := s.client.OAuth.LoginWithPassword(ctx, body, opts)
	if err != nil {
		s.log.Errorf("error creating session in auth0: %v", err)
		return nil, err
	}

	s.log.Info("session generated!")
	s.log.Debug(transform.StructToString(session))
	return &entities.Token{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		Type:         session.TokenType,
		ExpiresIn:    session.ExpiresIn,
	}, nil
}

func (s *authServiceImpl) RenewSession(ctx context.Context, req *entities.RefreshTokenReq) (*entities.Token, error) {
	var opts oauth.IDTokenValidationOptions
	s.log.Debug("start service renew session")

	body := oauth.RefreshTokenRequest{
		RefreshToken: req.Token,
		Scope:        s.scope,
	}

	s.log.Debug("auth0 body refresh token request")
	s.log.Debug(transform.StructToString(body))

	s.log.Debug("trying to renew session")
	session, err := s.client.OAuth.RefreshToken(ctx, body, opts)
	if err != nil {
		s.log.Errorf("error renewing session in auth0: %v", err)
		return nil, err
	}

	s.log.Info("session renewed!")
	s.log.Debug(transform.StructToString(session))
	return &entities.Token{
		AccessToken:  session.AccessToken,
		RefreshToken: session.RefreshToken,
		Type:         session.TokenType,
		ExpiresIn:    session.ExpiresIn,
	}, nil
}
