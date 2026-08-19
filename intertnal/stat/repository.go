package stat

import (
	"time"

	"github.com/ivansevryukov1995/url-shortening-service/pkg/db"
	"gorm.io/datatypes"
)

type StatRepository struct {
	*db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{
		Db: db,
	}
}

func (repo *StatRepository) AddClick(linkID uint) {
	var stat Stat

	currentDate := datatypes.Date(time.Now())

	repo.Db.Find(&stat, "link_id = ? and date = ?", linkID, currentDate)

	if stat.ID == 0 {
		repo.DB.Create(&Stat{
			LinkID: linkID,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		repo.Db.Save(&stat)
	}
}
