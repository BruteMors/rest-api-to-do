package rest_api_to_do

type TodoList struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Done        bool   `json:"done,omitempty"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
