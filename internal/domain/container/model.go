package container

type Container struct {
	ID     int             `json:"id"`
	UserId string          `json:"user_id"`
	Name   string          `json:"name"`
	Items  []ContainerItem `json:"items"`
}

type ContainerItem struct {
	ID          int    `json:"id"`
	ContainerId int    `json:"container_id"`
	Name        string `json:"name"`
	Symbol      string `json:"symbol"`
	Priority    int    `json:"priority"`
}
