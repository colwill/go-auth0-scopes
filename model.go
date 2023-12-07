package middleware

import (
	"context"
	"go/types"
)

type CustomAuthClaims struct {
	Scope string `json:"scope"`
}

func (c CustomAuthClaims) Validate(ctx context.Context) error {
	if c.Scope == "" {
		return types.Error{Msg: "request not valid"}
	}
	return nil
}
