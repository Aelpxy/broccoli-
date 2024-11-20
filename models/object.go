package models

import "time"

type Object struct {
	ID          uint   `gorm:"primaryKey;autoIncrement:true"`
	Key         string `gorm:"not null;size:255"`
	Size        int64  `gorm:"not null"`
	ContentType string `gorm:"not null;size:255"`
	Hash        string `gorm:"type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	BucketID    uint   `gorm:"not null"`
	Bucket      Bucket `gorm:"foreignKey:BucketID"`
}

type ObjectMetadata struct {
	ID         string   `msgpack:"id"`
	Path       string   `msgpack:"path"`
	Hash       string   `msgpack:"hash"`
	Private    bool     `msgpack:"private"`
	AccessKeys []string `msgpack:"access_keys"`
}
