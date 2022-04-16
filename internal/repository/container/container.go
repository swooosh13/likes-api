package container

import (
	"context"
	"errors"
	"fmt"
	"proj1/internal/domain/container"
	"proj1/pkg/logger"
	"proj1/pkg/pgdb"

	"github.com/jackc/pgconn"
)

type repository struct {
	client pgdb.Client
}

func NewRepository(client pgdb.Client) container.Storage {
	return &repository{
		client: client,
	}
}

func (r *repository) FindAll(ctx context.Context) (cs []container.Container, err error) {
	q := `
		SELECT id, user_id, name FROM public.container;
	`

	rows, err := r.client.Query(ctx, q)
	if err != nil {
		return nil, err
	}

	cs = make([]container.Container, 0)
	for rows.Next() {
		var c container.Container

		err = rows.Scan(&c.ID, &c.UserId, &c.Name)
		if err != nil {
			return nil, err
		}

		cs = append(cs, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cs, nil
}

func (r *repository) FindUserContainers(ctx context.Context, userId string) ([]container.Container, error) {
	q := `
		SELECT id, user_id, name FROM public.container WHERE user_id = $1;
	`

	cs := make([]container.Container, 0)
	rows, err := r.client.Query(ctx, q, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var c container.Container
		err = rows.Scan(&c.ID, &c.UserId, &c.Name)
		if err != nil {
			return nil, err
		}

		items, err := r.FindContainerItems(ctx, c.ID)
		if err != nil {
			return nil, err
		}
		c.Items = items

		cs = append(cs, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return cs, nil
}

func (r *repository) Create(ctx context.Context, createContainerDTO *container.CreateContainerDTO) error {
	q := `
		INSERT INTO public.container
			(user_id, name)
		VALUES
			($1, $2)
		returning id;
	`

	var id int
	if err := r.client.QueryRow(ctx, q, createContainerDTO.UserId, createContainerDTO.Name).Scan(&id); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			logger.Error(newErr.Error())
			return err
		}
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, containerId int) error {
	q := `
		DELETE FROM container_item WHERE container_id IN (SELECT id FROM container WHERE id = $1);
	`

	fmt.Println(containerId)
	tx, err := r.client.Begin(ctx)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, q, containerId)
	if err != nil {
		tx.Rollback(ctx)
		return nil
	}

	q = `DELETE FROM container WHERE id = $1;`

	_, err = tx.Exec(ctx, q, containerId)
	if err != nil {
		tx.Rollback(ctx)
		return nil
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	return nil
}
