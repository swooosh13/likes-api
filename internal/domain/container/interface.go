package container

import (
	"context"
)

type StorageItem interface {
	DeleteItem(ctx context.Context, itemId int) error
}

type Storage interface {
	FindAll(ctx context.Context) ([]Container, error)
	FindUserContainers(ctx context.Context, userId string) ([]Container, error)
	Create(ctx context.Context, createContainerDTO *CreateContainerDTO) error
	Delete(ctx context.Context, containerId int) error
	UpdateContainer(ctx context.Context, updateContainerDTO *UpdateContainerDTO, containerId int) error
	AddItemToContainer(ctx context.Context, createItemDTO *CreateItemDTO) error
	GetContainerItems(ctx context.Context, userId string, containerId int) ([]ContainerItem, error)

	StorageItem
}

type ServiceItem interface {
	DeleteItem(ctx context.Context, itemId int) error
}
type Service interface {
	FindAll(ctx context.Context) ([]Container, error)
	FindUserContainers(ctx context.Context, userId string) ([]Container, error)
	Create(ctx context.Context, createContainerDTO *CreateContainerDTO) error
	Delete(ctx context.Context, containerId int) error
	UpdateContainer(ctx context.Context, updateContainerDTO *UpdateContainerDTO, containerId int) error
	AddItemToContainer(ctx context.Context, createItemDTO *CreateItemDTO) error
	GetContainerItems(ctx context.Context, userId string, containerId int) ([]ContainerItem, error)

	ServiceItem
}
