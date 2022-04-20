package container

import "context"

func (s *service) DeleteItem(ctx context.Context, itemId int) error {
	return s.storage.DeleteItem(ctx, itemId)
}
