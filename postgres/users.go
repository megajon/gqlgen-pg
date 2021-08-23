package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/megajon/gqlgen-todos/graph/model"
)

type UsersRepo struct {
	DB *pg.DB
}

func (m *UsersRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := m.DB.Model(&users).Select()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (m *UsersRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := m.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
