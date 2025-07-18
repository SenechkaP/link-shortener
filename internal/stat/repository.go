package stat

import (
	"advpractice/pkg/db"
	"time"

	"gorm.io/datatypes"
)

type StatRepository struct {
	Database *db.Db
}

func NewStatRepository(database *db.Db) *StatRepository {
	return &StatRepository{Database: database}
}

// func (repo *StatRepository) GetStat(linkId uint, date datatypes.Date) (stats []Stat) {
// 	repo.Database.
// 		Table("stats").
// 		Where("deleted_at is null").
// 		Order("id ASC").
// 		Scan(&links)
// }

func (repo *StatRepository) AddClick(linkId uint) {
	today := datatypes.Date(time.Now())
	var stat Stat
	repo.Database.Find(&stat, "link_id = ? and date = ?", linkId, today)
	if stat.ID == 0 {
		repo.Database.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   today,
		})
	} else {
		stat.Clicks += 1
		repo.Database.Save(&stat)
	}
}
