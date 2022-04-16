package container

type CreateContainerDTO struct {
	UserId string `json:"user_id"`
	Name   string `json:"name"`
}

type CreateItemDTO struct {
	ContainerId int    `json:"container_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
}
