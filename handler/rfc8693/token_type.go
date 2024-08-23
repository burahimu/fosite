package rfc8693

import "context"

const (
	// AccessTokenType is the access token type issued by the same provider
	AccessTokenType string = "urn:ietf:params:oauth:token-type:access_token" // #nosec G101
	// RefreshTokenType is the refresh token type issued by the same provider
	RefreshTokenType string = "urn:ietf:params:oauth:token-type:refresh_token" // #nosec G101
	// IDTokenType is the id_token type issued by the same provider
	IDTokenType string = "urn:ietf:params:oauth:token-type:id_token" // #nosec G101
	// JWTTokenType is the JWT type that may be issued by a different provider
	JWTTokenType string = "urn:ietf:params:oauth:token-type:jwt" // #nosec G101
	// SAML1TokenType is the SAML 1.1 type that may be issued by a different provider
	SAML1TokenType string = "urn:ietf:params:oauth:token-type:saml1" // #nosec G101
	// SAML2TokenType is the SAML 2.0 type that may be issued by a different provider
	SAML2TokenType string = "urn:ietf:params:oauth:token-type:saml2" // #nosec G101
)

type DefaultTokenType struct {
	Name string
}

func (c *DefaultTokenType) GetName(ctx context.Context) string {
	return c.Name
}

func (c *DefaultTokenType) GetType(ctx context.Context) string {
	return c.Name
}
