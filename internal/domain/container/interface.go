package container

import (
	"context"
)

type Storage interface {
	FindAll(ctx context.Context) ([]Container, error)
	FindUserContainers(ctx context.Context, userId string) ([]Container, error)
	Create(ctx context.Context, createContainerDTO *CreateContainerDTO) error
	AddItemToContainer(ctx context.Context, createItemDTO *CreateItemDTO) error
	FindContainerItems(ctx context.Context, containerId int) ([]ContainerItem, error)
	Delete(ctx context.Context, containerId int) error
}

type Service interface {
	FindAll(ctx context.Context) ([]Container, error)
	FindUserContainers(ctx context.Context, userId string) ([]Container, error)
	Create(ctx context.Context, createContainerDTO *CreateContainerDTO) error
	AddItemToContainer(ctx context.Context, createItemDTO *CreateItemDTO) error
	Delete(ctx context.Context, containerId int) error
}
