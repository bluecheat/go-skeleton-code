package member

import (
	"context"
	"skeleton-code/database"
	"skeleton-code/proto/generated"
)

type IMemberService interface {
	RegisterMember(ctx context.Context, request *generated.RegisterMemberRequest) (*generated.Member, error)
	LoginMember(ctx context.Context, request *generated.LoginRequest) (*generated.AccessToken, error)
}

type memberService struct {
	db database.Database
}

func NewMemberService(db database.Database) IMemberService {
	return &memberService{
		db: db,
	}
}

func (m memberService) RegisterMember(ctx context.Context, request *generated.RegisterMemberRequest) (*generated.Member, error) {
	panic("implement me")
}

func (m memberService) LoginMember(ctx context.Context, request *generated.LoginRequest) (*generated.AccessToken, error) {
	panic("implement me")
}
