package todo

type TodoList struct {
	Id          int
	Title       string
	Description string
}

type ListItem struct {
	Id          int
	Title       string
	Description string
	Done        bool
}

type UserLists struct {
	Id     int
	UserId int
	ListId int
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}
