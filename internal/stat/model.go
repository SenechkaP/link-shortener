package stat

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	LinkId uint           `json:"link_id" gorm:"index"`
	Clicks int            `json:"clicks" gorm:"default:0"`
	Date   datatypes.Date `json:"date"`
}
