package services

import (
	"context"
	"graphql_sample/graph/db"
	"graphql_sample/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type repositoryService struct {
	exec boil.ContextExecutor
}

func convertRepository(repository *db.Repository) *model.Repository {
	return &model.Repository{
		ID:    repository.ID,
		Name:  repository.Name,
		Owner: &model.User{ID: repository.Owner},
	}
}

func (r *repositoryService) GetRepositoryByNameOwner(ctx context.Context, owner, name string) (*model.Repository, error) {
	repository, err := db.Repositories(
		qm.Select(db.RepositoryColumns.ID, db.RepositoryColumns.Name, db.RepositoryColumns.Owner), // select id, name, owner
		db.RepositoryWhere.Name.EQ(name),
		db.RepositoryWhere.Name.EQ(owner),
	).One(ctx, r.exec)

	if err != nil {
		return nil, err
	}

	return convertRepository(repository), nil
}
