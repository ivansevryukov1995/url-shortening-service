package auth_test

import (
	"errors"
	"testing"

	"github.com/ivansevryukov1995/url-shortening-service/intertnal/auth"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/model"
)

type MockUserRepository struct {
	database map[string]*model.User
}

func NewMockUserRepository(database map[string]*model.User) *MockUserRepository {
	return &MockUserRepository{database: database}
}

func (repo *MockUserRepository) Create(user *model.User) (*model.User, error) {
	repo.database[user.Email] = user
	return user, nil
}

func (repo *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	user, ok := repo.database[email]
	if !ok {
		return nil, nil
	}
	return user, auth.ErrUserExists
}
func TestRegisterService(t *testing.T) {
	testCases := []struct {
		tName string
		email string
		pass  string
		name  string

		emailExpec string
		err        error
	}{
		{
			tName: "RegisterSuccess",
			email: "a@a.ru",
			pass:  "1",
			name:  "John",

			emailExpec: "a@a.ru",
			err:        nil,
		},
		{
			tName: "UserExists",
			email: "a@a.ru",
			pass:  "1",
			name:  "John",

			emailExpec: "",
			err:        auth.ErrUserExists,
		},
	}

	database := make(map[string]*model.User, 0)
	userRepo := NewMockUserRepository(database)

	authService := auth.NewAuthService(userRepo)

	for _, tC := range testCases {
		t.Run(tC.tName, func(t *testing.T) {
			emailGot, err := authService.Register(tC.email, tC.pass, tC.name)
			if !errors.Is(err, tC.err) {
				t.Fatalf("got %v, expected %v", err, tC.err)
				return
			}

			if emailGot != tC.emailExpec {
				t.Fatalf("Email got %s not equal expected %s", emailGot, tC.emailExpec)
				return
			}
		})
	}
}
func TestLoginService(t *testing.T) {
	testCases := []struct {
		tName string
		email string
		pass  string

		emailExpec string
		err        error
	}{
		{
			tName: "LoginSuccess",
			email: "a@a.ru",
			pass:  "1",

			emailExpec: "a@a.ru",
			err:        nil,
		},
		{
			tName: "WrongCredentials",
			email: "a2@a.ru",
			pass:  "1",

			emailExpec: "",
			err:        auth.ErrWrongCredentials,
		},
	}

	database := make(map[string]*model.User, 0)
	userRepo := NewMockUserRepository(database)

	authService := auth.NewAuthService(userRepo)

	_, _ = authService.Register("a@a.ru", "1", "John")

	for _, tC := range testCases {
		t.Run(tC.tName, func(t *testing.T) {
			emailGot, err := authService.Login(tC.email, tC.pass)
			if !errors.Is(err, tC.err) {
				t.Fatalf("got %v, expected %v", err, tC.err)
				return
			}

			if emailGot != tC.emailExpec {
				t.Fatalf("Email got %s not equal expected %s", emailGot, tC.emailExpec)
				return
			}
		})
	}
}
