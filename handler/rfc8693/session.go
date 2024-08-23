package rfc8693

import "github.com/ory/fosite/handler/openid"

type Session interface {
	// SetSubject sets the session's subject.
	SetSubject(subject string)

	SetActorToken(token map[string]interface{})

	GetActorToken() map[string]interface{}

	SetSubjectToken(token map[string]interface{})

	GetSubjectToken() map[string]interface{}

	SetAct(act map[string]interface{})

	AccessTokenClaimsMap() map[string]interface{}
}

type DefaultSession struct {
	*openid.DefaultSession

	ActorToken   map[string]interface{} `json:"-"`
	SubjectToken map[string]interface{} `json:"-"`
	Extra        map[string]interface{} `json:"extra,omitempty"`
}
