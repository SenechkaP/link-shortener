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

func (repo *StatRepository) GetStat(data GetStatRequest) (stats []GetStatResponce) {
	var selectQuery string
	switch data.By {
	case GroupByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case GroupByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	repo.Database.Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", data.From, data.To).
		Group("period").
		Order("period").
		Scan(&stats)
	return
}
