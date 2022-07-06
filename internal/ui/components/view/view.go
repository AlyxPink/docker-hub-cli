package view

type Model struct {
	Id int
}

func NewModel(id int) Model {
	return Model{
		Id: id,
	}
}
