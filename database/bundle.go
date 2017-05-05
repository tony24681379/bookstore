package database

import "time"

//BundleMaster define
type BundleMaster struct {
	ID         uint `gorm:"AUTO_INCREMENT"`
	BundleName string
	Note       string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time
}

//BundleDetail define
type BundleDetail struct {
	BundleID  uint
	BookID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
