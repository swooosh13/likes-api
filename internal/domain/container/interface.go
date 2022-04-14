package container

import "context"

type Storage interface {
	FindAll(ctx context.Context) (c []Container, err error)
}

type Service interface {
}
