package oauth2

import (
	"context"
	"github.com/ory/fosite"
	"github.com/ory/x/errorsx"
)

type TokenExchangeGrantHandler struct {
	Config                   fosite.RFC8693ConfigProvider
	ScopeStrategy            fosite.ScopeStrategy
	AudienceMatchingStrategy fosite.AudienceMatchingStrategy
}

var tokenExchangeGrantType = string(fosite.GrantTypeTokenExchange)

func (c *TokenExchangeGrantHandler) HandleTokenEndpointRequest(ctx context.Context, request fosite.AccessRequester) error {
	if !c.CanHandleTokenEndpointRequest(ctx, request) {
		return errorsx.WithStack(fosite.ErrUnknownRequest)
	}

	client := request.GetClient()
	if client.IsPublic() {
		return errorsx.WithStack(fosite.ErrInvalidGrant.WithHintf("The OAuth 2.0 Client is marked as public and is thus not allowed to use authorization grant '%s'.", tokenExchangeGrantType))
	}

	if !client.GetGrantTypes().Has(tokenExchangeGrantType) {
		return errorsx.WithStack(fosite.ErrUnauthorizedClient.WithHintf("The OAuth 2.0 Client is not allowed to use authorization grant '%s'.", tokenExchangeGrantType))
	}
}

func (c *TokenExchangeGrantHandler) CanHandleTokenEndpointRequest(ctx context.Context, request fosite.AccessRequester) bool {
	return request.GetGrantTypes().Has(tokenExchangeGrantType)
}
