package container

import (
	"context"
	"errors"
	"fmt"
	"proj1/internal/domain/container"
	"proj1/pkg/logger"

	"github.com/jackc/pgconn"
)

func (r *repository) AddItemToContainer(ctx context.Context, createItemDTO *container.CreateItemDTO) error {
	q := `
		INSERT INTO container_item
			(container_id, name, symbol)
		VALUES
			($1, $2, $3)
		RETURNING id;
	`

	var id int
	if err := r.client.QueryRow(ctx, q, createItemDTO.ContainerId, createItemDTO.Name, createItemDTO.Symbol).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			logger.Error(newErr.Error())
			return newErr
		}
	}

	return nil
}

func (r *repository) DeleteItem(ctx context.Context, itemId int) error {
	q := `
		DELETE FROM container_item WHERE id = $1
	`

	_, err := r.client.Exec(ctx, q, itemId)
	if err != nil {
		return err
	}

	return nil
}
