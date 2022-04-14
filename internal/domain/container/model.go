package container

import "proj1/internal/domain/item"

type Container struct {
	ID    string      `json:"id"`
	Name  string      `json:"name"`
	Items []item.Item `json:"items"`
}
