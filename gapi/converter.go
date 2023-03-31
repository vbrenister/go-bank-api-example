package gapi

import (
	db "github.com/vbrenister/go-bank-api-example/db/sqlc"
	"github.com/vbrenister/go-bank-api-example/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:         user.Username,
		FullName:         user.FullName,
		Email:            user.Email,
		PasswordChagedAt: timestamppb.New(user.PasswordChanged),
		CreatedAt:        timestamppb.New(user.CreatedAt),
	}
}
