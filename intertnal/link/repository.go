package link

import (
	"github.com/ivansevryukov1995/url-shortening-service/intertnal/model"
	"github.com/ivansevryukov1995/url-shortening-service/pkg/db"
	"gorm.io/gorm/clause"
)

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepository(database *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: database,
	}
}

func (repo *LinkRepository) Create(link *model.Link) (*model.Link, error) {
	result := repo.Database.DB.Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*model.Link, error) {
	var link model.Link
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil

}

func (repo *LinkRepository) GetById(id uint) (*model.Link, error) {
	var link model.Link
	result := repo.Database.DB.First(&link, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &link, nil

}

func (repo *LinkRepository) Update(link *model.Link) (*model.Link, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}

	return link, nil

}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&model.Link{}, id)
	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (repo *LinkRepository) Count() (int64, error) {
	var count int64

	result := repo.Database.
		Table("links").
		Where("deleted_at is null").
		Count(&count)
	if result.Error != nil {
		return count, result.Error
	}

	return count, nil
}

func (repo *LinkRepository) GetAll(limit, offset int) ([]model.Link, error) {
	var links []model.Link

	result := repo.Database.
		Table("links").
		Where("deleted_at is null").
		Order("id asc").
		Limit(limit).
		Offset(offset).
		Scan(&links)
	if result.Error != nil {
		return nil, result.Error
	}

	return links, nil

}
