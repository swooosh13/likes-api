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

func (r *repository) FindContainerItems(ctx context.Context, containerId int) ([]container.ContainerItem, error) {
	q := `
		SELECT id, container_id, name, symbol, priority FROM container_item WHERE container_id = $1;
	`

	items := make([]container.ContainerItem, 0)
	rows, err := r.client.Query(ctx, q, containerId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c container.ContainerItem
		err = rows.Scan(&c.ID, &c.ContainerId, &c.Name, &c.Symbol, &c.Priority)
		if err != nil {
			return nil, fmt.Errorf("error in scan: %s", err.Error())
		}

		items = append(items, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
