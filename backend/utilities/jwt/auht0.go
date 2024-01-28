package jwt

import (
	"context"
	"errors"
	"fmt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jws"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"net/http"
	"strings"
	"time"
)

type Auth0Validator interface {
	ValidateToken(token string) (Token, error)
	SetContext(ctx context.Context) error
}

type auth0ValidatorImpl struct {
	algorithm      jwa.SignatureAlgorithm
	audience       string
	clientSecret   []byte
	clockTolerance time.Duration
	httpClient     *http.Client
	issuer         string
	jwks           *jwk.Cache
	jwksURL        string
	ignoreScope    bool
}

type Token struct {
	Aud    []string
	Iss    string
	Sub    string
	Claims map[string]any
}

func NewAuth0JWTValidator(client *http.Client, domain, audience, clientSecret, algorithm string, ignoreScope bool) (Auth0Validator, error) {
	if i := strings.Index(domain, "//"); i != -1 {
		domain = domain[i+2:]
	}

	alg, err := determineAlg(algorithm)
	if err != nil {
		return nil, err
	}

	auth0Validator := &auth0ValidatorImpl{
		algorithm:      alg,
		clockTolerance: time.Minute,
		issuer:         "https://" + domain + "/",
		audience:       audience,
		clientSecret:   []byte(clientSecret),
		httpClient:     client,
		ignoreScope:    ignoreScope,
	}

	ctx := context.TODO()
	if alg == jwa.RS256 {
		var registerOpts []jwk.RegisterOption
		auth0Validator.jwksURL = auth0Validator.issuer + ".well-known/jwks.json"
		auth0Validator.jwks = jwk.NewCache(ctx)
		if auth0Validator.httpClient != nil {
			registerOpts = append(registerOpts, jwk.WithHTTPClient(auth0Validator.httpClient))
		}

		err = auth0Validator.jwks.Register(auth0Validator.jwksURL, registerOpts...)
		if err != nil {
			return nil, err
		}

		_, err = auth0Validator.jwks.Refresh(ctx, auth0Validator.jwksURL)
		if err != nil {
			return nil, err
		}
	}

	return auth0Validator, nil
}

func (a *auth0ValidatorImpl) SetContext(ctx context.Context) error {
	a.jwks = jwk.NewCache(ctx)
	_, err := a.jwks.Refresh(ctx, a.jwksURL)
	return err
}

func (a *auth0ValidatorImpl) ValidateToken(token string) (Token, error) {
	var result Token
	validator := a.createCustomValidatorFn()

	decodedToken, err := jws.Parse([]byte(token))
	if err != nil {
		return result, err
	}

	headers := decodedToken.Signatures()[0].ProtectedHeaders()

	if headers.Algorithm() != jwa.HS256 && headers.Algorithm() != jwa.RS256 {
		return result, fmt.Errorf("signature algorithm \"%s\" is not supported. Expected the ID token to be signed with \"HS256\" or \"RS256\"", headers.Algorithm())
	}

	if headers.Algorithm() != a.algorithm {
		return result, fmt.Errorf("unexpected signature algorithm; found \"%s\" but expected \"%s\"", headers.Algorithm(), a.algorithm)
	}

	// These options run in the order specified, so changing the order may change the errors returned.
	// Our own validator func should always be ran last.
	keyOpts := []jwt.ParseOption{
		jwt.WithValidate(true),
		jwt.WithAcceptableSkew(a.clockTolerance),
		jwt.WithRequiredClaim("aud"),
		jwt.WithRequiredClaim("sub"),
		jwt.WithRequiredClaim("iss"),
		jwt.WithRequiredClaim("iat"),
		jwt.WithAudience(a.audience),
		jwt.WithIssuer(a.issuer),
		jwt.WithValidator(validator),
	}

	if a.algorithm == jwa.HS256 {
		keyOpts = append(keyOpts, jwt.WithKey(a.algorithm, a.clientSecret))
	} else {
		keyOpts = append(keyOpts, jwt.WithKeySet(jwk.NewCachedSet(a.jwks, a.jwksURL)))
	}

	tokenDetails, err := jwt.Parse([]byte(token), keyOpts...)

	claims := tokenDetails.PrivateClaims()
	if a.ignoreScope {
		delete(claims, "scope")
	}
	result = Token{
		Sub:    tokenDetails.Subject(),
		Iss:    tokenDetails.Issuer(),
		Aud:    tokenDetails.Audience(),
		Claims: claims,
	}

	return result, err
}

func (a *auth0ValidatorImpl) createCustomValidatorFn() jwt.ValidatorFunc {
	return func(_ context.Context, t jwt.Token) jwt.ValidationError {
		if t.Subject() == "" {
			return jwt.NewValidationError(errors.New("sub claim must be a string present in the ID token"))
		}

		if len(t.Audience()) > 1 {
			azp, azpExists := t.Get("azp")

			if azpExists == false {
				return jwt.NewValidationError(errors.New("azp claim must be a string present in the ID token when Audience (aud) claim has multiple values"))
			}

			if azp != a.audience {
				return jwt.NewValidationError(fmt.Errorf("azp claim mismatch in the ID token; expected \"%s\", found \"%s\"", a.audience, azp))
			}
		}

		// calculate expiration > now + 30 secs
		now := time.Now().Add(time.Second * 30)
		if t.Expiration().Sub(now) <= 0 {
			return jwt.NewValidationError(fmt.Errorf("token already expired"))
		}

		return nil
	}
}

func determineAlg(alg string) (jwa.SignatureAlgorithm, error) {
	switch alg {
	case jwa.HS256.String():
		return jwa.HS256, nil
	case jwa.RS256.String():
		return jwa.RS256, nil
	default:
		return "", fmt.Errorf("unsupported algorithm %s provided", alg)
	}
}
