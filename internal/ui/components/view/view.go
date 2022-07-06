package view

import "fmt"

type Model struct {
	Id int
}

func NewModel(id int) Model {
	return Model{
		Id: id,
	}
}

func (m Model) View() string {
	return fmt.Sprintf("Hello from view %d", m.Id)
}
