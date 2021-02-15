package handlers

import (
	"context"
	"skeleton-code/components"
	"skeleton-code/components/member"
	"skeleton-code/proto/generated"
)

type memberHandler struct {
	ms member.IMemberService
}

func NewMemberHandler(ctx components.Context) *memberHandler {
	return &memberHandler{
		ms: ctx.GetMemberService(),
	}
}

func (m memberHandler) RegisterMember(ctx context.Context, request *generated.RegisterMemberRequest) (*generated.Member, error) {
	return m.ms.RegisterMember(ctx, request)
}

func (m memberHandler) LoginMember(ctx context.Context, request *generated.LoginRequest) (*generated.AccessToken, error) {
	return m.ms.LoginMember(ctx, request)
}
