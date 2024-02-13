package repository

import (
	"errors"
	"fmt"
	"matchkeeper/internal/models"
)

var (
	NotFoundErr = errors.New("not found")
)

type InMemRepository struct {
	list map[int]models.Match
}

func NewInMemRepository() *InMemRepository {
	list := make(map[int]models.Match)
	return &InMemRepository{
		list,
	}
}

func (m InMemRepository) Create(match models.Match) error {
	fmt.Println("inMemStore Create Method: ", match)
	m.list[match.Id] = match
	//fmt.Println("List: ", m.list)
	return nil
}

func (m InMemRepository) List() (map[int]models.Match, error) {
	fmt.Println("In List Function: ", m.list)
	return m.list, nil
}

func (m InMemRepository) Get(id int) (models.Match, error) {

	for _, match := range m.list {
		if match.Id == id {
			return match, nil
		}
	}
	return models.Match{}, NotFoundErr
}

func (m InMemRepository) Update(id int, match models.Match) error {

	if match.Id == 0 {
		// Invalid Id, cannot update
		return NotFoundErr
	}

	for index, matchInput := range m.list {
		if matchInput.Id == id {
			m.list[index] = matchInput
		}
	}

	return NotFoundErr
}

func (m InMemRepository) Delete(id int) error {
	//slices.Delete(m.list, id, id)
	delete(m.list, id)
	return NotFoundErr
}
