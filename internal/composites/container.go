package composites

import (
	"proj1/internal/domain/container"
	"proj1/internal/handlers/api"
	container2 "proj1/internal/handlers/api/container"
	container1 "proj1/internal/repository/container"
)

type ContainerComposite struct {
	Storage container.Storage
	Service container.Service
	Handler api.Handler
}

func NewContainerComposite(pgComposite *PostgresDBComposite) *ContainerComposite {
	storage := container1.NewRepository(pgComposite.Client)
	service := container.NewService(storage)
	handler := container2.NewHandler(service)

	return &ContainerComposite{Storage: storage, Service: service, Handler: handler}
}
