package services

import (
	"context"
	"graphql_sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
}
type RepositoryService interface {
	GetRepositoryByNameOwner(ctx context.Context, name string, owner string) (*model.Repository, error)
	GetRepositoryById(ctx context.Context, id string) (*model.Repository, error)
}

type Services interface {
	UserService
	RepositoryService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	*repositoryService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
	}
}
