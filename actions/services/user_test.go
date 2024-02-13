package services

import (
	"meetme/be/actions/repositories/interfaces"
	"meetme/be/utils"
	"testing"
)

type mockUserRepo struct {
	getFunc func() (interface{}, error)
}

func (m *mockUserRepo) GetUsers() (interface{}, error) {
	return m.getFunc()
}

func TestUserService_GetUsers(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		_ = &mockUserRepo{
			getFunc: func() (interface{}, error) {
				return utils.DataResponse{Data: interfaces.UserResponse{
					Username: "admin",
				}}, nil
			},
		}
		//service := NewUserService(repo)
		//_, err := service.GetUsers()
		//assert.NoError(t, err)
	})
}
