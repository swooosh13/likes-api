package item

type CreateItemDTO struct {
	Name        string `json:"name"`
	ContainerID int    `json:"container_id"`
}
