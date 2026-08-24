package auth_test

import (
	"testing"

	"github.com/ivansevryukov1995/url-shortening-service/intertnal/auth"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/model"
)

type MockUserRepository struct {
}

func (repo *MockUserRepository) Create(user *model.User) (*model.User, error) {
	return &model.User{
		Email:    "a@a.ru",
		Password: "1",
		Name:     "John",
	}, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}

func TestRegisterSuccess(t *testing.T) {
	param := struct {
		email string
		pass  string
		name  string
	}{
		email: "a@a.ru",
		pass:  "1",
		name:  "John",
	}

	authService := auth.NewAuthService(&MockUserRepository{})

	email, err := authService.Register(param.email, param.pass, param.name)
	if err != nil {
		t.Fatal(err)
		return
	}
	if email != param.email {
		t.Fatalf("Email %s not equal %s", param.email, email)
		return
	}

}
