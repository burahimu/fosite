package rfc8693

import (
	"context"
	"github.com/ory/fosite"
	"github.com/ory/fosite/handler/oauth2"
	"time"
)

// RFC8693CoreStorage is the storage needed for the RFC8693 handler
type RFC8693CoreStorage interface {
	oauth2.CoreStorage

	// SetTokenExchangeCustomJWT marks a JTI as known for the given
	// expiry time. It should atomically check if the JTI
	// already exists and fail the request, if found.
	SetTokenExchangeCustomJWT(ctx context.Context, jti string, exp time.Time) error

	// GetSubjectForTokenExchange computes the session subject and is used for token types where there is no way
	// to know the subject value. For some token types, such as access and refresh tokens, the subject is well-defined
	// and this function is not called.
	GetSubjectForTokenExchange(ctx context.Context, requester fosite.Requester, subjectToken map[string]interface{}) (string, error)
}
