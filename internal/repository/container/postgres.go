package container

import (
	"context"
	"proj1/internal/domain/container"
	"proj1/pkg/postgresdb"
)

type repository struct {
	client postgresdb.Client
}

func NewRepository(client postgresdb.Client) container.Storage {
	return &repository{
		client: client,
	}
}

func (r *repository) FindAll(ctx context.Context) (c []container.Container, err error) {
	return nil, nil
}
