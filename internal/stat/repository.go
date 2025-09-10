package stat

import (
	"http_server/pakages/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	Database *db.Db
}

func NewStatRepository(db *db.Db) *StatRepository {
	return &StatRepository{Database: db}
}

func (repo *StatRepository) AddClick(linkId uint) {
	var stat Stat
	curentDate := datatypes.Date(time.Now())
	repo.Database.Find(&stat, "link_id = ? and date = ?", linkId, curentDate)

	if stat.ID == 0 {
		repo.Database.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   curentDate,
		})
	} else {
		stat.Clicks += 1
		repo.Database.Save(&stat)
	}

}

// func (repo *StatRepository) GetByEmail(email string) (*Stat, error) {
// 	var stat Stat
// 	result := repo.Database.DB.First(&stat, "email = ?", email)

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &stat, nil
// }
