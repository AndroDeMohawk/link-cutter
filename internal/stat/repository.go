package stat

import (
	"time"

	"github.com/AndroDeMohawk/link-cutter/pkg/db"
	"gorm.io/datatypes"
)

type Repository struct {
	*db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{
		Db: db,
	}
}

func (r *Repository) AddClick(linkId uint) {
	currentDate := datatypes.Date(time.Now())
	var stat Stat
	r.Db.Find(&stat, "link_id = ? and date =?", linkId, currentDate)
	if stat.ID == 0 {
		r.Db.Create(&Stat{
			LinkId: linkId,
			Clicks: 1,
			Date:   currentDate,
		})
	} else {
		stat.Clicks++
		r.Db.Save(&stat)
	}
}

func (r *Repository) GetStats(by string, from, to time.Time) []GetResponse {
	var stats []GetResponse
	var selectQuery string
	switch by {
	case FilterByDay:
		selectQuery = "to_char(date, 'YYYY-MM-DD') as period, sum(clicks)"
	case FilterByMonth:
		selectQuery = "to_char(date, 'YYYY-MM') as period, sum(clicks)"
	}
	r.DB.Table("stats").
		Select(selectQuery).
		Where("date BETWEEN ? AND ?", from, to).
		Group("period").
		Order("period").
		Scan(&stats)
	return stats
}
