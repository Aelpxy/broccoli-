package models

import "time"

type Object struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Key         string `gorm:"not null;size:255"`
	Size        int64  `gorm:"not null"`
	ContentType string `gorm:"not null;size:255"`
	Metadata    string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BucketID    uint   `gorm:"not null"`
	Bucket      Bucket `gorm:"foreignKey:BucketID"`
}
