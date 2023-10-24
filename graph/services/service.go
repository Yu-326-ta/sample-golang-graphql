package services

import (
	"context"
	"graphql_sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
	GetUserById(ctx context.Context, id string) (*model.User, error)
	ListUsersByID(ctx context.Context, IDs []string) ([]*model.User, error)
}
type RepositoryService interface {
	GetRepositoryByNameOwner(ctx context.Context, name string, owner string) (*model.Repository, error)
	GetRepositoryById(ctx context.Context, id string) (*model.Repository, error)
}

type IssueService interface {
	GetIssueByRepoAndNumber(ctx context.Context, repoID string, number int) (*model.Issue, error)
	ListIssueInRepository(ctx context.Context, repoID string, after *string, before *string, first *int, last *int) (*model.IssueConnection, error)
}

type ProjectService interface {
	GetProjectByOwnerAndNumber(ctx context.Context, ownerID string, number int) (*model.ProjectV2, error)
	GetProjectByID(ctx context.Context, id string) (*model.ProjectV2, error)
	ListProjectByOwner(ctx context.Context, ownerID string, after *string, before *string, first *int, last *int) (*model.ProjectV2Connection, error)
}

type Services interface {
	UserService
	RepositoryService
	IssueService
	ProjectService
	// issueテーブルを扱うIssueServiceなど、他のサービスインターフェースができたらそれらを追加していく
}

type services struct {
	*userService
	*repositoryService
	*issueService
	*projectService
	// issueテーブルを扱うissueServiceなど、他のサービス構造体ができたらフィールドを追加していく
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService:       &userService{exec: exec},
		repositoryService: &repositoryService{exec: exec},
		issueService:      &issueService{exec: exec},
		projectService:    &projectService{exec: exec},
	}
}
