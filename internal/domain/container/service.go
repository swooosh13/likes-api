package container

import "context"

type service struct {
	storage Service
}

func NewService(storage Storage) Service {
	return &service{storage: storage}
}

func (s *service) AddItemToContainer(ctx context.Context, createItemDTO *CreateItemDTO) error {
	return s.storage.AddItemToContainer(ctx, createItemDTO)
}

func (s *service) Create(ctx context.Context, createContainerDTO *CreateContainerDTO) error {
	return s.storage.Create(ctx, createContainerDTO)
}

func (s *service) FindAll(ctx context.Context) ([]Container, error) {
	return s.storage.FindAll(ctx)
}

func (s *service) FindUserContainers(ctx context.Context, userId string) ([]Container, error) {
	return s.storage.FindUserContainers(ctx, userId)
}

func (s *service) Delete(ctx context.Context, containerId int) error {
	return s.storage.Delete(ctx, containerId)
}
