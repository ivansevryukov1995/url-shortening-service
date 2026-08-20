package di

import (
	"time"

	"github.com/ivansevryukov1995/url-shortening-service/intertnal/http/dto"
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/model"
)

type IAuthService interface {
	Register(email, password, name string) (string, error)
	Login(email, password string) (string, error)
}

type IStatService interface {
	AddClick()
}

type IUserRepository interface {
	Create(user *model.User) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

type ILinkRepository interface {
	Create(link *model.Link) (*model.Link, error)
	GetByHash(hash string) (*model.Link, error)
	GetById(id uint) (*model.Link, error)
	Update(link *model.Link) (*model.Link, error)
	Delete(id uint) error
	Count() (int64, error)
	GetAll(limit, offset int) ([]model.Link, error)
}

type IStatRepository interface {
	AddClick(linkID uint)
	GetStats(from, to time.Time, by string) []dto.GetStatResponse
}
